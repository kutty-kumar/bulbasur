package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

func NewGRPCServer(logger *logrus.Logger, dbConnectionString string) (*grpc.Server, error){
	grpcServer := grpc.NewServer(
	grpc.KeepaliveParams(
		keepalive.ServerParameters{
			Time:    time.Duration(viper.GetInt("config.keepalive.time")) * time.Second,
			Timeout: time.Duration(viper.GetInt("config.keepalive.timeout")) * time.Second,
		},
	),
	grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			// logging middleware
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),

			// Request-Id interceptor
			requestid.UnaryServerInterceptor(),

			
			// Metrics middleware
			grpc_prometheus.UnaryServerInterceptor,

			// validation middleware
			grpc_validator.UnaryServerInterceptor(),

			// collection operators middleware
			gateway.UnaryServerInterceptor(),
			),
		),
	)
	
	// create new postgres database
	_, err := gorm.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, err
	}
	// register service implementation with the grpcServer
	userSvcConn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("An error %v occurred while instantiating connection with user service", err)
	}
	defer userSvcConn.Close()

	userSvc := pikachu_v1.NewUserServiceClient(userSvcConn)

	return grpcServer, nil
}

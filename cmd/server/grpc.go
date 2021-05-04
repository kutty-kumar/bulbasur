package main

import (
	"bulbasur/pkg/domain/entity"
	"bulbasur/pkg/repo"
	"bulbasur/pkg/svc"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
	"log"
	"os"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kutty-kumar/charminder/pkg"
	charminder "github.com/kutty-kumar/charminder/pkg"
	"github.com/kutty-kumar/ho_oh/bulbasur_v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

func NewGRPCServer(logger *logrus.Logger, userSvcConn *grpc.ClientConn) (*grpc.Server, error) {
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

	dbLogger := gLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gLogger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  gLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,        // Disable color
		},
	)

	// register database
	db, err := gorm.Open(mysql.Open(viper.GetString("database_config.dsn")), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return nil, err
	}

	createTables(db)

	// register repositories
	domainFactory := charminder.NewDomainFactory()
	domainFactory.RegisterMapping("refresh_token", func() charminder.Base {
		return &entity.RefreshToken{}
	})
	dbOption := charminder.WithDb(db)
	externalIdSetter := func(externalId string, base pkg.Base) pkg.Base {
		base.SetExternalId(externalId)
		return base
	}
	setterOption := pkg.WithExternalIdSetter(externalIdSetter)
	refreshTokenGormDao := charminder.NewBaseGORMDao(dbOption, charminder.WithCreator(domainFactory.GetMapping("refresh_token")), setterOption)
	refreshTokenGormRepo := repo.NewRefreshTokenGORMRepo(refreshTokenGormDao)

	// register services
	userSvc := svc.NewUserSvc(pikachu_v1.NewUserServiceClient(userSvcConn))
	authTokenSvc := svc.NewAuthTokenSvc(&refreshTokenGormRepo, userSvc)
	bulbasur_v1.RegisterAuthServiceServer(grpcServer, &authTokenSvc)

	return grpcServer, nil
}

func createTables(db *gorm.DB) {
	err := db.AutoMigrate(entity.RefreshToken{})
	if err != nil {
		log.Fatalf("An error %v occurred while automigrating", err)
	}
}

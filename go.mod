module bulbasur

go 1.14

replace google.golang.org/grpc v1.37.0 => google.golang.org/grpc v1.29.0

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/infobloxopen/atlas-app-toolkit v0.24.0
	github.com/jinzhu/gorm v1.9.16
	github.com/kutty-kumar/charminder v0.0.0-20210505122708-21e591ab714f
	github.com/kutty-kumar/ho_oh v0.0.0-20210504121712-7a8730c49d52
	github.com/prometheus/client_golang v1.10.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0 // indirect
	google.golang.org/grpc v1.37.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.6
	gorm.io/gorm v1.21.9
)

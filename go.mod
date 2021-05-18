module red-bean-anime-server

go 1.14

replace (
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc v1.36.1 => google.golang.org/grpc v1.26.0
)

require (
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/common v0.23.0
	github.com/spf13/viper v1.7.1
	go.etcd.io/etcd v3.3.25+incompatible // indirect
	go.uber.org/zap v1.16.0
	google.golang.org/genproto v0.0.0-20210506142907-4a47615972c2
	google.golang.org/grpc v1.36.1
	google.golang.org/protobuf v1.26.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.6
	gorm.io/gorm v1.21.9
)

module github.com/coder2m/reminder

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/coder2m/component v0.5.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gorilla/websocket v1.4.2 // indirect
	google.golang.org/grpc v1.27.0
	gorm.io/gorm v1.20.9
)

replace (
	github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

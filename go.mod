module github.com/fengyoutian/holingo-micro-gin

go 1.14

replace (
    // github.com/fengyoutian/holingo-util v0.0.0 => ../holingo-util
    // etcd v3版本和 目前grpc v1.27.0 和 v1.28.0 存在不兼容
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	// github.com/fengyoutian/holingo-util v0.0.0
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/wire v0.4.0
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.5.1-0.20200417165434-16db76bee2fb
	github.com/miekg/dns v1.1.29 // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	golang.org/x/tools v0.0.0-20200522201501-cb1345f3a375 // indirect
	google.golang.org/genproto v0.0.0-20200521103424-e9a78aa275b7 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)

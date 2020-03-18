module github.com/fengyoutian/holingo-micro-gin

go 1.14

replace github.com/fengyoutian/holingo-util v0.0.0 => ../holingo-util

require (
	github.com/fengyoutian/holingo-util v0.0.0
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/google/wire v0.4.0
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.2.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
)

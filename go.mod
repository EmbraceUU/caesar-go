module caesar-go

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190813141303-74dc4d7220e7
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190813064441-fde4db37ae7a
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190816200558-6889da9d5479
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-errors/errors v1.0.1
	github.com/go-redis/redis v0.0.0-20190609092923-f8704e4b6b43
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-querystring v1.0.0
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/websocket v1.4.0
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/micro/go-config v1.0.0
	github.com/okcoin-okex/open-api-v3-sdk v0.0.0-20190814030746-f1f75d870f17
	github.com/panjf2000/ants v1.1.1
	github.com/pkg/errors v0.8.1
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/shopspring/decimal v0.0.0-20191009025716-f1972eb1d1f5
	github.com/sirupsen/logrus v1.4.2
	github.com/streadway/amqp v0.0.0-20190815230801-eade30b20f1d
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/ugorji/go v1.1.7
	github.com/viant/toolbox v0.19.1
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.0
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.9.1
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	gopkg.in/go-playground/pool.v3 v3.1.1
)

go 1.12

module github.com/bianjieai/cosmos-sync

go 1.16

require (
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/kaifei-bianjie/msg-parser v0.0.0-20220714072603-982e5144b27f
	github.com/qiniu/qmgo v1.0.4
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.35.0
	github.com/weichang-bianjie/metric-sdk v1.0.0
	go.mongodb.org/mongo-driver v1.7.4
	go.uber.org/zap v1.17.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.44.2-irita-20211102
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.0-irita-210104.0.20210112015006-57e95aa6402f
	google.golang.org/grpc => google.golang.org/grpc v1.35.0
)

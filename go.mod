module github.com/woshidama323/ETHContractPractice

go 1.14

require (
	github.com/ethereum/go-ethereum v1.9.25
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/websocket v1.4.1-0.20190629185528-ae1634f6a989
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli/v2 v2.3.0
	github.com/woshidama323/config v1.0.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/woshidama323/config => ./config

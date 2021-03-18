package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
)

const (
	// host = "mainnet.infura.io"
	host = "rinkeby.infura.io"
	path = "/ws/v3/7a42b8e3d3c64e33ab5737041fc211f9"
	// httpmainnet = "https://mainnet.infura.io/v3/6707728235da4599b4a045c1b40ff0d9"
)

//WebsocketClient ws连接
type WebsocketClient struct {
	Conn *websocket.Conn
	Send chan []byte
}

//NewWebsocketClient 创建新的ws连接
func NewWebsocketClient() (*WebsocketClient, error) {

	// u := url.URL{Scheme: "wss", Host: host, Path: path}
	// rlog.Infoln("connecting to:", u.String())

	// c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	// if err != nil {
	// 	rlog.Errorln("dial:", err)
	// 	return nil, err
	// }

	Ws := &WebsocketClient{
		Conn: nil,
		Send: make(chan []byte),
	}
	if err := Ws.Reconnection(); err != nil {
		rlog.Fatalln("failed to connection to the ws,err:", err)
	}

	// defer c.Close()

	return Ws, nil
}

//Reconnection 重新获得ws连接
func (Wc *WebsocketClient) Reconnection() error {
	u := url.URL{Scheme: "wss", Host: host, Path: path}
	rlog.Infoln("connecting to:", u.String())
	var err error
	Wc.Conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		rlog.Errorln("dial:", err)
		return err
	}
	return nil

}

//MsgRead 读取ws的返回值
func (Wc *WebsocketClient) MsgRead() {
	defer func() {
		Wc.Conn.Close()
	}()

	for {
		msgType, message, err := Wc.Conn.ReadMessage()
		if err != nil {
			//重新连接，并且sub一下
			rlog.Errorln("failed to read msg from ws,starting reconnecting ,err:", err)
			if err := Wc.Reconnection(); err != nil {
				rlog.Println("failed to reconnect to infura,", err)
				time.Sleep(5 * time.Second)
				continue
			}
			Wc.MsgWrite()

		}

		rlog.Infoln("Get Message msgtype:", msgType, " message:", string(message))
		if msgType == websocket.TextMessage {
			var evt PendingTxType
			err := json.Unmarshal(message, &evt)
			if err != nil {
				rlog.Errorf("websocket: failed to decode JSON sent from client %s\n", err)
				continue
			}

			//start开始看下handl需要多久
			start := time.Now()

			err = Wc.HandlerTxInfo(evt)
			if err != nil {
				rlog.Errorf("failed to handle txinfo err:", err)
				continue
			}
			//结束看handle需要多久时间
			duration := time.Since(start)
			rlog.Infoln("===== HandlerTxInfo ==== duration:", duration.Milliseconds())
		}
	}
}

//MsgWrite 读取ws的消息
func (Wc *WebsocketClient) MsgWrite() {
	//这里写入subscribe
	postBody, _ := json.Marshal(map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "eth_subscribe",
		"params":  []string{"newPendingTransactions"},
	})
	err := Wc.Conn.WriteMessage(websocket.TextMessage, postBody)
	if err != nil {
		rlog.Fatalln("write:", err)
		return
	}
}

//HandlerTxInfo 获取某一条的tx的信息
func (Wc *WebsocketClient) HandlerTxInfo(txhash PendingTxType) error {

	client, err := ethclient.Dial(globalProvider)
	if err != nil {
		rlog.Fatal(err)
	}

	//hardcode 用来测试
	txhash.Params.Result = "0x4011cbff10481d8e1e6ff1895b9a28e81737489b380f5b68079324aa10dce908"
	txinfo, ispending, err := client.TransactionByHash(context.Background(), common.HexToHash(txhash.Params.Result))
	//解析出来info的数据
	if err != nil {
		return err
	}
	rlog.Infoln("txinfo TO:", txinfo.To().String(), " Payload:", hex.EncodeToString(txinfo.Data()), " ispending:", ispending, " err:", err)

	if strings.EqualFold(txinfo.To().String(), "0x7a250d5630b4cf539739df2c5dacb4c659f2488d") {
		uniswapabi, err := abi.JSON(strings.NewReader(MainABI))
		if err != nil {
			rlog.Fatal(err)
		}

		// var fillEvent AddETHLiquidityType
		receivedMap := map[string]interface{}{}
		// result, err := uniswapabi.Unpack("addLiquidityETH", txinfo.Data()[4:])

		method, err := uniswapabi.MethodById(txinfo.Data()[:4])
		if err != nil {
			rlog.Println("failed to get method by id err:", err)
		} else if method.Name == "addLiquidityETH" {
			//我了哥去
			if err := method.Inputs.UnpackIntoMap(receivedMap, txinfo.Data()[4:]); err != nil {
				rlog.Errorln("failed to get method.Inputs.UnpackIntoMap err:", err)
			} else {
				//成功监控到
				rlog.Infoln("ammmmmmmmmmmmm,", receivedMap, " price:", txinfo.GasPrice(), " gas:", txinfo.Gas())
				TestTransfe(txinfo.GasPrice(), txinfo.Gas())
			}
		} else {
			rlog.Println("current method is:", method.Name)
		}

		//只是abi的output部分
		// err = uniswapabi.UnpackIntoMap(receivedMap, "addLiquidityETH", txinfo.Data()[4:])
		// if err != nil {
		// 	rlog.Errorln("unpack err:", err)
		// 	// rlog.Fatal(err)
		// } else {
		// 	// var test AddETHLiquidityType
		// 	// json.Marshal(result, &test)
		// 	rlog.Panicln("find the target is :", receivedMap)
		// }

	}
	return nil

}

//PendingTxType pending tx的结果结构
type PendingTxType struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription string `json:"subscription"`
		Result       string `json:"result"`
	}
}

//infuraRequestDataType 统一的调用格式
type infuraRequestDataType struct {
	ID      uint64   `json:"id"`
	JSONRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

//AddETHLiquidityType 解析tx input
type AddETHLiquidityType struct {
	TokenContract      common.Address
	AmountTokenDesired *big.Int
	AmountTokenMin     *big.Int
	AmountETHMin       *big.Int
	ToAddress          common.Address
	Deadline           *big.Int
}

func TestTransferAuth(gasprice *big.Int, gaslimit uint64) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA("f36a83b3c3e5b506145f267fff3b986e499a2730fbc08e6c08a09b160e87ad83")
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	rlog.Errorln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// 	return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }
	nonce, err := ForTokenClient.PendingNonceAt(context.Background(), common.HexToAddress("0x31480704F726cD60010b1FEF1dB1c2a4f7bDD67f"))
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	// gasPrice, err := ForTokenClient.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	// log.Fatal(err)
	// 	return nil, err
	// }
	//todo 需要后期研究一下为什么这个接口获取不到正确的数据？
	// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	To:   &oneSplitAddress,
	// 	Data: DataWrapperForGasEstimate(source, destination, big.NewInt(1), dis.ReturnAmount, dis.Distribution, big.NewInt(0)),
	// })

	// rlog.Infof("gaslimit  is from network is :", gasLimit)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	//todo 这里是要做为需要交易多少eth 或者erc20的代币来制定的。
	auth.Value, _ = StringToBigInt("1", 18) // in wei

	// auth.GasLimit = uint64(3000000) // in unit  也同样需要乘以一个倍率
	// auth.GasLimit = Tcon.GasLimit
	// //13倍的gas费用 gas price调整为主网的130%
	// adjustPrice := gasPrice
	// if Tcon.GasPriceTimes > 0 {

	// 	PriceTimes := big.NewInt(0).Add(big.NewInt(100), big.NewInt(0).SetUint64(Tcon.GasPriceTimes))
	// 	adjustPriceTimes := big.NewInt(0).Mul(gasPrice, PriceTimes)
	// 	adjustPrice = big.NewInt(0).Div(adjustPriceTimes, big.NewInt(100))
	// 	rlog.Infoln("Gas times is:", Tcon.GasPriceTimes, " origin gas price:", gasPrice, " after adjust:", adjustPrice)
	// }
	// auth.GasPrice = adjustPrice
	// GasUsedSum := big.NewInt(0).Mul(big.NewInt(0).SetInt64(300000), gasPrice)

	// rlog.Infof("Get gas price:%s, gas limit:%d, gas total used:%d\n", auth.GasPrice, auth.GasLimit, GasUsedSum)
	auth.GasPrice = gasprice
	auth.GasLimit = gaslimit
	return auth, nil
}

func TestTransfe(gasprice *big.Int, gaslimit uint64) error {

	auth, err := TestTransferAuth(gasprice, gaslimit)
	if err != nil {
		rlog.Println("failed to create auth for transfer,err:", err)
		return err
	}

	distru := struct {
		ReturnAmount *big.Int
		Distribution []*big.Int
	}{}

	distru.ReturnAmount, _ = StringToBigInt("1000", 18)
	distru.Distribution = []*big.Int{big.NewInt(100), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)}

	//创建一个公用的函数
	// slipper := func(retValue *big.Int) *big.Int {
	// 	mul97 := big.NewInt(0).Mul(retValue, big.NewInt(0).Sub(big.NewInt(100), big.NewInt(0).SetUint64(50)))

	// 	curSlipper := big.NewInt(0).Div(mul97, big.NewInt(100))
	// 	rlog.Infof("PairName:%v, current slipper is:%v\n", "test", retValue)
	// 	return curSlipper
	// }
	minreserve, _ := StringToBigInt("1", 18)
	//																		rinkeby DAI
	tx, err := OneInchInstance.Swap(auth, common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"), common.HexToAddress("0x7ca4408137eb639570f8e647d9bd7b7e8717514a"), minreserve, big.NewInt(0), distru.Distribution, big.NewInt(0))
	if err != nil {
		rlog.Infof("got error:", err, " tx:", tx.Hash().String())
		return err
	}

	rlog.Infoln("Swap the eth to token successful tx:", tx.Hash().String())
	rlog.Panic()
	return nil
}

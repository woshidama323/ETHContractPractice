package main

import (
	"flag"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	wsmainnet   = "https://rinkeby.infura.io/v3/6707728235da4599b4a045c1b40ff0d9"
	wsrinkeby   = "wss://rinkeby.infura.io/ws/v3/6707728235da4599b4a045c1b40ff0d9"
	httprinkeby = "https://rinkeby.infura.io/v3/6707728235da4599b4a045c1b40ff0d9"
	httpmainnet = "https://mainnet.infura.io/v3/6707728235da4599b4a045c1b40ff0d9"
	ganachehttp = "http://127.0.0.1:8545"

	//eth address
	eth = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
	dai = "0x6b175474e89094c44da98b954eedeac495271d0f"

	//OneSplitMainnetAddress onesplit合约地址
	OneSplitMainnetAddress = "0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E"
)

//TInfos 全局变量 token的具体信息
var TInfos TokenInfos

//OneInchInstance 的实例
var OneInchInstance *Onesplitaudit

func main() {
	rlog.Info("start robot server....")
	// transferEvent()

	netmode := flag.String("netmode", "mainnet", "which net you want to use")
	flag.Parse()
	if netmode == nil {
		rlog.Fatal("please input a right network:%s\n", *netmode)
	}

	client, err := ethclient.Dial(*netmode)
	if err != nil {
		log.Fatal(err)
	}

	ForTokenClient = client

	//onesplitaddress 的合约地址
	oneSplitAddress := common.HexToAddress(OneSplitMainnetAddress)
	instance, err := NewOnesplitaudit(oneSplitAddress, client)
	if err != nil {
		rlog.Fatal(err)
	}

	OneInchInstance = instance

	if err := TInfos.LoadConfig(); err != nil {
		rlog.Info("failed to load config from default file:", err)
		return
	}

	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				rlog.Info("it's time to go to checking")
				//变量当前的配置文件中的token信息
				//需要的信息有：
				//1. source的地址   就是本金合约地址
				//2. 目标的地址      要得到的代币合约地址
				//3. 要交换的本金数量
				//4. 参与的dex的数量
				//5. flag的参数是什么 这里是0
				for index, value := range TInfos.TConfig {

					if value.Status == "disable" {
						rlog.Warn("Pair:%s has been disabled\n", value.Name)
						time.Sleep(1 * time.Second)
						continue
					}
					// value.MinReserverEthAmount()
					result, err := value.PriceMonitor()
					if err != nil {
						rlog.Info("Failed to monitor price for pair:", value.Name, " ,err:", err)
						continue
					}

					//检查当前是否有tx还处在pending状态
					//test
					// value.TxInfos = append(value.TxInfos, "0xab21585e21b45f65d758f7b0493333706e20cd44a7b2badb812817f43f89d299")
					//end test
					if status, pending, err := value.CheckTxStatus(); err != nil {
						rlog.Info("==%s\tCheckTxStatus\terr:%v\n", value.Name, err)
					} else if pending {
						continue
					} else if status == "disable" {
						TInfos.TConfig[index].Status = "disable"
						continue
					}

					value.Test()

					// value.MinReserverSAmount()

					//判断价格是否为0
					if result.ReturnAmount.Cmp(big.NewInt(0)) <= 0 {
						rlog.Warn("Got zero price for pair:%v\n", value.Name)
						continue
					}
					rlog.Info("................start Action................")
					value.Action(result)

					TInfos.TConfig[index].Status = value.Status
					TInfos.TConfig[index].TxInfos = value.TxInfos
					//先查余额
					//扣除必用项之后，还剩多少余额，然后全部交易掉
					//利用剩余余额计算可以收获多少对应的代币  进一步计算当前的价格
					//利用以上信息，当前价格等决定交易策略
					//进行交易

				}

			}

		}
	}()

	ChangeChannel := make(chan string, 1)
	ResponseChannel := make(chan string, 1)
	go GrpcServer(ChangeChannel, ResponseChannel)

	for {
		select {
		case task := <-ChangeChannel:
			if task == "updateconfig" {
				if err := TInfos.LoadConfig(); err != nil {
					rlog.Error("failed to load config from default file:", err)
					ResponseChannel <- err.Error()
					continue
				}
				ResponseChannel <- "success"
			} else if task == "approve" {
				for _, token := range TInfos.TConfig {
					txs, errs := token.ApproveForOneSplitAudit(token.SourceAddress, token.PrecisionSource)
					if errs != nil {
						ResponseChannel <- errs.Error()
						continue
					}

					txd, errd := token.ApproveForOneSplitAudit(token.Destination, token.PrecisionDestination)
					if errd != nil {
						ResponseChannel <- errd.Error()
						continue
					}

					txlist := "sourcetoken:" + txs + ",destinationtoken:" + txd
					ResponseChannel <- txlist

				}
			}

		}
	}

}

package main

import (
	"flag"
	"fmt"
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

	//onesplit合约地址
	OneSplitMainnetAddress = "0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E"
)

//全局变量
var TInfos TokenInfos

func main() {
	fmt.Println("start robot server....")
	// transferEvent()

	netmode := flag.String("netmode", "mainnet", "which net you want to use")
	flag.Parse()
	if netmode == nil {
		log.Fatalf("please input a right network:%s\n", *netmode)
	}

	client, err := ethclient.Dial(*netmode)
	if err != nil {
		log.Fatal(err)
	}

	//onesplitaddress 的合约地址
	oneSplitAddress := common.HexToAddress(OneSplitMainnetAddress)
	instance, err := NewOnesplitaudit(oneSplitAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	if err := TInfos.LoadConfig(); err != nil {
		fmt.Println("failed to load config from default file:", err)
		return
	}

	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("it's time to go to checking")
				//变量当前的配置文件中的token信息
				//需要的信息有：
				//1. source的地址   就是本金合约地址
				//2. 目标的地址      要得到的代币合约地址
				//3. 要交换的本金数量
				//4. 参与的dex的数量
				//5. flag的参数是什么 这里是0
				for _, value := range TInfos.TConfig {

					distri, err := instance.GetExpectedReturn(nil, common.HexToAddress(value.SourceAddress), common.HexToAddress(value.Destination), big.NewInt(1), big.NewInt(100), big.NewInt(0))
					if err != nil {
						fmt.Println("what's problem, err:", err)
						continue
					}

					fmt.Println("current token pair :", value.Name, " distribution is:", distri)
					if action, err := value.CheckStrategy(DistributionValue{
						distri.ReturnAmount,
						distri.Distribution,
					}); err != nil {
						fmt.Println("failed to check the strategy,err :", err)
					} else {
						fmt.Println("what have i get..:", action)
						//如果达到要求，那么就开始进行交易
						if action == "sell" {
							//eth 换成 dai
							if err := SwapPrepare(client, value.SourceAddress, value.Destination, value.TradeAddress, value.TradeAddressPriv, distri); err != nil {
								fmt.Println("failed to sell eth to erc20token err:", err)

							}
							// return
						} else if action == "buy" {
							//dai 换成 eth
							distribuy, err := instance.GetExpectedReturn(nil, common.HexToAddress(value.SourceAddress), common.HexToAddress(value.Destination), big.NewInt(1), big.NewInt(100), big.NewInt(0))
							if err != nil {
								fmt.Println("what's problem, err:", err)
								time.Sleep(1 * time.Second)
								continue
							}
							if err := SwapPrepare(client, value.Destination, value.SourceAddress, value.TradeAddress, value.TradeAddressPriv, distribuy); err != nil {
								fmt.Println("failed to sell erc20token to eth  err:", err)
							}
						}

					}

				}

			}

		}
	}()

	ChangeChannel := make(chan bool, 1)
	go GrpcServer(ChangeChannel)

	for {
		select {
		case <-ChangeChannel:
			if err := TInfos.LoadConfig(); err != nil {
				fmt.Println("failed to load config from default file:", err)
				continue
			}
		}
	}

}

//CheckStrategy 检查当前的价格条件是否满足要求
func CheckStrategy(dis DistributionValue) (bool, error) {
	//1. 价格是否大于一定的数值
	//比如当前设定为 eth-dai 800
	if dis.ReturnAmount.Cmp(big.NewInt(100)) >= 0 {
		fmt.Println("get the right price for swaping")
		return true, nil
	}
	return false, nil
}

//配置文件更新方法

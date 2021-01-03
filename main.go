package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

func main() {
	fmt.Println("start tasting code....")
	// transferEvent()
	client, err := ethclient.Dial(ganachehttp)
	if err != nil {
		log.Fatal(err)
	}

	//onesplitaddress 的合约地址
	oneSplitAddress := common.HexToAddress(OneSplitMainnetAddress)
	instance, err := NewOnesplitaudit(oneSplitAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				//变量当前的配置文件中的token信息
				//需要的信息有：
				//1. source的地址   就是本金合约地址
				//2. 目标的地址      要得到的代币合约地址
				//3. 要交换的本金数量
				//4. 参与的dex的数量
				//5. flag的参数是什么 这里是0

				distribution, err := instance.GetExpectedReturn(nil, common.HexToAddress(eth), common.HexToAddress(dai), big.NewInt(1), big.NewInt(100), big.NewInt(0))
				if err != nil {
					fmt.Println("what's problem, err:", err)
				}
				fmt.Println("distribution is:", distribution)
			}

		}
	}()

	//获取到当前到交易方案，这个时候，如果满足一定到条件

	ok, err := CheckStrategy(DistributionValue{
		distribution.ReturnAmount,
		distribution.Distribution,
	})

	//状态评估之后的结果，
	//如果评估成功，需要进行交易，则走下一步
	privateKey, err := crypto.HexToECDSA("f36a83b3c3e5b506145f267fff3b986e499a2730fbc08e6c08a09b160e87ad83")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	if ok {
		tx, err := instance.Swap(auth, common.HexToAddress(eth), common.HexToAddress(dai), big.NewInt(1), big.NewInt(700), distribution.Distribution, big.NewInt(0))
		if err != nil {
			fmt.Println("got error:", err)
		} else {
			fmt.Println("tx is:", tx)
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

//监控的方法

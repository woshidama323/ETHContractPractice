package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"os/user"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	defaultConfigFile = ".1inch/config.json"
)

//初始化为0.1eth
var initETHBalanceForToken = new(big.Float).SetFloat64(math.Pow10(17))

// var gain, _ = initETHBalance.Int(ethaccountminbalance)

//ForTokenClient 全局client 供token的操作用
var ForTokenClient *ethclient.Client

//TokenConfig 交易对信息
type TokenConfig struct {
	Name          string `jons:"name"`
	SourceAddress string `json:"sourceaddress"`
	Destination   string `json:"destinationaddress"`
	Upper         uint64 `json:"upper"` //价格波动的上限
	Lower         uint64 `json:"lower"` //价格波动的下限

	//交易的账户 该账户拥有或者打算拥有该交易对的币种
	TradeAddress     string `json:"tradeaddress"`
	TradeAddressPriv string `json:"tradeaddresspriv"`
}

//TradeAddrBalanceForSource tradeAddress对于source币种的余额
func (Tcon *TokenConfig) TradeAddrBalanceForSource() (*big.Int, error) {
	//计算合适的用于交易的erc20的数量
	sourceInstance, err := NewErc20token(common.HexToAddress(Tcon.SourceAddress), ForTokenClient)
	if err != nil {
		fmt.Println("failed to get instance of erc20 token err:", err)
		return nil, errors.New("account erc20 balance not enough for handling")
	}
	TradeBalanceToSource, err := sourceInstance.BalanceOf(nil, common.HexToAddress(Tcon.TradeAddress))
	if err != nil {
		fmt.Println("failed to get balance of erc20 token err:", err)
		return nil, errors.New("failed to get balance of erc20 token")
	}

	if TradeBalanceToSource.Cmp(big.NewInt(0)) <= 0 {
		return nil, errors.New("erc20 token balance is zero")
	}

	return TradeBalanceToSource, nil
}

//TradeAddrBalanceForDestination tradeAddress对于destination币种的余额
func (Tcon *TokenConfig) TradeAddrBalanceForDestination() (*big.Int, error) {
	//计算合适的用于交易的erc20的数量
	destinationInstance, err := NewErc20token(common.HexToAddress(Tcon.Destination), ForTokenClient)
	if err != nil {
		fmt.Println("failed to get instance of erc20 token err:", err)
		return nil, errors.New("account erc20 balance not enough for handling")
	}
	TradeBalanceToDestination, err := destinationInstance.BalanceOf(nil, common.HexToAddress(Tcon.TradeAddress))
	if err != nil {
		fmt.Println("failed to get balance of erc20 token err:", err)
		return nil, errors.New("failed to get balance of erc20 token")
	}

	if TradeBalanceToDestination.Cmp(big.NewInt(0)) <= 0 {
		return nil, errors.New("erc20 token balance is zero")
	}

	return TradeBalanceToDestination, nil
}

//TradeAddrBalanceForETH tradeAddress对于eth的余额
func (Tcon *TokenConfig) TradeAddrBalanceForETH() (*big.Int, error) {

	balance, err := ForTokenClient.BalanceAt(context.Background(), common.HexToAddress(Tcon.TradeAddress), nil)
	if err != nil {
		fmt.Println("failed to get the balance of account : ", Tcon.TradeAddress, " balance:", balance, " err:", err)
		return nil, err
	}

	return balance, nil
}

//CheckStrategy 用于判断当前是否达到要求
func (Tcon *TokenConfig) CheckStrategy(dis DistributionValue) (string, error) {
	//1. 价格是否大于一定的数值
	//returnamount的值会大于 max 这个时候卖出
	if dis.ReturnAmount.Cmp(new(big.Int).SetUint64(Tcon.Upper)) >= 0 {
		fmt.Println("get the right price for swaping")
		return "sell", nil
	} else if dis.ReturnAmount.Cmp(new(big.Int).SetUint64(Tcon.Lower)) <= 0 {
		return "buy", nil
	}
	return "", nil
}

//GetExpectedReturnForTokens 封装一下oneinch的getexpectedreturn
func (Tcon *TokenConfig) GetExpectedReturnForTokens() (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {

	//当前的交易对是属于哪一种？
	//这两种币种的余额 from账户中哪一个有  两个中 取多的，然后作为
	//正常情况下，erc20的代币，某一方一定是全部都是0 所以这里呢，如果都不是eth的交易对，那么可以用这个作为依据

	//eth余额获取
	ethBalance, err := Tcon.TradeAddrBalanceForETH()
	if err != nil {
		fmt.Printf("Failed to get eth balance from netowrk for account:[%s],err:", Tcon.TradeAddress, err)
		return struct {
			ReturnAmount *big.Int
			Distribution []*big.Int
		}{}, err
	}

	//erc20 sourcetoken 获取
	// sourceBalance, err := Tcon.TradeAddrBalanceForSource()
	_, err = Tcon.TradeAddrBalanceForSource()
	if err != nil {
		fmt.Printf("failed to get erc20 balance for source token:[%s] for account:[%s]", Tcon.SourceAddress, Tcon.TradeAddress)
		return struct {
			ReturnAmount *big.Int
			Distribution []*big.Int
		}{}, err
	}

	//erc20 destinationtoken获取
	// destinationBalance, err := Tcon.TradeAddrBalanceForDestination()
	_, err = Tcon.TradeAddrBalanceForDestination()
	if err != nil {
		fmt.Printf("failed to get erc20 balance for source token:[%s] for account:[%s]", Tcon.SourceAddress, Tcon.TradeAddress)
		return struct {
			ReturnAmount *big.Int
			Distribution []*big.Int
		}{}, err
	}

	initAmount := big.NewInt(0)
	initETHBalanceForToken.Int(initAmount)
	if strings.EqualFold(Tcon.SourceAddress, "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {

		//eth余额 大于 初始的量
		//eth余额大于最小额度，就认为当前的潜在交易对象是从这个方向开始的，
		if ethBalance.Cmp(initAmount) > 0 {

			canTradeEth := big.NewInt(0).Sub(ethBalance, initAmount)
			distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), canTradeEth, big.NewInt(100), big.NewInt(0))
			if err != nil {
				fmt.Println("failed to GetExpectedReturn from 1inch, err:", err)
				return struct {
					ReturnAmount *big.Int
					Distribution []*big.Int
				}{}, err
			}

			return distri, nil
		}
		return struct {
			ReturnAmount *big.Int
			Distribution []*big.Int
		}{}, errors.New("ethBalance is smaller than initAmount")
	}

	if strings.EqualFold(Tcon.Destination, "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
		//eth余额 大于 初始的量
		//eth余额大于最小额度，就认为当前的潜在交易对象是从这个方向开始的，
		if ethBalance.Cmp(initAmount) > 0 {

			canTradeEth := big.NewInt(0).Sub(ethBalance, initAmount)
			//
			distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), canTradeEth, big.NewInt(100), big.NewInt(0))
			if err != nil {
				fmt.Println("failed to GetExpectedReturn from 1inch, err:", err)
				return struct {
					ReturnAmount *big.Int
					Distribution []*big.Int
				}{}, err
			}

			return distri, nil
		}
		return struct {
			ReturnAmount *big.Int
			Distribution []*big.Int
		}{}, errors.New("ethBalance is smaller than initAmount")
	}

	// fmt.Println("current token pair :", value.Name, " distribution is:", distri)
	// if action, err := value.CheckStrategy(DistributionValue{
	// 	distri.ReturnAmount,
	// 	distri.Distribution,
	// }); err != nil {
	// 	fmt.Println("failed to check the strategy,err :", err)
	// } else {
	// 	fmt.Println("what have i get..:", action)
	// 	//如果达到要求，那么就开始进行交易
	// 	if action == "sell" {
	// 		//eth 换成 dai
	// 		if err := SwapPrepare(client, value.SourceAddress, value.Destination, value.TradeAddress, value.TradeAddressPriv, distri); err != nil {
	// 			fmt.Println("failed to sell eth to erc20token err:", err)

	// 		}
	// 		// return
	// 	} else if action == "buy" {
	// 		//dai 换成 eth
	// 		distribuy, err := instance.GetExpectedReturn(nil, common.HexToAddress(value.SourceAddress), common.HexToAddress(value.Destination), big.NewInt(1), big.NewInt(100), big.NewInt(0))
	// 		if err != nil {
	// 			fmt.Println("what's problem, err:", err)
	// 			time.Sleep(1 * time.Second)
	// 			continue
	// 		}
	// 		if err := SwapPrepare(client, value.Destination, value.SourceAddress, value.TradeAddress, value.TradeAddressPriv, distribuy); err != nil {
	// 			fmt.Println("failed to sell erc20token to eth  err:", err)
	// 		}
	// 	}

	// }
	// return nil
	return struct {
		ReturnAmount *big.Int
		Distribution []*big.Int
	}{}, nil
}

//PriceMonitor 设定交易对A<->B B作为基准, 但是两者B的金额相对于A较大，这样保证计算时为整数，比如B为ETH 1个B可以有1200个A
func (Tcon *TokenConfig) PriceMonitor() (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {

	distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), big.NewInt(1), big.NewInt(100), big.NewInt(0))
	if err != nil {
		fmt.Println("failed to GetExpectedReturn from 1inch, err:", err)
		return struct {
			ReturnAmount *big.Int
			Distribution []*big.Int
		}{}, err
	}
	return distri, nil
}

//Action 监控到当前价格之后 应该如何应对
func (Tcon *TokenConfig) Action(dis struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}) {
	//成功获得action之后啊，需要更具当前的价格，进行决策

	//1. 在 A <-> B 对应的前提下
	// 当前价格超过最大值，则
	if dis.ReturnAmount.Cmp(new(big.Int).SetUint64(Tcon.Upper)) >= 0 {
		fmt.Println("---Pair: [%s] Price: [%d] have been larger than Upper: [%d], starting swap from B -> A", Tcon.Name, dis.ReturnAmount, Tcon.Upper)
		// 检查各个币种的余额，如果满足条件则进行兑换
	} else if dis.ReturnAmount.Cmp(new(big.Int).SetUint64(Tcon.Lower)) <= 0 {
		// 当前价格低于最小值，则
		fmt.Println("---Pair [%s] Price: [%d] have been less than Lower [%d], starting swap from A -> B", Tcon.Name, dis.ReturnAmount, Tcon.Lower)

	} else {
		fmt.Println(".-.-.- Pair [%s] Price: [%d] is between Lower:[%d] and Upper:[%d], starting swap from A -> B", Tcon.Name, dis.ReturnAmount, Tcon.Lower, Tcon.Upper)

	}

}

//Action 监控到当前价格之后 应该如何应对
func (Tcon *TokenConfig) TokenSwap(strategy string) {

	if strategy == "fromAtoB" {
		//先重新获取到具体的数据，

	}
	tx, err := OneInchInstance.Swap(auth, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(destination), balancenogas, big.NewInt(1), dis.Distribution, big.NewInt(0))
	if err != nil {
		fmt.Println("got error:", err, " tx:", tx)
		return err
	}
}

// =======  总体token的策略 =========
//TokenInfos 用于解析tokens的数据
type TokenInfos struct {
	TConfig []TokenConfig `json:"tokens"`
}

//LoadConfig 将数据从config文件装载到本地服务中
func (TInfos *TokenInfos) LoadConfig() error {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr.HomeDir)
	configfile := usr.HomeDir + "/" + defaultConfigFile
	jsonFile, err := os.Open(configfile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var tokeninfos TokenInfos

	if err := json.Unmarshal(byteValue, &tokeninfos); err != nil {
		fmt.Println("failed to json the config from file: ", configfile, " err:", err)
		return err
	}
	TInfos.TConfig = tokeninfos.TConfig

	return nil
}

//eth-dai
//gas的费用大家都一样

//如何抽象的实现token的配置？
//1. 利用json数组来完成 token对的遍历

//增加json的marshal unmarshal的方法

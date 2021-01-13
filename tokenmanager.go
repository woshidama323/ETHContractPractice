package main

import (
	"context"
	"crypto/ecdsa"
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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	//保留交易地址的最小eth的量
	TradeAddressReserveEth string `json:"tradeaddressreserveeth"`

	//tradeaddress保留SourceAddress币种最小余额
	TradeAddressReserveS string `json:"tradeaddressreservesrc"`

	//tradeaddress保留SourceAddress币种最小余额
	TradeAddressReserveD string `json:"tradeaddressreservedest"`

	//增加币种有效数字
	PrecisionSource uint64 `json:"precisionsource"`

	//增加币种有效数字
	PrecisionDestination uint64 `json:"precisiondestination"`

	//test
	IamhereForTest *big.Int
}

//MinReserverSourceAmount tradeaddress保留最小的source币种余额
func (Tcon *TokenConfig) MinReserverSourceAmount() (*big.Int, error) {
	if minBalance, ok := big.NewInt(0).SetString(Tcon.TradeAddressReserveS, 10); ok {
		minBalance = big.NewInt(0).Mul(minBalance, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(Tcon.PrecisionSource)), big.NewInt(0)))
		fmt.Println("Current min Blance for TradeAddressReserveS: ", minBalance)
		return minBalance, nil
	}
	return nil, errors.New("failed to parse string to big int for TradeAddressReserveS")
}

//MinReserverDestinationAmount tradeaddress保留destination币种最小的余额
func (Tcon *TokenConfig) MinReserverDestinationAmount() (*big.Int, error) {
	if minBalance, ok := big.NewInt(0).SetString(Tcon.TradeAddressReserveD, 10); ok {
		minBalance = big.NewInt(0).Mul(minBalance, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(Tcon.PrecisionDestination)), big.NewInt(0)))
		fmt.Println("Current min Blance for TradeAddressReserveD: ", minBalance)
		return minBalance, nil
	}
	return nil, errors.New("failed to parse string to big int for TradeAddressReserveD")
}

//MinReserverEthAmount tradeaddress保留最小的eth的余额
func (Tcon *TokenConfig) MinReserverEthAmount() (*big.Int, error) {
	if minBalance, ok := big.NewInt(0).SetString(Tcon.TradeAddressReserveEth, 10); ok {
		minBalance = big.NewInt(0).Mul(minBalance, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0)))
		fmt.Println("Current min Blance for TradeAddressReserveEth: ", minBalance)
		return minBalance, nil
	}
	return nil, errors.New("failed to parse string to big int for TradeAddressReserveEth")
}

//StringToAmount tradeaddress保留最小的eth的余额
// func (Tcon *TokenConfig) StringToAmount(toconvert string) (*big.Int, error) {
// 	if minBalance, ok := big.NewInt(0).SetString(toconvert, 10); ok {
// 		minBalance = big.NewInt(0).Mul(minBalance, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0)))
// 		fmt.Println("Current min Blance for TradeAddressReserveEth: ", minBalance)
// 		return minBalance, nil
// 	}
// 	return nil, errors.New("failed to parse string to big int for TradeAddressReserveEth")
// }

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

//PriceMonitor 设定交易对A<->B B作为基准, 但是两者B的金额相对于A较大，这样保证计算时为整数，比如B为ETH 1个B可以有1200个A
func (Tcon *TokenConfig) PriceMonitor() (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {

	GetPrice := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(Tcon.PrecisionDestination)), big.NewInt(0))
	distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), GetPrice, big.NewInt(100), big.NewInt(0))
	fmt.Printf("monitor current price is %d, distribution is [%d]", distri.ReturnAmount, distri.Distribution)
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
		fmt.Printf("---Pair: [%s] Price: [%d] have been larger than Upper: [%d], starting swap from B -> A\n", Tcon.Name, dis.ReturnAmount, Tcon.Upper)
		// 检查各个币种的余额，如果满足条件则进行兑换
		if err := Tcon.TokenSwap("fromBtoA"); err != nil {
			fmt.Println("Action Got error:", err)
		}
	} else if dis.ReturnAmount.Cmp(new(big.Int).SetUint64(Tcon.Lower)) <= 0 {
		// 当前价格低于最小值，则
		fmt.Printf("---Pair [%s] Price: [%d] have been less than Lower [%d], starting swap from A -> B\n", Tcon.Name, dis.ReturnAmount, Tcon.Lower)
		if err := Tcon.TokenSwap("fromAtoB"); err != nil {
			fmt.Println("Action Got error:", err)
		}
	} else {
		//不交易
		fmt.Printf(".-.-.- Pair [%s] Price: [%d] is between Lower:[%d] and Upper:[%d], starting swap from A -> B\n", Tcon.Name, dis.ReturnAmount, Tcon.Lower, Tcon.Upper)

	}

}

//TokenSwap 监控到当前价格之后 应该如何应对
func (Tcon *TokenConfig) TokenSwap(strategy string) error {
	var tx *types.Transaction
	var err error
	auth, err := Tcon.BuildAuth()
	if err != nil {
		return err
	}

	//创建一个公用的函数
	slipper := func(retValue *big.Int) *big.Int {
		mul97 := big.NewInt(0).Mul(retValue, big.NewInt(97))
		return big.NewInt(0).Div(mul97, big.NewInt(100))
	}
	if strategy == "fromAtoB" {
		//先重新获取到具体的数据，
		//当明确知道哪一个交易方向的时候,才去看看余额情况，如果这个方向上的source余额为空则拒绝交易
		// var sBanlance *big.Int

		initAmount := big.NewInt(0)
		initETHBalanceForToken.Int(initAmount)

		if strings.EqualFold(Tcon.SourceAddress, "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
			//source是eth
			ethBalance, err := Tcon.TradeAddrBalanceForETH()
			if err != nil {
				return err
			}
			//这个时候应该从eth 兑换成其他的币种，这个时候，

			if ethBalance.Cmp(initAmount) > 0 {
				canTradeEth := big.NewInt(0).Sub(ethBalance, initAmount)
				distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), canTradeEth, big.NewInt(100), big.NewInt(0))
				if err != nil {
					fmt.Println("failed to GetExpectedReturn from 1inch, err:", err)
					return err
				}
				tx, err = OneInchInstance.Swap(auth, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), canTradeEth, slipper(distri.ReturnAmount), distri.Distribution, big.NewInt(0))
				if err != nil {
					fmt.Println("got error:", err, " tx:", tx)
					return err
				}
			} else {
				fmt.Printf("account :%s has no enough eth for swap balance:%d", Tcon.TradeAddress, ethBalance)
				return errors.New("")
			}

		} else {
			sBanlanceErc20, err := Tcon.TradeAddrBalanceForSource()
			if err != nil {
				fmt.Println("erc20 token balance is zero for address:", Tcon.TradeAddress, " current token address:", Tcon.SourceAddress)
				return err
			}
			if sBanlanceErc20.Cmp(big.NewInt(0)) > 0 {
				//最小的return效果可以用滑点的方法
				distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), sBanlanceErc20, big.NewInt(1), big.NewInt(0))
				if err != nil {
					fmt.Println("failed to GetExpectedReturn from 1inch, err:", err)
					return err
				}
				tx, err = OneInchInstance.Swap(auth, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), sBanlanceErc20, slipper(distri.ReturnAmount), distri.Distribution, big.NewInt(0))
				if err != nil {
					fmt.Println("got error:", err, " tx:", tx)
					return err
				}
			} else {
				fmt.Printf("account :%s has no enough eth for swap balance:%s\n", Tcon.TradeAddress, sBanlanceErc20)
			}

		}

	} else if strategy == "fromBtoA" {
		//先重新获取到具体的数据，
		//当明确知道哪一个交易方向的时候,才去看看余额情况，如果这个方向上的source余额为空则拒绝交易
		// var sBanlance *big.Int

		initAmount := big.NewInt(0)
		initETHBalanceForToken.Int(initAmount)

		if strings.EqualFold(Tcon.Destination, "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
			//source是eth
			dBalanceEth, err := Tcon.TradeAddrBalanceForETH()
			if err != nil {
				return err
			}
			//这个时候应该从eth 兑换成其他的币种，这个时候，

			if dBalanceEth.Cmp(initAmount) > 0 {
				canTradeEth := big.NewInt(0).Sub(dBalanceEth, initAmount)
				fmt.Println("++++++ current sub initamount:", canTradeEth)
				canTradeEth = big.NewInt(0).Sub(canTradeEth, big.NewInt(0).Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit))))
				fmt.Println("++++++ current sub Gas:", canTradeEth)

				if canTradeEth.Cmp(big.NewInt(0)) <= 0 {
					fmt.Println("Current balance for address:", Tcon.TradeAddress, " is not enough for trading,canTradeEth:", canTradeEth)
					return errors.New("is not enough for trading")
				}

				distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), canTradeEth, big.NewInt(100), big.NewInt(0))
				if err != nil {
					fmt.Println("failed to GetExpectedReturn from 1inch, err:", err)
					return err
				}

				auth.Value = canTradeEth //big.NewInt(1) //

				fmt.Println("........... distribution between swap", distri)

				tx, err = OneInchInstance.Swap(auth, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), canTradeEth, slipper(distri.ReturnAmount), distri.Distribution, big.NewInt(0))
				if err != nil {
					fmt.Println("got error:", err, " tx:", tx)
					return err
				}
				fmt.Println("*************** tx:", tx.Hash().String())
			} else {
				fmt.Printf("account :%s has no enough eth for swap balance:%s", Tcon.TradeAddress, dBalanceEth)
				return errors.New("")
			}

		} else {
			dBalanceErc20, err := Tcon.TradeAddrBalanceForDestination()
			if err != nil {
				return err
			}
			if dBalanceErc20.Cmp(big.NewInt(0)) > 0 {
				//最小的return效果可以用滑点的方法
				distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), dBalanceErc20, big.NewInt(1), big.NewInt(0))
				if err != nil {
					fmt.Println("failed to GetExpectedReturn from 1inch, err:", err)
					return err
				}
				tx, err = OneInchInstance.Swap(auth, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), dBalanceErc20, slipper(distri.ReturnAmount), distri.Distribution, big.NewInt(0))
				if err != nil {
					fmt.Println("got error:", err, " tx:", tx)
					return err
				}
			} else {
				fmt.Printf("account :%s has no enough eth for swap balance:%d\n", Tcon.TradeAddress, dBalanceErc20)
			}

		}
	}

	fmt.Println("*************** tx:", tx.Hash().String())
	return nil
}

//BuildAuth 创建auth为 合约操作
func (Tcon *TokenConfig) BuildAuth() (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(Tcon.TradeAddressPriv)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Printf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fmt.Println("Is this the same address provided from account:", crypto.PubkeyToAddress(*publicKeyECDSA), " == ", Tcon.TradeAddress)
	nonce, err := ForTokenClient.PendingNonceAt(context.Background(), common.HexToAddress(Tcon.TradeAddress))
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	gasPrice, err := ForTokenClient.SuggestGasPrice(context.Background())
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	//todo 需要后期研究一下为什么这个接口获取不到正确的数据？
	// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	To:   &oneSplitAddress,
	// 	Data: DataWrapperForGasEstimate(source, destination, big.NewInt(1), dis.ReturnAmount, dis.Distribution, big.NewInt(0)),
	// })

	// fmt.Println("gaslimit  is from network is :", gasLimit)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	//todo 这里是要做为需要交易多少eth 或者erc20的代币来制定的。
	// auth.Value = big.NewInt(1000001) // in wei
	auth.GasLimit = uint64(3000000) // in unit  也同样需要乘以一个倍率

	//13倍的gas费用 gas price调整为主网的130%
	GasPrice13 := big.NewInt(0).Mul(gasPrice, big.NewInt(13))     //是否需要乘上一个倍率
	auth.GasPrice = big.NewInt(0).Div(GasPrice13, big.NewInt(10)) //是否需要乘上一个倍率

	GasUsedSum := big.NewInt(0).Mul(big.NewInt(0).SetInt64(300000), gasPrice)

	fmt.Printf("Get gas price:%s, gas limit:%d, gas totle used:%d\n", auth.GasPrice, auth.GasLimit, GasUsedSum)
	return auth, nil
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

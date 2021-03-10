package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
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
	Upper         string `json:"upper"` //价格波动的上限
	Lower         string `json:"lower"` //价格波动的下限

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

	//保存每个交易对 交易对为key value为交易列表（key 为tx，value为当前状态）
	TxInfos []string

	//如果发生过交易失败，则标记为disable
	Status string `json:"status"`

	//slipper
	Slipper uint64 `json:"slipper"`

	//gaslimit
	GasLimit uint64 `json:"gaslimit"`

	//gasprice倍率
	GasPriceTimes uint64 `json:"gaspricetimes"`
}

//StringToBigInt tradeaddress保留最小的source币种余额
func (Tcon *TokenConfig) StringToBigInt(reserve string, precision uint64) (*big.Int, error) {

	unit := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(precision)), big.NewInt(0))
	if minBalance, ok := big.NewFloat(0).SetString(reserve); ok {
		minBalance = big.NewFloat(0).Mul(minBalance, big.NewFloat(0).SetInt(unit))

		minBalance.Int(unit)
		return unit, nil
	}
	return nil, errors.New("failed to parse string to big int")
}

//MinReserverDestinationAmount tradeaddress保留destination币种最小的余额
// func (Tcon *TokenConfig) MinReserverDestinationAmount() (*big.Int, error) {
// 	if minBalance, ok := big.NewInt(0).SetString(Tcon.TradeAddressReserveD, 10); ok {
// 		minBalance = big.NewInt(0).Mul(minBalance, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(Tcon.PrecisionDestination)), big.NewInt(0)))
// 		rlog.Info("Current min Blance for TradeAddressReserveD: ", minBalance)
// 		return minBalance, nil
// 	}
// 	return nil, errors.New("failed to parse string to big int for TradeAddressReserveD")
// }

// //MinReserverEthAmount tradeaddress保留最小的eth的余额
// func (Tcon *TokenConfig) MinReserverEthAmount() (*big.Int, error) {
// 	if minBalance, ok := big.NewInt(0).SetString(Tcon.TradeAddressReserveEth, 10); ok {
// 		minBalance = big.NewInt(0).Mul(minBalance, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0)))
// 		rlog.Info("Current min Blance for TradeAddressReserveEth: ", minBalance)
// 		return minBalance, nil
// 	}
// 	return nil, errors.New("failed to parse string to big int for TradeAddressReserveEth")
// }

//StringToAmount tradeaddress保留最小的eth的余额
// func (Tcon *TokenConfig) StringToAmount(toconvert string) (*big.Int, error) {
// 	if minBalance, ok := big.NewInt(0).SetString(toconvert, 10); ok {
// 		minBalance = big.NewInt(0).Mul(minBalance, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0)))
// 		rlog.Info("Current min Blance for TradeAddressReserveEth: ", minBalance)
// 		return minBalance, nil
// 	}
// 	return nil, errors.New("failed to parse string to big int for TradeAddressReserveEth")
// }

//TradeAddrBalanceForSource tradeAddress对于source币种的余额
func (Tcon *TokenConfig) TradeAddrBalanceForSource() (*big.Int, error) {
	//计算合适的用于交易的erc20的数量
	sourceInstance, err := NewErc20token(common.HexToAddress(Tcon.SourceAddress), ForTokenClient)
	if err != nil {
		rlog.Error("failed to get instance of erc20 token err:", err)
		return nil, errors.New("account erc20 balance not enough for handling")
	}
	TradeBalanceToSource, err := sourceInstance.BalanceOf(nil, common.HexToAddress(Tcon.TradeAddress))
	if err != nil {
		rlog.Error("failed to get balance of erc20 token err:", err)
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
		rlog.Error("failed to get instance of erc20 token err:", err)
		return nil, errors.New("account erc20 balance not enough for handling")
	}
	TradeBalanceToDestination, err := destinationInstance.BalanceOf(nil, common.HexToAddress(Tcon.TradeAddress))
	if err != nil {
		rlog.Error("failed to get balance of erc20 token err:", err)
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
		rlog.Error("failed to get the balance of account : ", Tcon.TradeAddress, " balance:", balance, " err:", err)
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
	rlog.Info("monitor current price is %d, distribution is [%d]\n", distri.ReturnAmount, distri.Distribution)
	if err != nil {
		rlog.Info("failed to GetExpectedReturn from 1inch, err:", err)
		return struct {
			ReturnAmount *big.Int
			Distribution []*big.Int
		}{}, err
	}
	if distri.ReturnAmount.Cmp(big.NewInt(0)) <= 0 {
		rlog.Info("name:%s,returnamount is zero", Tcon.Name)
		return struct {
			ReturnAmount *big.Int
			Distribution []*big.Int
		}{}, errors.New("returnamount is zero")
	}
	return distri, nil
}

//CheckTxStatus 检查自己的tx列表中是否有还处在pending的状态的
func (Tcon *TokenConfig) CheckTxStatus() (status string, isPending bool, err error) {

	for index, txinfomap := range Tcon.TxInfos {
		txHash := common.HexToHash(txinfomap)
		_, isPending, err = ForTokenClient.TransactionByHash(context.Background(), txHash)
		if err != nil {
			rlog.Error("== %s\tErr:%v\n", Tcon.Name, err)
			return "", false, err
		}
		if isPending {
			rlog.Info("== %s\tisPending:%v\n", Tcon.Name, txinfomap)
			return "", true, nil
		}

		//是否是错误的
		if receipt, err := ForTokenClient.TransactionReceipt(context.Background(), txHash); err != nil {
			rlog.Error("receipt got err:", err)
		} else {
			rlog.Info("*** receipt:", receipt)
			if receipt.Status == 0 {

				rlog.Info(" %v disable current pair\n", receipt)

				//移除当前的tx

				Tcon.TxInfos = append(Tcon.TxInfos[:index], Tcon.TxInfos[index+1:]...)
				return "disable", false, nil
			}
		}
	}
	//如果都没有pending的则直接删除掉
	Tcon.TxInfos = nil

	return "", false, nil

}

//Test 用于测试
func (Tcon *TokenConfig) Test() {

	txHash := common.HexToHash("0xab21585e21b45f65d758f7b0493333706e20cd44a7b2badb812817f43f89d299")
	_, isPending, err := ForTokenClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		rlog.Error("== %s\tErr:%v\n", Tcon.Name, err)
		return
	}
	if isPending {
		rlog.Info("== %s\tisPending:%v\n", Tcon.Name, txHash)
	}

	//是否是错误的
	if receipt, err := ForTokenClient.TransactionReceipt(context.Background(), txHash); err != nil {
		rlog.Error("receipt got err:", err)
	} else {
		rlog.Info("*** receipt:", receipt.Status)

	}

}

//Action 监控到当前价格之后 应该如何应对
func (Tcon *TokenConfig) Action(dis struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}) {
	//成功获得action之后啊，需要更具当前的价格，进行决策

	//1. 在 A <-> B 对应的前提下
	// 当前价格超过最大值，则
	var tx string
	var err error
	var UpperBigInt, LowerBigInt *big.Int

	// if strings.EqualFold(Tcon.Destination, eth) {
	// 	UpperBigInt, err = Tcon.StringToBigInt(Tcon.Upper, 18)
	// 	if err != nil {
	// 		rlog.Info("faile to convert upper to big int err:%v", err)
	// 		return
	// 	}

	// 	LowerBigInt, err = Tcon.StringToBigInt(Tcon.Lower, 18)
	// 	if err != nil {
	// 		rlog.Info("faile to convert upper to big int err:%v", err)
	// 		return
	// 	}
	// } else {
	UpperBigInt, err = Tcon.StringToBigInt(Tcon.Upper, Tcon.PrecisionDestination)
	if err != nil {
		rlog.Info("faile to convert upper to big int err:%v", err)
		return
	}

	LowerBigInt, err = Tcon.StringToBigInt(Tcon.Lower, Tcon.PrecisionDestination)
	if err != nil {
		rlog.Info("faile to convert upper to big int err:%v", err)
		return
	}
	rlog.Info("convert String to big int for upper:%v and lower:%v\n", UpperBigInt, LowerBigInt)
	// }

	if dis.ReturnAmount.Cmp(UpperBigInt) >= 0 {
		rlog.Info("---Pair: [%s] Price: [%s] have been larger than Upper: [%s], starting swap from B -> A\n", Tcon.Name, dis.ReturnAmount, UpperBigInt)
		// 检查各个币种的余额，如果满足条件则进行兑换
		tx, err = Tcon.TokenSwap("fromBtoA")
		if err != nil {
			rlog.Info("Action Got error:%v\n", err)
		}
	} else if dis.ReturnAmount.Cmp(LowerBigInt) <= 0 {
		// 当前价格低于最小值，则
		rlog.Info("---Pair [%s] Price: [%s] have been less than Lower [%s], starting swap from A -> B\n", Tcon.Name, dis.ReturnAmount, LowerBigInt)
		tx, err = Tcon.TokenSwap("fromAtoB")
		if err != nil {
			rlog.Info("Action Got error:%v\n", err)
		}
	} else {
		//不交易
		rlog.Info(".-.-.- Pair [%s] Price: [%s] is between Lower:[%s] and Upper:[%s], do nothing\n", Tcon.Name, dis.ReturnAmount, Tcon.Lower, Tcon.Upper)

	}
	if len(tx) != 0 {
		Tcon.TxInfos = append(Tcon.TxInfos, tx)
	}

}

//TokenSwap 监控到当前价格之后 应该如何应对
func (Tcon *TokenConfig) TokenSwap(strategy string) (string, error) {
	var tx *types.Transaction
	var err error
	auth, err := Tcon.BuildAuth()
	if err != nil {
		return "", err
	}

	//创建一个公用的函数
	slipper := func(retValue *big.Int) *big.Int {
		mul97 := big.NewInt(0).Mul(retValue, big.NewInt(0).Sub(big.NewInt(100), big.NewInt(0).SetUint64(Tcon.Slipper)))

		curSlipper := big.NewInt(0).Div(mul97, big.NewInt(100))
		rlog.Info("PairName:%v, current slipper is:%v\n", Tcon.Name, Tcon.Slipper)
		return curSlipper
	}

	ethmin, err := Tcon.StringToBigInt(Tcon.TradeAddressReserveEth, 18)
	if err != nil {
		return "", err
	}
	rlog.Info("Current min Blance of ETH for [%s] is [%v] \n", Tcon.Name, ethmin)
	if strategy == "fromAtoB" {
		//先重新获取到具体的数据，
		//当明确知道哪一个交易方向的时候,才去看看余额情况，如果这个方向上的source余额为空则拒绝交易
		// var sBanlance *big.Int

		if strings.EqualFold(Tcon.SourceAddress, "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
			//source是eth
			ethBalance, err := Tcon.TradeAddrBalanceForETH()
			if err != nil {
				return "", err
			}
			//这个时候应该从eth 兑换成其他的币种，这个时候，

			if ethBalance.Cmp(ethmin) > 0 {
				canTradeEth := big.NewInt(0).Sub(ethBalance, ethmin)
				distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), canTradeEth, big.NewInt(100), big.NewInt(0))
				if err != nil {
					rlog.Info("failed to GetExpectedReturn from 1inch, err:", err)
					return "", err
				}
				tx, err = OneInchInstance.Swap(auth, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), canTradeEth, slipper(distri.ReturnAmount), distri.Distribution, big.NewInt(0))
				if err != nil {
					rlog.Info("got error:", err, " tx:", tx)
					return "", err
				}
			} else {
				rlog.Info("account :%s has no enough eth for swap balance:%d\n", Tcon.TradeAddress, ethBalance)
				return "", errors.New("")
			}

		} else {
			sBanlanceErc20, err := Tcon.TradeAddrBalanceForSource()
			if err != nil {
				rlog.Info("erc20 token balance is zero for address:", Tcon.TradeAddress, " current token address:", Tcon.SourceAddress)
				return "", err
			}

			minReserveSource, err := Tcon.StringToBigInt(Tcon.TradeAddressReserveS, Tcon.PrecisionSource)
			if err != nil {
				rlog.Info("failed to get erc20 min token availabe in config file,tokenaddress:", Tcon.SourceAddress, " for address:", Tcon.TradeAddress, "ethsource:", minReserveSource)
				return "", err
			}
			rlog.Info("Current min Blance of source Reserve for [%s] is [%v] \n", Tcon.Name, minReserveSource)
			if sBanlanceErc20.Cmp(minReserveSource) > 0 {
				//最小的return效果可以用滑点的方法
				sBerc20Fortrade := big.NewInt(0).Sub(sBanlanceErc20, minReserveSource)
				distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), sBerc20Fortrade, big.NewInt(100), big.NewInt(0))
				if err != nil {
					rlog.Info("failed to GetExpectedReturn from 1inch, err:", err)
					return "", err
				}
				log.Println("+++distri:", distri)
				tx, err = OneInchInstance.Swap(auth, common.HexToAddress(Tcon.SourceAddress), common.HexToAddress(Tcon.Destination), sBerc20Fortrade, slipper(distri.ReturnAmount), distri.Distribution, big.NewInt(0))
				if err != nil {
					rlog.Info("got error:", err, " tx:", tx)
					Tcon.Status = "disable"
					return "", err
				}
			} else {
				rlog.Info("account :%s has no enough token for swap balance:%s\n", Tcon.TradeAddress, sBanlanceErc20)
				return "", errors.New("has no enough token for swap balance")
			}

		}

	} else if strategy == "fromBtoA" {
		//先重新获取到具体的数据，
		//当明确知道哪一个交易方向的时候,才去看看余额情况，如果这个方向上的source余额为空则拒绝交易
		// var sBanlance *big.Int

		// initAmount := big.NewInt(0)
		// initETHBalanceForToken.Int(initAmount)

		if strings.EqualFold(Tcon.Destination, "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
			//source是eth
			dBalanceEth, err := Tcon.TradeAddrBalanceForETH()
			if err != nil {
				return "", err
			}
			//这个时候应该从eth 兑换成其他的币种，这个时候，

			if dBalanceEth.Cmp(ethmin) > 0 {
				canTradeEth := big.NewInt(0).Sub(dBalanceEth, ethmin)
				rlog.Info("++++++ current sub initamount:", canTradeEth)
				canTradeEth = big.NewInt(0).Sub(canTradeEth, big.NewInt(0).Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit))))
				rlog.Info("++++++ current sub Gas:", canTradeEth)

				if canTradeEth.Cmp(big.NewInt(0)) <= 0 {
					rlog.Info("Current balance for address:", Tcon.TradeAddress, " is not enough for trading,canTradeEth:", canTradeEth)
					return "", errors.New("is not enough for trading")
				}

				distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), canTradeEth, big.NewInt(100), big.NewInt(0))
				if err != nil {
					rlog.Info("failed to GetExpectedReturn from 1inch, err:", err)
					return "", err
				}

				auth.Value = canTradeEth //big.NewInt(1) //

				rlog.Info("........... distribution between swap", distri)

				// data := DataWrapperForGasEstimate(Tcon.Destination, Tcon.SourceAddress, canTradeEth, slipper(distri.ReturnAmount), distri.Distribution, big.NewInt(0))
				// forpoint := common.HexToAddress(Tcon.SourceAddress)
				// gasLimittest, err := ForTokenClient.EstimateGas(context.Background(), ethereum.CallMsg{
				// 	To:   &forpoint,
				// 	Data: data,
				// })
				// rlog.Info("how to get gas....%v\n", gasLimittest)
				// if err != nil {
				// 	log.Fatal(err)
				// }
				tx, err = OneInchInstance.Swap(auth, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), canTradeEth, slipper(distri.ReturnAmount), distri.Distribution, big.NewInt(0))
				if err != nil {
					rlog.Info("got error:", err, " tx:", tx)
					return "", err
				}
			} else {
				rlog.Info("account :%s has no enough eth for swap balance:%s", Tcon.TradeAddress, dBalanceEth)
				return "", errors.New("has no enough eth for swap balance")
			}

		} else {
			dBalanceErc20, err := Tcon.TradeAddrBalanceForDestination()
			if err != nil {
				return "", err
			}
			minReserveDestination, err := Tcon.StringToBigInt(Tcon.TradeAddressReserveD, Tcon.PrecisionDestination)
			if err != nil {
				rlog.Info("failed to get erc20 min token availabe in config file,tokenaddress:", Tcon.SourceAddress, " for address:", Tcon.TradeAddress, "minReserveDestination:", minReserveDestination)
				return "", err
			}
			rlog.Info("Current min Blance of destination Reserve for [%s] is [%v] \n", Tcon.Name, minReserveDestination)

			if dBalanceErc20.Cmp(minReserveDestination) > 0 {
				//最小的return效果可以用滑点的方法
				subResult := big.NewInt(0).Sub(dBalanceErc20, minReserveDestination)
				distri, err := OneInchInstance.GetExpectedReturn(nil, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), subResult, big.NewInt(100), big.NewInt(0))
				if err != nil {
					rlog.Info("failed to GetExpectedReturn from 1inch, err:", err)
					return "", err
				}
				log.Println("+++distri:", distri) //slipper(distri.ReturnAmount)
				tx, err = OneInchInstance.Swap(auth, common.HexToAddress(Tcon.Destination), common.HexToAddress(Tcon.SourceAddress), subResult, slipper(distri.ReturnAmount), distri.Distribution, big.NewInt(0))
				if err != nil {
					rlog.Info("got error:", err, " tx:", tx)
					Tcon.Status = "disable"
					return "", err
				}
			} else {
				rlog.Info("account :%s has no enough tokens for swap balance:%d\n", Tcon.TradeAddress, dBalanceErc20)
				return "", errors.New("has no enough tokens for swap balance")
			}

		}
	}

	rlog.Info("*************** tx:", tx.Hash().String())
	Tcon.TxInfos = append(Tcon.TxInfos, tx.Hash().String())
	return tx.Hash().String(), nil
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

	rlog.Info("Is this the same address provided from account:", crypto.PubkeyToAddress(*publicKeyECDSA), " == ", Tcon.TradeAddress)
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

	// rlog.Info("gaslimit  is from network is :", gasLimit)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	//todo 这里是要做为需要交易多少eth 或者erc20的代币来制定的。
	// auth.Value = big.NewInt(1000001) // in wei
	// auth.GasLimit = uint64(3000000) // in unit  也同样需要乘以一个倍率
	auth.GasLimit = Tcon.GasLimit
	//13倍的gas费用 gas price调整为主网的130%
	adjustPrice := gasPrice
	if Tcon.GasPriceTimes > 0 {

		PriceTimes := big.NewInt(0).Add(big.NewInt(100), big.NewInt(0).SetUint64(Tcon.GasPriceTimes))
		adjustPriceTimes := big.NewInt(0).Mul(gasPrice, PriceTimes)
		adjustPrice = big.NewInt(0).Div(adjustPriceTimes, big.NewInt(100))
		rlog.Info("Gas times is:", Tcon.GasPriceTimes, " origin gas price:", gasPrice, " after adjust:", adjustPrice)
	}
	auth.GasPrice = adjustPrice
	GasUsedSum := big.NewInt(0).Mul(big.NewInt(0).SetInt64(300000), gasPrice)

	rlog.Info("Get gas price:%s, gas limit:%d, gas total used:%d\n", auth.GasPrice, auth.GasLimit, GasUsedSum)
	return auth, nil
}

//ApproveForOneSplitAudit approval 一个地址
func (Tcon *TokenConfig) ApproveForOneSplitAudit(source string, precision uint64) (string, error) {

	sourceInstance, err := NewErc20token(common.HexToAddress(source), ForTokenClient)
	if err != nil {
		rlog.Info("failed to get instance of erc20 token err:", err)
		return "", errors.New("failed to get instance of erc20 token")
	}

	auth, err := Tcon.BuildAuth()
	if err != nil {
		return "", err
	}

	prec := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(0).SetUint64(precision), big.NewInt(0))
	tx, err := sourceInstance.Approve(auth, common.HexToAddress(OneSplitMainnetAddress), big.NewInt(0).Mul(prec, big.NewInt(9223372036854775800)))
	if err != nil {
		rlog.Info("failed to Approve for token:", source, " to spender:", OneSplitMainnetAddress, " precision:", precision, " err:", err)
		return "", errors.New("failed to Approve for onesplitaudit")
	}
	if tx != nil {

		Tcon.TxInfos = append(Tcon.TxInfos, tx.Hash().String())
		return tx.Hash().String(), nil
	}
	return "", errors.New("tx is nil")

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
	rlog.Info(usr.HomeDir)
	configfile := usr.HomeDir + "/" + defaultConfigFile
	jsonFile, err := os.Open(configfile)
	if err != nil {
		rlog.Info(err)
		return err
	}
	rlog.Info("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var tokeninfos TokenInfos

	if err := json.Unmarshal(byteValue, &tokeninfos); err != nil {
		rlog.Info("failed to json the config from file: ", configfile, " err:", err)
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

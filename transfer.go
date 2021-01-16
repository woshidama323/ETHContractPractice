package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	// "github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

const (
	//approval value
	approval int64 = 1000000*10 ^ 18
)

//初始化为0.1eth
var initETHBalance = new(big.Float).SetFloat64(math.Pow10(17))

var ethaccountminbalance *big.Int // = big.NewInt(0).SetUint64(hello.Uint64()) // 0.1ETH

var gain, _ = initETHBalance.Int(ethaccountminbalance)

//SwapPrepare 用于交易
func SwapPrepare(client *ethclient.Client, source, destination, fromaddress, private string, dis DistributionValue) error {

	// hellof := new(big.Float).SetFloat64(math.Pow10(17))

	// var ethaccountmin *big.Int // = big.NewInt(0).SetUint64(hello.Uint64()) // 0.1ETH

	// gain, _ := hellof.Int(ethaccountmin)
	// fmt.Println(gain, "   --     ", ethaccountmin)
	//获取到当前到交易方案，这个时候，如果满足一定到条件
	//状态评估之后的结果，
	//如果评估成功，需要进行交易，则走下一步
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		log.Fatal(err)
	}

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }

	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(fromaddress))
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//onesplitaddress 的合约地址
	oneSplitAddress := common.HexToAddress(OneSplitMainnetAddress)
	instance, err := NewOnesplitaudit(oneSplitAddress, client)
	if err != nil {
		fmt.Println("failed to create instance for onesplit,err:", err)
		return err
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
	auth.GasLimit = uint64(300000) // in unit  也同样需要乘以一个倍率

	//13倍的gas费用 gas price调整为主网的130%
	GasPrice13 := big.NewInt(0).Mul(gasPrice, big.NewInt(13))     //是否需要乘上一个倍率
	auth.GasPrice = big.NewInt(0).Div(GasPrice13, big.NewInt(10)) //是否需要乘上一个倍率

	GasUsedSum := big.NewInt(0).Mul(big.NewInt(0).SetInt64(300000), gasPrice)
	account := common.HexToAddress(fromaddress)
	//当前操作账户的eth余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		fmt.Println("failed to get the balance of account : ", account, " balance:", balance, " err:", err)
		return err
	}

	var balancenogas *big.Int
	// var ercbalance *big.Int
	if strings.EqualFold(source, "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {

		//todo gain cmp
		if gain == nil {
			fmt.Println("reserve balance is zero")
			return errors.New("reserve balance is zero")
		}
		if balance.Cmp(gain) > 0 && balance.Cmp(GasUsedSum) > 0 {

			balancenogas = big.NewInt(0).Sub(balance, gain)
			fmt.Println("+++++ balancenogas", balancenogas, " ethaccountminbalance:", ethaccountminbalance)
		} else {
			return errors.New("eth balance not enough for handling")
		}
		auth.Value = balancenogas // in wei
	} else {
		// if balance.Cmp(gain) < 0 || balance.Cmp(GasUsedSum) < 0 {
		if balance.Cmp(GasUsedSum) < 0 {
			return errors.New("eth balance not enough for handling erc20")
		}
		//计算合适的用于交易的erc20的数量
		ercinstance, err := NewErc20token(common.HexToAddress(source), client)
		if err != nil {
			fmt.Println("failed to get instance of erc20 token err:", err)
			return errors.New("account erc20 balance not enough for handling")
		}
		balancenogas, err = ercinstance.BalanceOf(nil, account)
		if err != nil {
			fmt.Println("failed to get balance of erc20 token err:", err)
			return errors.New("failed to get balance of erc20 token")
		}

		if balancenogas.Cmp(big.NewInt(0)) <= 0 {
			return errors.New("erc20 token balance is zero")
		}

	}

	//balance - gasprice * gaslimit

	fmt.Println("account:", account, "  balance:", balance, " gasPrice:", gasPrice, " GasLimit:300000", " value:", 2000000, " GasUsedSum", GasUsedSum)

	// test := balancenogas.SetUint64(1000001)
	// auth.Value = balancenogas
	// test = balancenogas.SetUint64(1000000)
	tx, err := instance.Swap(auth, common.HexToAddress(source), common.HexToAddress(destination), balancenogas, big.NewInt(1), dis.Distribution, big.NewInt(0))
	if err != nil {
		fmt.Println("got error:", err, " tx:", tx)
		return err
	}

	fmt.Println("tx is:", tx.Hash().String())

	account = common.HexToAddress(fromaddress)
	balance, err = client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		fmt.Println("failed to get the balance of account : ", account, " balance:", balance, " err:", err)
	}
	fmt.Println("account:", account, "  balance:", balance)
	return nil
}

func approveForErc20(client *ethclient.Client, tokenContract, sourceAddress, toAuther, private string) {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		fmt.Println("got error when crypto err:", err)
	}

	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(sourceAddress))
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	//todo 这里是要做为需要交易多少eth 或者erc20的代币来制定的。
	// auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in unit  也同样需要乘以一个倍率
	auth.GasPrice = gasPrice       //是否需要乘上一个倍率
	ercinstance, err := NewErc20token(common.HexToAddress(tokenContract), client)
	if err != nil {
		fmt.Println("failed to get instance of erc20 token")
		return
	}

	tx, err := ercinstance.Approve(auth, common.HexToAddress(toAuther), big.NewInt(approval))
	if err != nil {
		fmt.Println("failed to send approval,err:", err)
		return
	}
	fmt.Println("tx is :", tx.Hash().String())

}

//DataWrapperForGasEstimate 评估gaslimit
func DataWrapperForGasEstimate(source, destination string, amount, percentage *big.Int, distri []*big.Int, flag *big.Int) []byte {
	// getExpectedReturnFnSignature := []byte("Swap(address,address,uint256,uint256,uint256[],uint256)")
	getExpectedReturnFnSignature := []byte("transfer(address,uint256)")
	// hash := sha3.NewKeccak256()
	hash := sha3.NewLegacyKeccak256()
	hash.Write(getExpectedReturnFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(common.HexToAddress(source).Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))
	amount1 := new(big.Int)
	amount1.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount1.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	// paddedSourceAddress := common.LeftPadBytes(common.HexToAddress(source).Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedSourceAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	// addedDestinationAddress := common.LeftPadBytes(common.HexToAddress(destination).Bytes(), 32)
	// fmt.Println(hexutil.Encode(addedDestinationAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	// paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	// paddedPercentage := common.LeftPadBytes(percentage.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedPercentage)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	// var fordistri []byte
	// for _, dis := range distri {
	// 	disbytes := common.LeftPadBytes(dis.Bytes(), 32)
	// 	fordistri = append(fordistri, disbytes...)
	// }
	// // paddedDist := common.LeftPadBytes(distri.Bytes(), 32)
	// // fmt.Println(hexutil.Encode(paddedPercentage)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	// paddedFlag := common.LeftPadBytes(flag.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedFlag)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	// var data []byte
	// data = append(data, methodID...)
	// data = append(data, paddedSourceAddress...)
	// data = append(data, addedDestinationAddress...)
	// data = append(data, paddedAmount...)
	// data = append(data, paddedPercentage...)
	// data = append(data, fordistri...)
	// data = append(data, paddedFlag...)
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	return data

}

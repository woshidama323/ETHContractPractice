package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
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
var hello = new(big.Float).SetFloat64(math.Pow10(17))

var ethaccountminbalance *big.Int // = big.NewInt(0).SetUint64(hello.Uint64()) // 0.1ETH

var gain, _ = hello.Int(ethaccountminbalance)

func transferEvent() {

	client, err := ethclient.Dial(ganachehttp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	// _ = client // we'll use this in the upcoming sections/

	//获取一个账户的余额
	account := common.HexToAddress("0xcDB24EAB8179C807E30b508A1B95D265162d9858")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		fmt.Println("failed to get the balance of account : ", account, " balance:", balance, " err:", err)
	}
	fmt.Println("account:", account, "  balance:", balance)

	//创建交易  inch账户的私钥字符串
	privateKey, err := crypto.HexToECDSA("0ce8111d23a0f272a56f7234c9eff1d2709563805676139fde71599c23bf4c6f")
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

	fmt.Println("from account nonce:", nonce, " address:", fromAddress.Hex())
	value := big.NewInt(10000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//inch 2 账户
	toAddress := common.HexToAddress("0x05a0291E474D3A3911B431bf9653cC6Fb0AC8b25")

	var data []byte
	//这里是eth的转账，所以data可以为空，然后 gaslimit需要设置，price可以根据主网的价格来 ，两个只要选择一个就够
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Println("raw tx hex..:", rawTxHex)

	//开始广播出去
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f

}

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
	getExpectedReturnFnSignature := []byte("Swap(address,address,uint256,uint256,uint256[],uint256)")

	// hash := sha3.NewKeccak256()
	hash := sha3.NewLegacyKeccak256()
	hash.Write(getExpectedReturnFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedSourceAddress := common.LeftPadBytes(common.HexToAddress(source).Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedSourceAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	addedDestinationAddress := common.LeftPadBytes(common.HexToAddress(destination).Bytes(), 32)
	fmt.Println(hexutil.Encode(addedDestinationAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	paddedPercentage := common.LeftPadBytes(percentage.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedPercentage)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	paddedFlag := common.LeftPadBytes(flag.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedFlag)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedSourceAddress...)
	data = append(data, addedDestinationAddress...)
	data = append(data, paddedAmount...)
	data = append(data, paddedPercentage...)
	data = append(data, paddedFlag...)

	return data

}

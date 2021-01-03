package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

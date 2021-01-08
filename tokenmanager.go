package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/user"
)

const (
	defaultConfigFile = ".1inch/config.json"
)

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

//TokenInfos 用于解析tokens的数据
type TokenInfos struct {
	TConfig []TokenConfig `json:"tokens"`
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

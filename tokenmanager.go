package main

import (
	"fmt"
	"math/big"
)

//TokenConfig 交易对信息
type TokenConfig struct {
	Name          string `jons:"name"`
	SourceAddress string `json:"sourceaddress"`
	Destination   string `json:"destinationaddress"`
	Upper         string `json:"upper"` //价格波动的上限
	Lower         string `json:"lower"` //价格波动的下限
}

//CheckStrategy 用于判断当前是否达到要求
func (Tcon *TokenConfig) CheckStrategy(dis DistributionValue) (bool, error) {
	//1. 价格是否大于一定的数值
	//比如当前设定为 eth-dai 800
	if dis.ReturnAmount.Cmp(big.NewInt(100)) >= 0 {
		fmt.Println("get the right price for swaping")
		return true, nil
	}
	return false, nil
}

//eth-dai
//gas的费用大家都一样

//如何抽象的实现token的配置？
//1. 利用json数组来完成 token对的遍历

//增加json的marshal unmarshal的方法

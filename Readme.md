### 按照配置在1inch上自动交易配置币种

#### 服务组成；

    1. 守护进程 oneinchrobot
    2. cli     robotclient 

#### 配置方法

```shell
    
{
"tokens":[ {
	        "name" : "dai-bas",                                               //交易对名 一般规定较大的在后面 A<->B
            "sourceaddress" : "0x6b175474e89094c44da98b954eedeac495271d0f",   //A币种合约
            "destinationaddress" : "0xa7ED29B253D8B4E3109ce07c80fc570f81B63696", //B币种合约
            "upper" : "100",                                                  //价格监控的上限 B 卖成 A
            "lower" : "60",                                                   //价格监控的下限 A 买成 B 
            "tradeaddress":"0x61c7Cbd3fa89113E9d9cb40dB43587fD6C3032c5",      //操作的地址
            "tradeaddresspriv":"xxxxx",                                       //私钥
            "tradeaddressreserveeth":"0.1",                                   //操作地址eth的保留最小余额
            "tradeaddressreservesrc":"0.1",                                   //操作地址source币种的保留最小余额
            "tradeaddressreservedest":"0.1",                                  //操作地址destination币种的保留最小余额
            "precisionsource":18,                                             //source币种的精度
            "precisiondestination":18,                                        //destination币种的精度
            "status":"enable",                                                //当前交易对的状态
            "slipper":3,                                                      //交易滑点 设置范围 0-100
            "gaslimit":3000000,                                               //gaslimit 可调 目前自动获取遇到问题，暂时配置来做
			"gaspricetimes":10
	}]
}

```

#### 使用方法

编译主程序为 oneinchrobot   
编译cmd目录下的cli为 robotclient


1. 将配置准备好，然后将需要交易的交易对的状态status设置成disable 
2. 启动oneinchrobot ，调用命令robotclient --approve true 来调用approve两个币种合约地址  source 和 destination   spender 设置为onesplitaudit的合约地址
3. 等待上一步执行成功，确保交易已经上链之后，将交易对中的status改成enable  并执行./robotclient --config true
4. 生效后系统将按照 价格上限和下限 低买高卖



#### 下一步计划
1. 增加更多的交易策略
2. 代码重构，完善单元测试和性能测试
### https://github.com/woshidama323/ETHContractPractice.githttps://github.com/woshidama323/ETHContractPractice.git去中心化聚合交易所

#### 背景：

dex 领域 也是资金流转的新的领域

这个部分可以补充中心化交易所的不足，但是未来是一个新的监控系统的监控对象，

选择用golang实现，好处可以很好的继承到现有的监控系统中，



选择1inch交易所，其是一个聚合交易所，学习基本的框架，这样有利于理解dex的运行机制



### 目标：

1. 自动执行系统 
2. 可以提供命令行



+ 可以执行合约接口
+ 可以监控合约行为
+ 可以增加策略



### 现有资源

1. 官方的学习资料

   ​	https://goethereumbook.org/zh/util-go/

2. 1inch的基本资料

   https://docs.1inch.exchange/api/



### 需要调研的数据

1. 如何利用golang 调用合约

   ```shell
   ### 学习体验 golang调用rinkeby合约
   
   ```

   

2. 1inch的合约接口有那一些，如何利用golang进行调用

   

3. 需求明确

   ```shell
   1. 两个服务，一个后台可执行程序，一个cli工具
   2. 可以运行在mac等linux系统上
   3. 可以设置 买卖单的触发价格，gas费用的上线，
   4. 可以设置 开启与关闭策略
   5. 数据存入到sqlite3的数据库中 （前期放到内存中）
   
   6. 写入到json
   ```

   

### 方案设计

```shell
## 监控系统
1. 监控log （前提是如何知道需要监控那一些log ？ ） 走infura的路径
   获取price相关的信息（如果可以行的话） 然后存入到本地，作为价格的监控与趋势的判断
   
   多个币种：
   1. erc20 
   
   如何将币种信息获取到
   
## action系统
1. 如果价格达到设定的范围，则需要触发合约调用，
2. 根据配置信息，更新action的条件

### 难点：
	多个币种

## 配置设置
2. 提供命令行，更新配置文件，将配置文件中的数据更新到内存中，用于每次自动交易的发起放
    这里需要提前判断eth 的余额 
    erc20的余额
    

### 关键信息收集
1. 1inch支持那些交易多 uinswap 0x bancer等等
2. 顺序是什么 在github上

### 核心数据结构设计
1. 交易对
2. 交易对的address
3. 价格是什么 
### 接口的定义是什么
1.监控某块
设计一个单独协程，用于定时进行url的请求 更新设计好的交易对的当前价格

    
```

![image-20210102120537197](/Users/mk/Library/Application Support/typora-user-images/image-20210102120537197.png)



### 关键操作记录

```shell

solc --abi onesplitauditor.sol  -o build
solc --bin onesplitauditor.sol -o build
abigen --bin=OneSplitAudit.bin --abi=OneSplitAudit.abi --pkg=onesplitaudit --out=onesplitaudit.go
将在build文件夹中产生出go 代码 

注意：
solc的版本，
https://docs.soliditylang.org/en/v0.5.3/installing-solidity.html

mac下注意版本
1. brew search solc
然后按照需要的版本下载安装就可以了
```



### 如何理解getreturn信息

```shell
10 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0


```



### truffle中辅助测试

```shell
let dai = new web3.eth.Contract(IERC20.abi, "0x6b175474e89094c44da98b954eedeac495271d0f")

let dai = new web3.eth.Contract(IERC20.abi, "0xa7ED29B253D8B4E3109ce07c80fc570f81B63696")


  dai.methods.approve("0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E","0xffffffffffffffffffffffffffffffffffffffff").send({from: [0]})

dai.methods.balanceOf(accounts[0]).call()

web3.eth.getBalance(accounts[0])

dai.methods.decimals(accounts[0]).call()



```



```shell
ganache-cli -f https://mainnet.infura.io/v3/6707728235da4599b4a045c1b40ff0d9 -i 66 -l 0x31480704F726cD60010b1FEF1dB1c2a4f7bDD67f 800000000000
```



```shell
{
"tokens":[ {
	                "name" : "dai-bas",   //交易对名 一般规定较大的在后面 A<->B
                        "sourceaddress" : "0x6b175474e89094c44da98b954eedeac495271d0f", A
                        "destinationaddress" : "0xa7ED29B253D8B4E3109ce07c80fc570f81B63696",
                        "upper" : "100",
                        "lower" : "60",
                        "tradeaddress":"0x61c7Cbd3fa89113E9d9cb40dB43587fD6C3032c5",
                        "tradeaddresspriv":"xxxxx",
                        "tradeaddressreserveeth":"0.1",
                        "tradeaddressreservesrc":"0.1",
                        "tradeaddressreservedest":"0.1",
                        "precisionsource":18,
                        "precisiondestination":18,
                        "status":"enable",
                        "slipper":3,
                        "gaslimit":3000000,
			"gaspricetimes":10
	}]
}


```

### 先approve 给source 和 destination的合约地址

```shell
用命令 ./cmd --config true 更新配置 
用命令 ./cmd --approve true approve 两个币种到 onesplitaudit合约

```


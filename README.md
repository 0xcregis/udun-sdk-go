### udun-sdk-go

为集成优盾钱包提供golang版本SDK

### Install
>golang 版本需1.16以上

go get -u github.com/uduncloud/udun-sdk-go

### Usage

1.init client

   client:=NewUdunClient(config.Config{Url: url})
   > Config.url: node server address 
   >
   > Config.Timeout: http.connect time out
   
2.call function
  
   - [CheckAddress](https://www.uduncloud.com/geteway-interface?index=4)  
   - [ExistAddress](https://www.uduncloud.com/geteway-interface?index=6)
   - [GetSupportCoinsByMerchant](https://www.uduncloud.com/geteway-interface?index=5)
   - [CreateAddress](https://www.uduncloud.com/geteway-interface?index=1)
   - [Withdraw](https://www.uduncloud.com/geteway-interface?index=2)

### example
```
	c := NewUdunClient(config.Config{Url: url})
	base := common.Request{Timestamp: time.Now().Unix(), Nonce: 1000, Key: key}
	log.Println(c.GetSupportCoinsByMerchant(base, "309914", true))
```

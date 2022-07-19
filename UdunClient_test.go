package udun_sdk_go

import (
	"github.com/uduncloud/udun-sdk-go/common"
	"github.com/uduncloud/udun-sdk-go/config"
	"github.com/uduncloud/udun-sdk-go/utils"
	"log"
	"testing"
	"time"
)

var url = "https://sig10.udun.io"
var key = "5ae108d94ffbade8564ef6a0da2dad34"

func TestUtil(t *testing.T) {
	log.Println(utils.Signature(1001, time.Now().Unix(), "hello", key))
}

func TestCheckAddress(t *testing.T) {
	c := NewUdunClient(config.Config{Url: url})
	base := common.Request{Timestamp: time.Now().Unix(), Nonce: 1000, Key: key}
	v := common.VerifyRequest{MerchantId: "309914", MainCoinType: "60", Address: "0xf2c48ea7819b69221984b94ec8e188b72b015b88"}
	log.Println(c.CheckAddress(base, []*common.VerifyRequest{&v}))
}

func TestExistAddress(t *testing.T) {
	c := NewUdunClient(config.Config{Url: url})
	base := common.Request{Timestamp: time.Now().Unix(), Nonce: 1000, Key: key}
	v := common.VerifyRequest{MerchantId: "309914", MainCoinType: "60", Address: "0xf2c48ea7819b69221984b94ec8e188b72b015b88"}
	log.Println(c.ExistAddress(base, []*common.VerifyRequest{&v}))
}

func TestGetSupportCoinsByMerchant(t *testing.T) {
	c := NewUdunClient(config.Config{Url: url})
	base := common.Request{Timestamp: time.Now().Unix(), Nonce: 1000, Key: key}
	log.Println(c.GetSupportCoinsByMerchant(base, "309914", true))
}

func TestCreateAddress(t *testing.T) {
	c := NewUdunClient(config.Config{Url: url})
	base := common.Request{Timestamp: time.Now().Unix(), Nonce: 1000, Key: key}
	call := common.CallbackRequest{MerchantId: "309914", MainCoinType: 60, CallUrl: "http://localhost:8080/callBack"}
	log.Println(c.CreateAddress(base, []*common.CallbackRequest{&call}))
}

func TestWithdraw(t *testing.T) {
	c := NewUdunClient(config.Config{Url: url})
	base := common.Request{Timestamp: time.Now().Unix(), Nonce: 1000, Key: key}
	call := common.WithdrawRequest{MerchantId: "309914", MainCoinType: "60", CoinType: "60", Amount: "0.11", CallUrl: "http://localhost:8080/callBack", Address: "0x465316d9e9b603c306ab6443857bbb653c687081", BusinessId: "9001"}
	log.Println(c.Withdraw(base, []*common.WithdrawRequest{&call}))
}

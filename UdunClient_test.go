package udun_sdk_go

import (
	"github.com/uduncloud/udun-sdk-go/common"
	"github.com/uduncloud/udun-sdk-go/config"
	"github.com/uduncloud/udun-sdk-go/utils"
	"log"
	"testing"
	"time"
)

var url = "http://gateway01.ruibai.com"
var key = "ebd7a5ad84a71da31143116e4ed1ea24"

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
	call := common.CallbackRequest{MerchantId: "400044", MainCoinType: 520, CallUrl: "http://localhost:8080/callBack"}
	log.Println(c.CreateAddress(base, []*common.CallbackRequest{&call}))
}

func TestWithdraw(t *testing.T) {
	c := NewUdunClient(config.Config{Url: url})
	base := common.Request{Timestamp: time.Now().Unix(), Nonce: 1000, Key: key}
	call := common.WithdrawRequest{MerchantId: "400044", MainCoinType: "195", CoinType: "195", Amount: "0.11", CallUrl: "http://localhost:8080/callBack", Address: "TDFwiy9qNjT9ryFS2Wsyacfa5PHS3hUwxn", BusinessId: "9003"}
	log.Println(c.Withdraw(base, []*common.WithdrawRequest{&call}))
}

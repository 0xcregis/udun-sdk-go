package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"udun-sdk-go/common"
	"udun-sdk-go/config"
	"udun-sdk-go/utils"
)

type Tx struct {
	client *http.Client
	config *config.Config
}

func NewTxClient(c *http.Client, config *config.Config) common.TXImpl {
	return &Tx{
		client: c,
		config: config,
	}
}

func (t *Tx) CreateAddress(baseReq common.Request, callback []*common.CallbackRequest) (code int64, msg string, data string) {
	reqBody, _ := json.Marshal(callback)
	sign := utils.Signature(baseReq.Nonce, baseReq.Timestamp, string(reqBody), baseReq.Key)
	url := fmt.Sprintf("%v%v", t.config.Url, common.CreateAddress)
	return utils.SendRequest(t.client, url, baseReq.Nonce, baseReq.Timestamp, string(reqBody), sign)
}

func (t *Tx) Withdraw(baseReq common.Request, withdraw []*common.WithdrawRequest) (code int64, msg string) {
	reqBody, _ := json.Marshal(withdraw)
	sign := utils.Signature(baseReq.Nonce, baseReq.Timestamp, string(reqBody), baseReq.Key)
	url := fmt.Sprintf("%v%v", t.config.Url, common.Withdraw)
	code, msg, _ = utils.SendRequest(t.client, url, baseReq.Nonce, baseReq.Timestamp, string(reqBody), sign)
	return
}

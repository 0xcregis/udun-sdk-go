package internal

import (
	"encoding/json"
	"fmt"
	"github.com/uduncloud/udun-sdk-go/common"
	"github.com/uduncloud/udun-sdk-go/config"
	"github.com/uduncloud/udun-sdk-go/utils"
	"net/http"
)

type Verify struct {
	client *http.Client
	config *config.Config
}

func NewVerifyClient(c *http.Client, config *config.Config) common.VerifyImpl {
	return &Verify{
		client: c,
		config: config,
	}
}

func (v *Verify) CheckAddress(baseReq common.Request, verify []*common.VerifyRequest) (code int64, msg string) {
	reqBody, _ := json.Marshal(verify)
	sign := utils.Signature(baseReq.Nonce, baseReq.Timestamp, string(reqBody), baseReq.Key)
	url := fmt.Sprintf("%v%v", v.config.Url, common.CheckAddress)
	code, msg, _ = utils.SendRequest(v.client, url, baseReq.Nonce, baseReq.Timestamp, string(reqBody), sign)
	return
}

func (v *Verify) ExistAddress(baseReq common.Request, verify []*common.VerifyRequest) (code int64, msg string) {
	reqBody, _ := json.Marshal(verify)
	sign := utils.Signature(baseReq.Nonce, baseReq.Timestamp, string(reqBody), baseReq.Key)
	url := fmt.Sprintf("%v%v", v.config.Url, common.ExistAddress)
	code, msg, _ = utils.SendRequest(v.client, url, baseReq.Nonce, baseReq.Timestamp, string(reqBody), sign)
	return
}

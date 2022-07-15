package internal

import (
	"encoding/json"
	"fmt"
	"gitcool.co/htsun/udun-sdk-go/common"
	"gitcool.co/htsun/udun-sdk-go/config"
	"gitcool.co/htsun/udun-sdk-go/utils"
	"net/http"
)

type Merchant struct {
	client *http.Client
	config *config.Config
}

func NewMerchantClient(c *http.Client, config *config.Config) common.MerchantImpl {
	return &Merchant{
		client: c,
		config: config,
	}
}

func (m *Merchant) GetSupportCoinsByMerchant(baseReq common.Request, merchantId string, showBalance bool) (code int64, msg string, extra string) {
	mp := make(map[string]interface{}, 0)
	mp["merchantId"] = merchantId
	mp["showBalance"] = showBalance
	reqBody, _ := json.Marshal(mp)
	sign := utils.Signature(baseReq.Nonce, baseReq.Timestamp, string(reqBody), baseReq.Key)
	url := fmt.Sprintf("%v%v", m.config.Url, common.SupportCoins)
	code, msg, extra = utils.SendRequest(m.client, url, baseReq.Nonce, baseReq.Timestamp, string(reqBody), sign)
	return
}

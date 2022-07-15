package udun_sdk_go

import (
	"net/http"
	"udun-sdk-go/common"
	"udun-sdk-go/config"
	"udun-sdk-go/internal"
)

type UdunClient struct {
	common.MerchantImpl
	common.TXImpl
	common.VerifyImpl
	Client *http.Client
	Config config.Config
}

func NewUdunClient(config config.Config) *UdunClient {
	c := &http.Client{}
	return &UdunClient{internal.NewMerchantClient(c, &config), internal.NewTxClient(c, &config), internal.NewVerifyClient(c, &config), c, config}
}

func (u *UdunClient) Http(f func(h *http.Client, config config.Config)) {
	f(u.Client, u.Config)
}

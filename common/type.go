package common

const (
	CreateAddress = "/mch/address/create"
	Withdraw      = "/mch/withdraw"
	CheckAddress  = "/mch/check/address"
	SupportCoins  = "/mch/support-coins"
	ExistAddress  = "/mch/exist/address"
)

type MerchantImpl interface {
	GetSupportCoinsByMerchant(baseReq Request, merchantId string, showBalance bool) (code int64, msg string, extra string)
}

type TXImpl interface {
	CreateAddress(baseReq Request, callback []*CallbackRequest) (code int64, msg string, extra string)

	Withdraw(baseReq Request, withdraw []*WithdrawRequest) (code int64, msg string)
}

type VerifyImpl interface {
	CheckAddress(baseReq Request, verify []*VerifyRequest) (code int64, msg string)
	ExistAddress(baseReq Request, verify []*VerifyRequest) (code int64, msg string)
}

type VerifyRequest struct {
	MerchantId   string `json:"merchantId"`
	MainCoinType string `json:"mainCoinType"`
	Address      string `json:"address"`
}

type CallbackRequest struct {
	MerchantId   string `json:"merchantId"`
	MainCoinType int    `json:"mainCoinType"`
	CallUrl      string `json:"callUrl"`
	WalletId     string `json:"walletId"`
	Alias        string `json:"alias"`
}

type WithdrawRequest struct {
	/**
	{
	        "address":"raadSxrUhG5EQVCY75CSGaVLWCeXd6yH6s",
	        "amount":"0.11",
	        "merchantId":"100109",
	        "mainCoinType":"144",
	        "coinType":"144",
	        "callUrl":"http://localhost:8080/mch/callBack",
	        "businessId":"15",
	        "memo":"10112"
	    }
	*/

	Address      string `json:"address"`
	Amount       string `json:"amount"`
	MerchantId   string `json:"merchantId"`
	MainCoinType string `json:"mainCoinType"`
	CoinType     string `json:"coinType"`
	CallUrl      string `json:"callUrl"`
	BusinessId   string `json:"businessId"`
	WalletId     string `json:"walletId"`
	Memo         string `json:"memo"`
}
type Request struct {
	Timestamp int64  `json:"timestamp"`
	Nonce     int    `json:"nonce"`
	Key       string `json:"key"`
}

type HttpRequest struct {
	Timestamp int64  `json:"timestamp"`
	Nonce     int    `json:"nonce"`
	Sign      string `json:"sign"`
	Body      string `json:"body"`
}

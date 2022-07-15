package utils

import (
	"bytes"
	"encoding/json"
	"gitcool.co/htsun/udun-sdk-go/common"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func SendRequest(client *http.Client, url string, nonce int, timestamp int64, reqBody string, sign string) (code int64, msg string, extra string) {
	q := common.HttpRequest{Nonce: nonce, Timestamp: timestamp, Sign: sign, Body: reqBody}
	body, _ := json.Marshal(q)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return http.StatusBadRequest, err.Error(), ""
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadGateway, err.Error(), ""
	}

	js := gjson.ParseBytes(r)
	code = js.Get("code").Int()
	msg = js.Get("message").String()
	if js.Get("data").Exists() {
		extra = js.Get("data").String()
	}
	return
}

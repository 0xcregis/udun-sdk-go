package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Signature(nonce int, timestamp int64, body string, key string) string {
	msg := fmt.Sprintf("%v%v%v%v", body, key, nonce, timestamp)
	h := md5.New()
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}

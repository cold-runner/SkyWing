package service

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "skylab.org.cn"

func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}
func decryptPassword(data string) (result string) {
	decodeString, _ := hex.DecodeString(data)
	return string(decodeString)
}

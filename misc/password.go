package misc

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Password(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	h.Write([]byte("@iguiyu.com"))
	return hex.EncodeToString(h.Sum(nil))
}

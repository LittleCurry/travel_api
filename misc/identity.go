package misc

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateId(seed string) string {
	h := md5.New()
	h.Write([]byte(seed))
	h.Write([]byte("Id"))
	return hex.EncodeToString(h.Sum(nil))
}

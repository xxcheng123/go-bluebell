package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(content string) (encryptContent string) {
	secret := "xxcheng.cn"
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(content)))
}

package core

import (
	"crypto/md5"
	"encoding/hex"
)

func HashString(requestResponse string) string {

	hash := md5.Sum([]byte(requestResponse))
	return hex.EncodeToString(hash[:])

}

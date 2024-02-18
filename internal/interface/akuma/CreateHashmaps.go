package akuma

import (
	"crypto/md5"
	"encoding/hex"
)

func CreateHashmaps(name string) string {
	hasher := md5.New()
	hasher.Write([]byte(name))
	return hex.EncodeToString(hasher.Sum(nil))
}

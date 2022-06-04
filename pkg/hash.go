package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/fnv"
)

func HashPassword(password string) string {
	return GetMD5Hash(password)
}

func HashString(str string) string {
	h := fnv.New32a()
	h.Write([]byte(str))
	return fmt.Sprint(h.Sum32())
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
 }
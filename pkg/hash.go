package pkg

import (
	"fmt"
	"hash/fnv"
)

func HashPassword(password string) string {
	return password
}

func HashString(str string) string {
	h := fnv.New32a()
	h.Write([]byte(str))
	return fmt.Sprint(h.Sum32())
}
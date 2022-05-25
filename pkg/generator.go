package pkg

import (
	"math/rand"
	"time"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	source      = rand.NewSource(time.Now().UnixNano())
	random      = rand.New(source)
)

func RandSID(n int) string {
	sid := make([]rune, n)
	for i := range sid {
		sid[i] = letterRunes[random.Intn(len(letterRunes))]
	}
	return string(sid)
}

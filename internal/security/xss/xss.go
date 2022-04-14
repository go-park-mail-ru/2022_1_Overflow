package xss

import (
	"github.com/microcosm-cc/bluemonday"
)

var P *bluemonday.Policy

func init() {
	P = bluemonday.StrictPolicy()
}
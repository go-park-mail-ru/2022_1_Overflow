package xss

import (
	"html"

	"github.com/microcosm-cc/bluemonday"
)

var P *bluemonday.Policy

func init() {
	P = bluemonday.StrictPolicy()
}

func EscapeInput(input string) string {
	return html.EscapeString(input)
}
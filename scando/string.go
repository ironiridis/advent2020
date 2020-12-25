package scando

import (
	"io/ioutil"
	"strings"
)

func String(s string) chan string {
	return lines(ioutil.NopCloser(strings.NewReader(s)))
}

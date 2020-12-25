package scando

import (
	"bufio"
	"io"
	"strings"
)

func lines(r io.ReadCloser) (c chan string) {
	c = make(chan string)
	go func() {
		defer r.Close()
		defer close(c)
		s := bufio.NewScanner(r)
		for s.Scan() {
			c <- strings.TrimSpace(s.Text())
		}
		if err := s.Err(); err != nil {
			panic(err)
		}
	}()
	return c
}

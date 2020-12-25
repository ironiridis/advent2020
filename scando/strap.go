package scando

import "testing"

type Strappable func(chan string) (string, error)

func Strap(fn Strappable, in, exp string, t *testing.T) {
	t.Helper()
	ans, err := fn(String(in))
	if err != nil {
		t.Errorf("Cannot determine answer: %v\n", err)
		return
	}
	if ans != exp {
		t.Errorf("Expected %q, got %q\n", exp, ans)
	}
}

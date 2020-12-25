package scando

import "os"

func File(fn string) chan string {
	fp, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	return lines(fp)
}

func Stdin() chan string {
	return lines(os.Stdin)
}

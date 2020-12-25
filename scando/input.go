package scando

import (
	"fmt"
	"os"
)

func Input() chan string {
	_, err := os.Stat("input.txt")
	if err != nil {
		fmt.Println("No input.txt file. Paste input data:")
		return Stdin()
	}
	fmt.Println("Reading from input.txt")
	return File("input.txt")
}

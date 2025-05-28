package main

import (
	"os"
	"fmt"
)

func main() {
	fileName := "memo.txt"
	fp, err := os.Open(fileName)
	if err != nil {
		fmt.Print("1")
	}
	defer fp.Close()

	var lines []string

	for {
		buf := make([]byte, 64)
		n, err := fp.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Print("2")
		}
		lines = append(lines, string(buf))
	}
	fmt.Print(lines)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	const maxCapacity = 1024*1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var a string
	scanner.Scan()
	fmt.Sscanln(scanner.Text(), &a)

	var b []rune

	for _, rune := range a {
		if rune == '<' {
			if len(b) > 0 {
				b = b[:len(b)-1]
			}
		} else {
			b = append(b, rune)
		}
	}

	if len(b) > 0 {
		fmt.Println(string(b))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var a int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &a)

	for i := 0; i < a; i++ {
		scanner.Scan()
		s:=scanner.Text()

		if len(s) > 11 && s[:10] == "simon says" {
			fmt.Println(s[11:])
		} else {
			fmt.Println()
		}
	}
}

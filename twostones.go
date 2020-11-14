package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var a int64
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &a)

	res := "Alice"
	if a%2 == 0 {
		res = "Bob"
	}

	fmt.Print(res)
}

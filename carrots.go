package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	var a,b int
	fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)

	fmt.Println(b)
}

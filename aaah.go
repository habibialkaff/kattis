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
	a := scanner.Text()
	scanner.Scan()
	b := scanner.Text()

	if len(a) < len(b) {
		fmt.Print("no")
	} else {
		fmt.Println("go")
	}
}

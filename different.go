package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var a,b int64
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d%d\n", &a, &b)
		c := a-b
		if c<0 {
			c = c*-1
		}
		fmt.Println(c)
	}
}

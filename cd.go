package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for {
		var a, b int64
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d%d\n", &a, &b)
		if a == 0 && b == 0 {
			break
		}

		dict := make(map[string]bool, a)

		var i int64
		for i = 0; i < a; i++ {
			scanner.Scan()
			dict[scanner.Text()] = true
		}

		result := 0

		for i = 0; i < b; i++ {
			scanner.Scan()
			if dict[scanner.Text()] {
				result++
			}
		}

		fmt.Printf("%d\n", result)
	}
}

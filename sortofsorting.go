package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	printBlank := false

	for scanner.Scan() {
		var a int
		fmt.Sscanf(scanner.Text(), "%d", &a)

		if a == 0 {
			break
		}

		if printBlank {
			fmt.Println()
		}
		printBlank = true

		var b []string
		for i := 0; i < a; i++ {
			scanner.Scan()
			b = append(b, scanner.Text())
		}

		sort.SliceStable(b, func(i, j int) bool {
			if b[i][0] < b[j][0] {
				return true
			}
			return b[i][0] == b[j][0] && b[i][1] < b[j][1]
		})

		for i := 0; i < a; i++ {
			fmt.Println(b[i])
		}
	}
}

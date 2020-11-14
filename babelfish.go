package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	isDict := true
	dict := make(map[string]string)

	for scanner.Scan() {
		a := scanner.Text()

		if a == "" {
			isDict = false
			continue
		}

		if isDict {
			var b, c string
			fmt.Sscanf(a, "%s %s", &b, &c)
			dict[c] = b
		} else {
			var b string
			fmt.Sscanf(a, "%s", &b)

			if val, ok := dict[b]; ok {
				fmt.Println(val)
			} else {
				fmt.Println("eh")
			}
		}
	}
}

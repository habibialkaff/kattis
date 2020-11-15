package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	scanner.Scan()

	for scanner.Scan() {
		var ct int64
		var val []int64

		for scanner.Scan() {
			t := scanner.Text()

			if t == "" {
				break
			}

			var b int64
			fmt.Sscanf(t, "%d", &b)
			val = append(val, b)
		}

		ct = int64(len(val))
		var res int64

		for i := 0; i < len(val); i++ {
			res += val[i] % ct
		}

		if res%ct == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

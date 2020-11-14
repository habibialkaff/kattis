package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	for scanner.Scan() {

		a:=strings.Split(scanner.Text(), " ")

		a = a[1:]

		var b []int

		for _, i := range a {
			j, _ := strconv.Atoi(i)
			b = append(b, j)
		}

		sum := 0

		for i := 0; i < len(b); i++ {
			sum += b[i]
		}

		avg := (float64(sum)) / (float64(len(b)))

		ct:= 0

		for i := 0; i < len(b); i++ {
			if float64(b[i]) > avg {
				ct++
			}
		}

		fmt.Printf("%.3f%%\n", (float64(ct) * 100) / float64(len(b)))
	}
}

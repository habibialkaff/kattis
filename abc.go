package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getVal(c byte, arr []int) int {
	if c == 'A' {
		return arr[0]
	}

	if c == 'B' {
		return arr[1]
	}

	return arr[2]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var a,b,c int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d %d", &a, &b, &c)

	arr := []int{a,b,c}
	sort.Ints(arr)

	var c1,c2,c3 byte
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%c%c%c", &c1, &c2, &c3)

	r1 := getVal(c1, arr)
	r2 := getVal(c2, arr)
	r3 := getVal(c3, arr)

	fmt.Printf("%d %d %d", r1, r2, r3)
}

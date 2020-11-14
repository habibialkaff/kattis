package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var arr [16]int

	for i := 0; i < 4; i++ {
		scanner.Scan()
		var a, b, c, d int
		fmt.Sscanf(scanner.Text(), "%d %d %d %d", &a, &b, &c, &d)
		arr[i*4] = a
		arr[i*4+1] = b
		arr[i*4+2] = c
		arr[i*4+3] = d
	}

	scanner.Scan()
	var mv int
	fmt.Sscanf(scanner.Text(), "%d", &mv)

	var res [16]int

	for i := 0; i < 4; i++ {
		var a, b, c, d int
		if mv == 0 {
			a = i * 4
			b = i*4 + 1
			c = i*4 + 2
			d = i*4 + 3
		} else if mv == 1 {
			a = i
			b = i + 4
			c = i + 8
			d = i + 12
		} else if mv == 2 {
			a = i*4 + 3
			b = i*4 + 2
			c = i*4 + 1
			d = i * 4
		} else {
			a = i + 12
			b = i + 8
			c = i + 4
			d = i
		}

		var temp [4]int
		x := 0

		if arr[a] != 0 {
			temp[x] = arr[a]
			x++
		}
		if arr[b] != 0 {
			temp[x] = arr[b]
			x++
		}
		if arr[c] != 0 {
			temp[x] = arr[c]
			x++
		}
		if arr[d] != 0 {
			temp[x] = arr[d]
			x++
		}

		if temp[0] == temp[1] {
			res[a] = 2 * temp[0]
			res[b] = temp[2]
			res[c] = temp[3]
			res[d] = 0
			if temp[2] == temp[3] {
				res[b] = 2 * temp[2]
				res[c] = 0
			}
		} else {
			res[a] = temp[0]
			if temp[1] == temp[2] {
				res[b] = 2 * temp[1]
				res[c] = temp[3]
				res[d] = 0
			} else {
				res[b] = temp[1]
				if temp[2] == temp[3] {
					res[c] = 2 * temp[2]
					res[d] = 0
				} else {
					res[c] = temp[2]
					res[d] = temp[3]
				}
			}
		}
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("%d %d %d %d\n", res[i*4], res[i*4+1], res[i*4+2], res[i*4+3])
	}
}

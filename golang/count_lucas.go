package main

import (
	"os"
	"strconv"
	"fmt"
)

func lucas() func() (int, int) {
	prev, next, count := 2, 1, 1
	return func() (int, int) {
		prev, next, count = next, prev+next, count+1
		return next, count
	}
}

func main() {
	args := os.Args
	num, _ := strconv.Atoi(args[1])

	if num <= 2 {
		fmt.Println(num)
	} else {
		f := lucas()
		for {
			n, c := f()
			if n > num || n < 0 {
				fmt.Println(c)
				break
			}
		}
	}
}
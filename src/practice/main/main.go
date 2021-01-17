package main

import "fmt"

func zipTo2DArray(xs string, ys string) {
	n := len(xs)
	if len(ys) < len(xs) {
		n = len(ys)
	}
	zipped := make([][] int32, n)
	for i, x := range xs {
		// len(ys) < len(xs)
		if len(zipped) == i {
			break
		}
		zipped[i] = make([]int32, 2)
		zipped[i][0] = x
	}

	for i, y := range ys {
		// len(ys) < len(xs)
		if len(zipped) == i {
			break
		}
		zipped[i][1] = y
	}

	fmt.Println(zipped)
}

func main() {
	zipTo2DArray("", "a")
	zipTo2DArray("a", "")
	zipTo2DArray("abab", "abcd")
	zipTo2DArray("abab", "abc")
	zipTo2DArray("abc", "abab")
}

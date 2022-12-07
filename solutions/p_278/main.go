package main

import "fmt"

func main() {
	n := 5
	fmt.Println(firstBadVersion(n))
}

var bad = 4

func isBadVersion(v int) bool {
	return v == bad
}

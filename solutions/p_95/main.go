package main

import "fmt"

func main() {
	n := 3
	for _, tree := range generateTrees(n) {
		fmt.Println(tree)
	}
}

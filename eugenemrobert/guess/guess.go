package main

import (
	"fmt"
)

func functionD(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Print(a)
	}
	fmt.Println()
}
func main() {
	functionD("ha", 4)
}

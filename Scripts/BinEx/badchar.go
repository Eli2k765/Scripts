package main

import "fmt"

func main() {
	for x := 1; x < 256; x++ {
		fmt.Printf("\\x" + "%02x", x)
	}
	fmt.Println()
}

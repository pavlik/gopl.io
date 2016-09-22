package main

import (
	"fmt"
	"os"
)

func main() {
	for n, arg := range os.Args[1:] {
		fmt.Printf("Index: %d\t value: %s\n", n+1, arg)
	}

}

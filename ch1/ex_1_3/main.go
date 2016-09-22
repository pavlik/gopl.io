package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Slow algoritm")
	start := time.Now()
	slowAlgorithm()
	fmt.Printf("%dns elapsed\n", time.Since(start).Nanoseconds())

	fmt.Printf("\n\n%s\n", "Fast algoritm")
	start = time.Now()
	fastAlgorithm()
	fmt.Printf("%dns elapsed\n", time.Since(start).Nanoseconds())

}

func slowAlgorithm() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func fastAlgorithm() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

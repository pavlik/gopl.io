package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	mapFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, mapFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, mapFiles)
			f.Close()
		}
		for line, n := range counts {
			if n >= 1 {
				fmt.Printf("%d\t %s\t %s\n", n, line, mapFiles[line])
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, mapFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		mapFiles[input.Text()] = append(mapFiles[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

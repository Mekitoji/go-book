package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type fLine struct {
	count     int
	filenames []string
}

func main() {
	counts := make(map[string]*fLine)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
		}
	}

	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\n%s\n\n", n.count, line, strings.Join(n.filenames, " "))
		}
	}
}

func countLines(f *os.File, c map[string]*fLine) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		if _, ok := c[key]; ok {
			c[key].count++
			if ok := contains(f.Name(), c[key].filenames); !ok {
				c[key].filenames = append(c[key].filenames, f.Name())
			}
		} else {
			// c[key] = fLine{1, append(c[key].filenames, f.Name())}
			c[key] = new(fLine)
			c[key].count++
			c[key].filenames = append(c[key].filenames, f.Name())
		}
	}
}

func contains(value string, s []string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

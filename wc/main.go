package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(reader io.Reader, countLines bool) int {

	scanner := bufio.NewScanner(reader)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "count lines")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

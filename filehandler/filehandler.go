package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// pushes a function call onto a list.
	//defer file.close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("read.txt")
	if err != nil {
		//Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		fmt.Println(i, line)
	}

	if err := writeLines(lines, "write.txt"); err != nil {
		log.Fatalf("WriteLines: %s", err)
	}
}

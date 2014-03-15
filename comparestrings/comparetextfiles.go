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
	// ensure file closes even if call to Open fails.
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

func difference(slice1, slice2 []string) []string {

	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. Add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func main() {
	slice1, err := readLines("read1.txt")
	slice2, err2 := readLines("read2.txt")
	if err != nil && err2 != nil {
		//Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
		log.Fatalf("readLines: %s", err)
		log.Fatalf("readLines: %s", err2)
	}

	fmt.Println("%+v\n", difference(slice1, slice2))

	if err := writeLines(difference(slice1, slice2), "write.txt"); err != nil {
		log.Fatalf("WriteLines: %s", err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
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

func difference(slice1, slice2 []string) {

	var dmp = diffmatchpatch.New()
	str1 := string(slice2[:])
	str2 := string(slice1[:])

	var diff = dmp.DiffMain(str1, str2, false)

	// if there is some changes made return them. If not return stored document.
	if diff != 1 {
		return diff
	}

	return slice1

}

func main() {
	slice1, err := readLines("read1.txt")  // stored document.
	slice2, err2 := readLines("read2.txt") // document with changes.
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

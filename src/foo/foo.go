package main

import "wmic"

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"os"
)

func main() {
	// Open the file.
	file, err := os.Open("volume.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read it as UTF-16 LE.
	enc := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM)
	scanner := bufio.NewScanner(transform.NewReader(file, enc.NewDecoder()))
	for scanner.Scan() {
		fmt.Printf("Read line: %s\n", scanner.Bytes())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
}

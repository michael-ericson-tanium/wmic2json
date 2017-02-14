package main

import (
	"encoding/json"
	"fmt"
	"os"
	"wmic"
)

func main() {
	var output []wmic.Output
	var err error
	if len(os.Args) <= 1 {
		output, err = wmic.Translate(os.Stdin)
	} else {
		output, err = wmic.Exec(os.Args[1:]...)
	}
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonBytes)
}

package main

import (
	"encoding/json"
	"fmt"
	"recursiveZip/internals/reader"
)

func main() {
	jsonmarshall, err := json.MarshalIndent(reader.Reader("learn_tests"), "", "  ")
	if err != nil {
		print(err.Error())
	}
	fmt.Println(string(jsonmarshall))
}

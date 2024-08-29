package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"recursiveZip/internals/reader"
)

func Write(dirName, fileName string) error {
	content, err := json.MarshalIndent(reader.Reader(dirName, ""), "", " ")
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s.txt", fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

package decompresser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"recursiveZip/internals/structs"
)

func Decompress(fileName string) error {
	data, err := convertFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	createDir(data.DirectoryName)
	recursiveUnZip(data)

	return nil
}

func recursiveUnZip(data structs.DirectoryContent) {
	for _, v := range data.Files {
		createAndWrite(v)
	}
	for _, v := range data.InterDirectories {
		createDir(v.DirectoryName)
		recursiveUnZip(v)
	}
}

func createDir(dirName string) {
	os.Mkdir(dirName, 0777)
}

func createAndWrite(fileData structs.FileContent) {
	err := os.WriteFile(fileData.FileName, []byte(fileData.Content), 0777)
	if err != nil {
		fmt.Printf("unable to write file: %s", err.Error())
	}
}

func convertFile(fileName string) (structs.DirectoryContent, error) {
	var data structs.DirectoryContent

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

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
	createDir(data.DirectoryName, "")
	recursiveUnZip(data, data.DirectoryName)

	return nil
}

func recursiveUnZip(data structs.DirectoryContent, dirName string) {
	dirName = dirName+"/"
	for _, v := range data.Files {
		createAndWrite(v, dirName)
	}
	for _, v := range data.InterDirectories {
		createDir(v.DirectoryName, dirName)
		recursiveUnZip(v, dirName+v.DirectoryName)
	}
}

func createDir(dirName, parentDirName string) {
	os.Mkdir(parentDirName + dirName, 0777)
}

func createAndWrite(fileData structs.FileContent, dirName string) {
	err := os.WriteFile(dirName+fileData.FileName, []byte(fileData.Content), 0777)
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

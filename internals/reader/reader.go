package reader

import (
	"fmt"
	"os"
	"path/filepath"
	"recursiveZip/internals/structs"
)

func Reader(dirname, parentDir string) structs.DirectoryContent {
	var dr structs.DirectoryContent
	dr.DirectoryName = dirname

	fileslist, dirlist := GetListOfFiles(parentDir+dirname)
	dr.Files = AppendFileContent(fileslist, parentDir+dirname)

	for _, v := range dirlist {
		dr.InterDirectories = append(dr.InterDirectories, Reader(v, parentDir+dirname+"/"))
	}

	return dr
}

func AppendFileContent(fileslist []string, dirname string) []structs.FileContent {
	tempfile := make([]structs.FileContent, 0, len(fileslist))
	for _, v := range fileslist {
		temp := structs.FileContent{
			FileName: v,
			Content:  ReadFile(dirname + "/" + v),
		}
		tempfile = append(tempfile, temp)
	}
	return tempfile
}

func ReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}

	fileContent := string(content)

	return fileContent
}

func GetListOfFiles(dirname string) ([]string, []string) {
	var filelist []string
	var dirlist []string

	entries, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return nil, nil
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			filelist = append(filelist, entry.Name())
		} else {
			dirlist = append(dirlist, entry.Name())
		}
	}

	return filelist, dirlist
}

func GetCurrentDir() string {
	dir := GetCurrentDirFullPath()

	lastDir := filepath.Base(dir)

	return lastDir
}

func GetCurrentDirFullPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return ""
	}
	return dir
}

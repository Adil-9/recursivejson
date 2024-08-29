package reader

import (
	"fmt"
	"os"
	"path/filepath"
	"recursiveZip/internals/structs"
)

func Reader(dirname string) structs.DirectoryContent {
	var dr structs.DirectoryContent
	dr.DirectoryName = dirname

	fileslist, dirlist := GetListOfFiles(dirname)
	dr.Files = AppendFileContent(fileslist)

	for _, v := range dirlist {
		dr.InterDirectories = append(dr.InterDirectories, Reader(v))
	}

	return dr
}

func AppendFileContent(fileslist []string) []structs.FileContent {
	tempfile := make([]structs.FileContent, 0, len(fileslist))
	for _, v := range fileslist {
		temp := structs.FileContent{
			FileName: v,
			Content:  ReadFile(v),
		}
		tempfile = append(tempfile, temp)
	}
	return tempfile
}

func ReadFile(fileName string) string {
	// Read the entire file content
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}

	// Convert to string if needed
	fileContent := string(content)

	// Print or process the content
	return fileContent
}

func GetListOfFiles(dirname string) ([]string, []string) {
	var filelist []string
	var dirlist []string

	// Read the directory contents
	entries, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return nil, nil
	}

	// Iterate over the directory entries
	for _, entry := range entries {
		fullPath := filepath.Join(dirname, entry.Name())
		if !entry.IsDir() { // Check if it's not a directory
			filelist = append(filelist, fullPath) // Print file name
		} else {
			dirlist = append(dirlist, fullPath)
		}
	}

	return filelist, dirlist
}

func GetCurrentDir() string {
	dir := GetCurrentDirFullPath()

	// Get the last directory name
	lastDir := filepath.Base(dir)

	return lastDir
}

func GetCurrentDirFullPath() string {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return ""
	}
	return dir
}

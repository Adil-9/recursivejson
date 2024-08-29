package structs

type DirectoryContent struct {
	DirectoryName    string             `json:"directoryName"`
	Files            []FileContent      `json:"files"`
	InterDirectories []DirectoryContent `json:"interdirectories"`
}

type FileContent struct {
	FileName string `json:"fileName"`
	Content  string `json:"content"`
}

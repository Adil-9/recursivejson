# RecursiveJSON

This Go program recursively compresses all files and directories from a specified folder into a single file. It also allows for decompression on any other device using the generated file.

## Features

- Compress a directory and its contents into a single file.
- Decompress the generated file to recreate the directory structure and contents.

## Usage

### Requirements

- [Go](https://golang.org/doc/install) must be installed on your machine.
- Clone this repository to your local machine.

### Compressing a Folder

To compress all data from a specified folder into a file, use the following command:

```bash
go run cmd/main.go -dir {{directory}}
```
Replace {{directory}} with the path to the folder you want to compress.

Example:

```bash
go run cmd/main.go -dir /path/to/folder
```
This will create a compressed file that contains all the files and directories within /path/to/folder.
Decompressing a File

To decompress the previously created file on any other device, use this command:

```bash
go run cmd/main.go -d -f={{filename}}
```
Replace {{filename}} with the name of the file you want to decompress.

Example:

```bash
go run cmd/main.go -d -f=compressed_file
```
This will extract the contents back to the original folder structure.
Options

    -dir : Specify the directory to compress.
    -d : Enable decompression mode.
    -f : Specify the filename for decompression.

Note

Make sure to use the same Go program on both compression and decompression for compatibility.

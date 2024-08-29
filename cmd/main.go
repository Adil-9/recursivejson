package main

import (
	"flag"
	"log"
	"recursiveZip/internals/decompresser"
	"recursiveZip/internals/writer"
)

var (
	dirName    = flag.String("dir", "", "select directory to compress")
	fileName   = flag.String("f", "filename", "select file Name, default is filename")
	deCompress = flag.Bool("d", false, "select to decompress")
)

func main() {
	flag.Parse()

	if *dirName == "" && !*deCompress {
		log.Fatal("directory was not selected!")
		return
	}

	if !*deCompress {
		err := writer.Write(*dirName, *fileName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := decompresser.Decompress(*fileName)
		if err != nil {
			log.Fatal(err)
		}
	}
}

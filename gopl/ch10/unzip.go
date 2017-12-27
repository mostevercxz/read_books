package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func decompress(path string) error {
	files, err := zip.OpenReader(path)

	if err != nil {
		return err
	}
	defer files.Close()

	for _, f := range files.File {
		// get file content in zip files
		contentf, err := f.Open()
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse file failed:%s,%s\n", f.Name, err)
			continue
		}
		defer contentf.Close()

		// mkdir the parent directory of the file
		dirName := filepath.Dir(f.Name)
		err = os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mkdir failed:%s,%s\n", dirName, err)
			continue
		}

		// check if the file is directory, if it is, continue; otherwise,go on
		finfo, err := os.Stat(f.Name)
		if err == nil && finfo.Mode().IsDir() {
			continue
		}

		// create file and copy content to dest file
		destf, err := os.Create(f.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "create file failed:%s,%s\n", f.Name, err)
			continue
		}

		_, err = io.Copy(destf, contentf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Copy file failed:%s,%s\n", f.Name, err)
		}
	}

	return nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please specify the filename to unzip or untar")
		os.Exit(1)
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Println("File not exists, please check!!")
		os.Exit(1)
	}
	decompress(os.Args[1])
}

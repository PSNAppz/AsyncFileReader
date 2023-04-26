package main

import (
	"async-file-reader/pkg/filereader"
	"fmt"
)

func main() {
	filePaths := []string{
		"files/file1.txt",
		"files/file2.txt",
		"files/file3.txt",
		"files/file4.txt",
	}

	results, err := filereader.ReadFilesAsync(filePaths)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, content := range results {
		fmt.Printf("File %d content: %s\n", i+1, content)
	}
}

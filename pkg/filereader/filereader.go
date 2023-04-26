package filereader

import (
	"os"
	"sync"
)

func ReadFilesAsync(filePaths []string) ([]string, error) {
	var wg sync.WaitGroup
	fileContents := make([]string, len(filePaths))
	errors := make(chan error, len(filePaths))

	for i, path := range filePaths {
		wg.Add(1)
		go func(index int, filePath string) {
			defer wg.Done()

			content, err := os.ReadFile(filePath)
			if err != nil {
				errors <- err
				return
			}

			fileContents[index] = string(content)
		}(i, path)
	}

	wg.Wait()
	close(errors)

	if len(errors) > 0 {
		return nil, <-errors
	}

	return fileContents, nil
}

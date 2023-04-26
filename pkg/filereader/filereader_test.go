package filereader

import (
	"path/filepath"
	"testing"
)

func TestReadFilesAsync(t *testing.T) {
	absPath, err := filepath.Abs("../../files")
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	filePaths := []string{
		filepath.Join(absPath, "file1.txt"),
		filepath.Join(absPath, "file2.txt"),
		filepath.Join(absPath, "file3.txt"),
		filepath.Join(absPath, "file4.txt"),
	}

	results, err := ReadFilesAsync(filePaths)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(results) != len(filePaths) {
		t.Errorf("Expected %d results, but got %d", len(filePaths), len(results))
	}
}

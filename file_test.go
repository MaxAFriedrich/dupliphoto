package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetPaths(t *testing.T) {
	// Create a temporary directory and some files for testing
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.png")
	file3 := filepath.Join(tempDir, "file3.jpg")
	err := os.WriteFile(file1, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file1: %v", err)
	}
	err = os.WriteFile(file2, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file2: %v", err)
	}
	err = os.WriteFile(file3, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file3: %v", err)
	}

	// Call the function to be tested
	allPaths := getPaths(tempDir)

	// Check if the expected paths are returned
	expectedPaths := []string{file1, file2, file3}
	for _, path := range expectedPaths {
		found := false
		for _, p := range allPaths {
			if path == p {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected path not found: %s", path)
		}
	}
}

func TestIsImage(t *testing.T) {
	// Test with different image extensions
	imageExtensions := []string{
		"jpg", "jpeg", "png", "gif", "bmp", "tiff", "tif", "webp",
		"jfif", "jxr", "hdp", "wdp", "ico", "svg", "heic", "heif",
	}

	for _, ext := range imageExtensions {
		path := "/path/to/image." + ext
		if !isImage(path) {
			t.Errorf("Expected %s to be recognized as an image", ext)
		}
	}

	// Test with a non-image extension
	nonImageExtension := "txt"
	path := "/path/to/non-image." + nonImageExtension
	if isImage(path) {
		t.Errorf("Expected %s to be recognized as a non-image", nonImageExtension)
	}
}

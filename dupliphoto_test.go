package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// Helper function to compare two slices of strings
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Test outputHashes function
func TestOutputHashes(t *testing.T) {
	// Test data
	hashes := [][]string{
		{"hash1", "file1"},
		{"hash2", "file2"},
	}
	path := "test_output.csv"

	// Call the function
	outputHashes(hashes, path)

	// Read the CSV file and compare its content with the original data
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("Error reading CSV file: %v", err)
	}

	expectedContent := "hash1,file1\nhash2,file2\n"
	if string(fileContent) != expectedContent {
		t.Errorf("OutputHashes: unexpected CSV content. Expected: %s, Got: %s", expectedContent, fileContent)
	}

	// Clean up
	err = os.Remove(path)
	if err != nil {
		t.Fatalf("Error cleaning up test file: %v", err)
	}
}

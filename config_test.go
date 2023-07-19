package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	// Create a temporary YAML file for testing
	tempFile, err := ioutil.TempFile("", "config_test")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Define some test YAML content
	testYAML := `
blocks:
  - target: target1
    sources:
      - source1
      - source2
  - target: target2
    sources:
      - source3
      - source4
`

	// Write test YAML content to the temporary file
	_, err = tempFile.WriteString(testYAML)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	// Close the file after writing
	tempFile.Close()

	// Call the getConfig function to read the temporary file
	config := getConfig(tempFile.Name())

	// Define the expected result
	expectedConfig := Config{
		Blocks: []Block{
			{Target: "target1", Sources: []string{"source1", "source2"}},
			{Target: "target2", Sources: []string{"source3", "source4"}},
		},
	}

	// Compare the actual result with the expected result using the testify/assert package
	assert.Equal(t, expectedConfig, config, "getConfig returned unexpected config")
}

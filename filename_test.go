package main

import (
	"reflect"
	"testing"
)

func TestGetStandardName(t *testing.T) {
	tests := []struct {
		filename      string
		expectedOut   standardName
		expectedFound bool
	}{
		// Test cases where the filename matches the expected pattern
		{
			filename: "2023_05_06_0001.jpeg",
			expectedOut: standardName{
				year:     2023,
				month:    5,
				day:      6,
				index:    1,
				fullDate: "2023_05_06",
			},
			expectedFound: true,
		},
		{
			filename: "2021_12_25_9999.png",
			expectedOut: standardName{
				year:     2021,
				month:    12,
				day:      25,
				index:    9999,
				fullDate: "2021_12_25",
			},
			expectedFound: true,
		},
		// Test cases where the filename does not match the expected pattern
		{
			filename:      "random_filename.txt",
			expectedOut:   standardName{},
			expectedFound: false,
		},
		{
			filename:      "2023_07_20.jpeg", // Missing index
			expectedOut:   standardName{},
			expectedFound: false,
		},
	}

	for _, test := range tests {
		out, found := getStandardName(test.filename)

		if found != test.expectedFound {
			t.Errorf("Expected found=%v, but got found=%v for filename=%s", test.expectedFound, found, test.filename)
		}

		if !reflect.DeepEqual(out, test.expectedOut) {
			t.Errorf("Expected out=%v, but got out=%v for filename=%s", test.expectedOut, out, test.filename)
		}
	}
}

func TestGetMaxIndex(t *testing.T) {
	testData := []string{
		"/path/to/2023_05_06_0001.jpeg", // Valid filename
		"/path/to/2021_12_25_0020.png",  // Valid filename
		"/path/to/2019_10_31_0200.docx", // Valid filename
		"/path/to/2021_12_25_0123.pdf",  // Valid filename

		"/path/to/random_file_without_ext",      // Invalid filename (missing extension)
		"/path/to/invalid_date_format_2023.pdf", // Invalid filename (invalid date format)
		"/path/to/2020_11_28_999999.png",        // Invalid filename (index too large)
		"/path/to/random_filename.txt",          // Invalid filename
		"/path/to/2023_07_20.jpeg",              // Invalid filename (missing index)
	}
	if getMaxIndex(testData, "2021_12_25") != 123 {
		t.Error()
	}
}

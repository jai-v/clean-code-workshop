package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

// Approach-3 with go testing library with multiple multiple tests scenerios and name for each scenerio(Table Driven Test)
func TestToReadableSizeMultipleApproachTwo(t *testing.T) {
	tt := []struct {
		name     string
		input    int64
		expected string
	}{
		{"byte_return", 125, "125 B"},
		{"kilobyte_return", 1010, "1 KB"},
		{"megabyte_return", 1988909, "1 MB"},
		{"gigabyte_return", 29121988909, "29 GB"},
		{"gigabyte_return", 890929121988909, "890 TB"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output := toReadableSize(tc.input)
			if output != tc.expected {
				t.Errorf("input %d, unexpected output: %s", tc.input, output)
			}
		})
	}
}

func TestTraverseDir(t *testing.T) {
	hashes := map[string]string{}
	duplicates := map[string]string{}
	var dupeSize int64
	directory := "./duplicates_files_directory"
	entries, _ := ioutil.ReadDir(directory)
	traverseDir(hashes, duplicates, &dupeSize, entries, directory)
	expectedHashes := map[string]string{
		"2dc2a49f5873c9fe21e3ce737a7da25a0c600ac8": "duplicates_files_directory/one copy.txt",
		"907a9be60387970483664a0237a9eb0e9b9590a5": "duplicates_files_directory/two copy.txt",
	}
	lenExpectedHashes := len(expectedHashes)
	expectedDuplicates := map[string]string{
		"duplicates_files_directory/one copy.txt": "duplicates_files_directory/one.txt",
		"duplicates_files_directory/two copy.txt": "duplicates_files_directory/two.txt",
	}
	lenExpectedDuplicates := len(expectedDuplicates)

	if lenExpectedHashes != len(hashes) {
		t.Errorf("got %v, wanted %v", len(hashes), lenExpectedHashes)
	}

	if lenExpectedDuplicates != len(duplicates) {
		t.Errorf("got %v, wanted %v", len(duplicates), lenExpectedDuplicates)
	}

	assert.Equal(t, expectedHashes, hashes)
	assert.Equal(t, expectedDuplicates, duplicates)

}

package main

import (
	directoryOptions "clean-code-workshop/directory"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Approach-3 with go testing library with multiple multiple tests scenerios and name for each scenerio(Table Driven Test)
func TestToReadableSizeMultipleApproachTwo(t *testing.T) {
	tt := []struct {
		name     string
		input    int64
		expected string
	}{
		{"byte_return", 125, "125.00 B"},
		{"kilobyte_return", 1045, "1.02 KB"},
		{"megabyte_return", 1988909, "1.90 MB"},
		{"gigabyte_return", 29121988909, "27.12 GB"},
		{"gigabyte_return", 890929121988909, "810.30 TB"},
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
	d := directoryOptions.NewDuplicates()
	directory := "./duplicates_files_directory"
	d.TraverseDir(directory)
	expectedHashes := map[string]string{
		"2dc2a49f5873c9fe21e3ce737a7da25a0c600ac8": "duplicates_files_directory/one copy.txt",
		"907a9be60387970483664a0237a9eb0e9b9590a5": "duplicates_files_directory/two copy.txt",
		"da39a3ee5e6b4b0d3255bfef95601890afd80709": "duplicates_files_directory/subdirectory/testfile",
	}
	lenExpectedHashes := len(expectedHashes)
	expectedDuplicates := map[string]string{
		"duplicates_files_directory/one copy.txt": "duplicates_files_directory/one.txt",
		"duplicates_files_directory/two copy.txt": "duplicates_files_directory/two.txt",
	}
	lenExpectedDuplicates := len(expectedDuplicates)

	if lenExpectedHashes != len(d.Hashes) {
		t.Errorf("got %v, wanted %v", len(d.Hashes), lenExpectedHashes)
	}

	if lenExpectedDuplicates != len(d.DuplicateSlice) {
		t.Errorf("got %v, wanted %v", len(d.DuplicateSlice), lenExpectedDuplicates)
	}

	assert.Equal(t, expectedHashes, d.Hashes)
	assert.Equal(t, expectedDuplicates, d.DuplicateSlice)

}

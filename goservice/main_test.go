package main

import "testing"

func TestTablesortBooks(t *testing.T) {
	var tests = []struct {
		input1   []string
		input2   []book
		expected []book
	}{
		{
			[]string{"Of Mice and Men", "The Pearl", "East of Eden"},
			[]book{
				{ID: "1", Title: "Of Mice and Men", Author: "John Steinbeck", Genre: "fiction"},
				{ID: "2", Title: "The Pearl", Author: "John Steinbeck", Genre: "fiction"},
				{ID: "3", Title: "East of Eden", Author: "John Steinbeck", Genre: "fiction"},
			},
			[]book{
				{ID: "3", Title: "East of Eden", Author: "John Steinbeck", Genre: "fiction"},
				{ID: "1", Title: "Of Mice and Men", Author: "John Steinbeck", Genre: "fiction"},
				{ID: "2", Title: "The Pearl", Author: "John Steinbeck", Genre: "fiction"},
			},
		},
	}

	for _, test := range tests {
		output := sortBooks(test.input1, test.input2)
		for i := 0; i < len(output); i++ {
			if output[i] != test.expected[i] {
				t.Error("Test Failed (expected : actual) =>", test.expected, output)
				break
			}
		}
	}
}

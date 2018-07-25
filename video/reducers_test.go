package video

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchReducer(t *testing.T) {
	tests := map[string]struct {
		Input    []Video
		Search   string
		Expected []Video
	}{
		"empty input": {
			Input:    []Video{},
			Search:   "a",
			Expected: []Video{},
		},
		"empty search": {
			Input:    []Video{{Title: "alpha"}, {Title: "beta"}, {Title: "charlie"}},
			Search:   "",
			Expected: []Video{},
		},
		"exact": {
			Input:    []Video{{Title: "alpha"}, {Title: "beta"}, {Title: "charlie"}},
			Search:   "beta",
			Expected: []Video{{Title: "beta"}},
		},
		"all": {
			Input:    []Video{{Title: "alpha"}, {Title: "beta"}, {Title: "charlie"}},
			Search:   "a",
			Expected: []Video{{Title: "alpha"}, {Title: "beta"}, {Title: "charlie"}},
		},
		"subset": {
			Input:    []Video{{Title: "alpha"}, {Title: "beta"}, {Title: "charlie"}},
			Search:   "e",
			Expected: []Video{{Title: "beta"}, {Title: "charlie"}},
		},
		"none": {
			Input:    []Video{{Title: "alpha"}, {Title: "beta"}, {Title: "charlie"}},
			Search:   "delta",
			Expected: []Video{},
		},
		"series": {
			Input:    []Video{{Title: "alpha", Series: "delta"}, {Title: "beta"}, {Title: "charlie"}},
			Search:   "delta",
			Expected: []Video{{Title: "alpha", Series: "delta"}},
		},
		"multiple words": {
			Input:    []Video{{Title: "alpha"}, {Title: "beta"}, {Title: "charlie"}},
			Search:   "alpha beta",
			Expected: []Video{{Title: "alpha"}, {Title: "beta"}},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reduce := NewSearchReducer(tc.Search)
			assert.Equal(t, tc.Expected, reduce(tc.Input))
		})
	}
}

func TestLimitReducer(t *testing.T) {
	tests := map[string]struct {
		Input    []Video
		Limit    int
		Expected []Video
	}{
		"empty": {
			Input:    []Video{},
			Limit:    3,
			Expected: []Video{},
		},
		"zero": {
			Input:    []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
			Limit:    0,
			Expected: []Video{},
		},
		"less_than": {
			Input:    []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
			Limit:    2,
			Expected: []Video{{Title: "a"}, {Title: "b"}},
		},
		"equal": {
			Input:    []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
			Limit:    3,
			Expected: []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
		},
		"greater_than": {
			Input:    []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
			Limit:    4,
			Expected: []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reduce := NewLimitReducer(tc.Limit)
			assert.Equal(t, tc.Expected, reduce(tc.Input))
		})
	}
}

func TestStartReducer(t *testing.T) {
	tests := map[string]struct {
		Input    []Video
		Start    int
		Expected []Video
	}{
		"empty": {
			Input:    []Video{},
			Start:    3,
			Expected: []Video{},
		},
		"zero": {
			Input:    []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
			Start:    0,
			Expected: []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
		},
		"less_than": {
			Input:    []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
			Start:    1,
			Expected: []Video{{Title: "b"}, {Title: "c"}},
		},
		"equal": {
			Input:    []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
			Start:    3,
			Expected: []Video{},
		},
		"greater_than": {
			Input:    []Video{{Title: "a"}, {Title: "b"}, {Title: "c"}},
			Start:    4,
			Expected: []Video{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reduce := NewStartReducer(tc.Start)
			assert.Equal(t, tc.Expected, reduce(tc.Input))
		})
	}
}

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
			Input:    []Video{{Name: "alpha"}, {Name: "beta"}, {Name: "charlie"}},
			Search:   "",
			Expected: []Video{},
		},
		"exact": {
			Input:    []Video{{Name: "alpha"}, {Name: "beta"}, {Name: "charlie"}},
			Search:   "beta",
			Expected: []Video{{Name: "beta"}},
		},
		"all": {
			Input:    []Video{{Name: "alpha"}, {Name: "beta"}, {Name: "charlie"}},
			Search:   "a",
			Expected: []Video{{Name: "alpha"}, {Name: "beta"}, {Name: "charlie"}},
		},
		"subset": {
			Input:    []Video{{Name: "alpha"}, {Name: "beta"}, {Name: "charlie"}},
			Search:   "e",
			Expected: []Video{{Name: "beta"}, {Name: "charlie"}},
		},
		"none": {
			Input:    []Video{{Name: "alpha"}, {Name: "beta"}, {Name: "charlie"}},
			Search:   "delta",
			Expected: []Video{},
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
			Input:    []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
			Limit:    0,
			Expected: []Video{},
		},
		"less_than": {
			Input:    []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
			Limit:    2,
			Expected: []Video{{Name: "a"}, {Name: "b"}},
		},
		"equal": {
			Input:    []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
			Limit:    3,
			Expected: []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
		},
		"greater_than": {
			Input:    []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
			Limit:    4,
			Expected: []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
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
			Input:    []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
			Start:    0,
			Expected: []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
		},
		"less_than": {
			Input:    []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
			Start:    1,
			Expected: []Video{{Name: "b"}, {Name: "c"}},
		},
		"equal": {
			Input:    []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
			Start:    3,
			Expected: []Video{},
		},
		"greater_than": {
			Input:    []Video{{Name: "a"}, {Name: "b"}, {Name: "c"}},
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

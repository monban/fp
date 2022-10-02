package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/monban/fp"
)

func TestQuicksort(t *testing.T) {
	in := []int{1, 9, 2, 8, 7, 3, 6, 4, 5}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	is := is.New(t)
	out := fp.QuickSortSlice(in)
	is.Equal(out, expected)
}

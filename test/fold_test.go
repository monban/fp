package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/monban/fp"
)

func TestFold(t *testing.T) {
	is := is.New(t)
	foo := []int{1, 2, 3, 4}
	out := fp.Fold(foo, sum)
	is.Equal(out, 10)
}

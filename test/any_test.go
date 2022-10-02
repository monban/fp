package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/monban/fp"
)

func TestAny(t *testing.T) {
	is := is.New(t)
	foo := []int{2, 3, 4}
	is.Equal(fp.Any(foo, isOdd), true)
}

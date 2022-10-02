package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/monban/fp"
)

func TestAllTrue(t *testing.T) {
	is := is.New(t)

	foo := []int{1, 3, 5}
	is.Equal(fp.All(foo, isOdd), true)
}

func TestAllFalse(t *testing.T) {
	is := is.New(t)

	foo := []int{1, 4, 5}
	is.Equal(fp.All(foo, isOdd), false)
}

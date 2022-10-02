package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/monban/fp"
)

func TestPartition(t *testing.T) {
	is := is.New(t)
	foo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	ae := []int{1, 2, 3, 4, 5, 6}
	be := []int{7, 8, 9, 10, 11, 12}
	pred := lowerThan(7)
	ao, bo := fp.Partition(foo, pred)
	is.Equal(ae, ao)
	is.Equal(be, bo)
}

func lowerThan(i int) func(int) bool {
	return func(j int) bool {
		return j < i
	}
}

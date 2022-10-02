package tests

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
	"github.com/monban/fp"
)

func TestMapReturnsArraySameLength(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{1, 2},
		{0, 1, 2},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
			is := is.New(t)
			out := fp.Map(test, doNothing)
			is.Equal(out, test)
		})
	}
}

func TestMapCallsProcedureForEachElement(t *testing.T) {
	foo := [][]int{
		{},
		{0},
		{1, 2, 3, 4, 5, 6},
	}
	for c, testData := range foo {
		t.Run(fmt.Sprintf("case #%d", c), func(t *testing.T) {
			is := is.New(t)
			var i int
			fp.Map(testData, func(j int) int {
				i++
				return 0
			})
			is.Equal(i, len(testData))
		})
	}
}

func TestMapAddOne(t *testing.T) {
	tests := []struct {
		i []int
		e []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 3, 4, 5}},
		{[]int{-1, -2, -3, -4}, []int{0, -1, -2, -3}},
		{[]int{100, 2000, 30000, 400000}, []int{101, 2001, 30001, 400001}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
			is := is.New(t)
			out := fp.Map(test.i, addOne)
			is.Equal(out, test.e)
		})
	}
}

func TestMapReverseString(t *testing.T) {
	test := []string{"hannah", "moon", "apple", "cotton"}
	expected := []string{"hannah", "noom", "elppa", "nottoc"}

	is := is.New(t)

	out := fp.Map(test, reverseString)
	is.Equal(out, expected)
}

func doNothing(i int) int {
	return i
}

func addOne(i int) int {
	return i + 1
}

func reverseString(a string) string {
	var b string
	for _, c := range a {
		b = string(c) + b
	}
	return b
}

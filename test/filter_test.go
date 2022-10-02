package tests

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
	"github.com/monban/fp"
)

func TestFilterCallsPredicateForEachElement(t *testing.T) {
	foo := [][]int{
		{},
		{0},
		{1, 2, 3, 4, 5, 6},
	}
	for c, testData := range foo {
		t.Run(fmt.Sprintf("case #%d", c), func(t *testing.T) {
			is := is.New(t)
			var i int
			fp.FilterSlice(testData, func(j int) bool {
				i++
				return true
			})
			is.Equal(i, len(testData))
		})
	}
}

func TestFilterReturnsEmptyWhenAllFalse(t *testing.T) {
	foo := [][]int{
		{},
		{0},
		{1, 2, 3, 4, 5, 6},
	}
	for c, testData := range foo {
		t.Run(fmt.Sprintf("case #%d", c), func(t *testing.T) {
			is := is.New(t)
			out := fp.FilterSlice(testData, func(j int) bool {
				return false
			})
			is.Equal(0, len(out))
		})
	}
}

func TestFilterSameLengthWhenTrue(t *testing.T) {
	foo := [][]int{
		{},
		{0},
		{1, 2, 3, 4, 5, 6},
	}
	for c, testData := range foo {
		t.Run(fmt.Sprintf("case #%d", c), func(t *testing.T) {
			is := is.New(t)
			out := fp.FilterSlice(testData, func(j int) bool {
				return true
			})
			is.Equal(len(testData), len(out))
		})
	}
}

func TestFilterOdd(t *testing.T) {
	foo := []struct{ in, expected []int }{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{7}, []int{7}},
		{[]int{8}, []int{}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 3, 5}},
	}
	for c, testData := range foo {
		t.Run(fmt.Sprintf("case #%d", c), func(t *testing.T) {
			is := is.New(t)
			out := fp.FilterSlice(testData.in, isOdd)
			is.Equal(out, testData.expected)
		})
	}
}

func TestFilterStringStartsWith(t *testing.T) {
	someStrings := []string{"apple", "green", "great", "banana", ""}
	tests := []struct {
		letter   string
		in       []string
		expected []string
	}{
		{"a", someStrings, []string{"apple"}},
		{"g", someStrings, []string{"green", "great"}},
		{"b", someStrings, []string{"banana"}},
	}

	for _, testData := range tests {
		t.Run("case #%d", func(t *testing.T) {
			is := is.New(t)
			out := fp.FilterSlice(testData.in, startsWith(testData.letter))
			is.Equal(out, testData.expected)
		})
	}
}

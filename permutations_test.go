// Package permutations is library of functions to create permutations
//
// Ralf Poeppel 2018
package permutations

import (
	"reflect"
	"testing"
)

func TestEnumerateInterval(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 3}, []int{1, 2, 3}},
		{[]int{2, 7}, []int{2, 3, 4, 5, 6, 7}},
	}
	for _, c := range cases {
		got := EnumerateInterval(c.in[0], c.in[1])
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("EnumerateInterval(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFaculty(t *testing.T) {
	cases := []struct {
		in   int
		want int64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{9, 362880},
		{11, 39916800},
		{20, 2432902008176640000},
		//{21, 1}, // produces overflow over max int64
	}
	for _, c := range cases {
		got := Faculty(c.in)
		if got != c.want {
			t.Errorf("Faculty(%v) == %v, want %v", c.in, got, c.want)
		}
	}

}

func TestFacultyRange(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{-1, "underflow for n=-1, out of range of definition"},
		{21, "overflow at i=21, out of range of implementation"},
		{22, "overflow at i=21, out of range of implementation"},
	}
	for _, c := range cases {
		func(in int, want string) {
			defer func(in int, want string) {
				r := recover()
				if r == nil {
					t.Errorf("Faculty(%v) did not panic", in)
				} else {
					if r != want {
						t.Errorf("Faculty(%v) == Panic(%v), want Panic(%v)", in, r, want)
					}
				}
			}(c.in, c.want)
			Faculty(c.in)
		}(c.in, c.want)
	}
}

func TestGeneratePermutations(t *testing.T) {
	cases := []struct {
		in   []interface{}
		want [][]interface{}
	}{
		{[]interface{}{1}, [][]interface{}{{1}}},
		{[]interface{}{1, 2}, [][]interface{}{{1, 2}, {2, 1}}},
		{[]interface{}{'a', 'b'}, [][]interface{}{{'a', 'b'}, {'b', 'a'}}},
		{[]interface{}{"alpha", "beta"}, [][]interface{}{{"alpha", "beta"}, {"beta", "alpha"}}},
		{[]interface{}{1, 2, 3}, [][]interface{}{{1, 2, 3}, {2, 1, 3},
			{3, 1, 2}, {1, 3, 2},
			{3, 2, 1}, {2, 3, 1}}},
		{[]interface{}{1, 2, 3, 4},
			[][]interface{}{{1, 2, 3, 4}, {2, 1, 3, 4}, {3, 1, 2, 4}, {1, 3, 2, 4}, {3, 2, 1, 4}, {2, 3, 1, 4},
				{4, 1, 2, 3}, {1, 4, 2, 3}, {1, 2, 4, 3},
				{4, 2, 1, 3}, {2, 4, 1, 3}, {2, 1, 4, 3},
				{4, 3, 1, 2}, {3, 4, 1, 2}, {3, 1, 4, 2},
				{4, 1, 3, 2}, {1, 4, 3, 2}, {1, 3, 4, 2},
				{4, 3, 2, 1}, {3, 4, 2, 1}, {3, 2, 4, 1},
				{4, 2, 3, 1}, {2, 4, 3, 1}, {2, 3, 4, 1}}},
	}
	for _, c := range cases {
		got := GeneratePermutations(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("GeneratePermutations(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

// verify first and last element
func TestGeneratePermutationsLast(t *testing.T) {
	cases := []struct {
		in   []interface{}
		want [][]interface{}
	}{
		{[]interface{}{1}, [][]interface{}{{1}, {1}}},
		{[]interface{}{1, 2}, [][]interface{}{{1, 2}, {2, 1}}},
		{[]interface{}{1, 2, 3}, [][]interface{}{
			{1, 2, 3}, {2, 3, 1}}},
		{[]interface{}{1, 2, 3, 4}, [][]interface{}{
			{1, 2, 3, 4}, {2, 3, 4, 1}}},
		{[]interface{}{1, 2, 3, 4, 5}, [][]interface{}{
			{1, 2, 3, 4, 5}, {2, 3, 4, 5, 1}}},
		{[]interface{}{1, 2, 3, 4, 5, 6}, [][]interface{}{
			{1, 2, 3, 4, 5, 6}, {2, 3, 4, 5, 6, 1}}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7}, [][]interface{}{
			{1, 2, 3, 4, 5, 6, 7}, {2, 3, 4, 5, 6, 7, 1}}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8},
			[][]interface{}{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{2, 3, 4, 5, 6, 7, 8, 1}}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[][]interface{}{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{2, 3, 4, 5, 6, 7, 8, 9, 1}}},
	}
	for _, c := range cases {
		got := GeneratePermutations(c.in)
		n := len(got) - 1
		if !reflect.DeepEqual(got[0], c.want[0]) {
			t.Errorf("GeneratePermutations(%v) == [0]%v, want[0] %v", c.in, got[0], c.want[0])
		}
		if !reflect.DeepEqual(got[n], c.want[1]) {
			t.Errorf("GeneratePermutations(%v) == [n]%v, want[1] %v", c.in, got[n], c.want[1])
		}
	}
}

func TestGeneratePermutationsSortedAll(t *testing.T) {
	cases := []struct {
		in   []interface{}
		want [][]interface{}
	}{
		{[]interface{}{1}, [][]interface{}{{1}}},
		{[]interface{}{1, 2}, [][]interface{}{{1, 2}, {2, 1}}},
		{[]interface{}{'a', 'b'}, [][]interface{}{{'a', 'b'}, {'b', 'a'}}},
		{[]interface{}{"alpha", "beta"}, [][]interface{}{{"alpha", "beta"}, {"beta", "alpha"}}},
		{[]interface{}{1, 2, 3}, [][]interface{}{
			{1, 2, 3}, {1, 3, 2},
			{2, 1, 3}, {2, 3, 1},
			{3, 1, 2}, {3, 2, 1}}},
		{[]interface{}{1, 2, 3, 4}, [][]interface{}{
			{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4},
			{1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2},
			{2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4},
			{2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1},
			{3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4},
			{3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1},
			{4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3},
			{4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1}}},
	}
	for _, c := range cases {
		got := GeneratePermutationsSorted(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("GeneratePermutationsSorted(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

// verify first and last element
func TestGeneratePermutationsSortedLast(t *testing.T) {
	cases := []struct {
		in   []interface{}
		want [][]interface{}
	}{
		{[]interface{}{1}, [][]interface{}{{1}, {1}}},
		{[]interface{}{1, 2}, [][]interface{}{{1, 2}, {2, 1}}},
		{[]interface{}{1, 2, 3}, [][]interface{}{
			{1, 2, 3}, {3, 2, 1}}},
		{[]interface{}{1, 2, 3, 4}, [][]interface{}{
			{1, 2, 3, 4}, {4, 3, 2, 1}}},
		{[]interface{}{1, 2, 3, 4, 5}, [][]interface{}{
			{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}}},
		{[]interface{}{1, 2, 3, 4, 5, 6}, [][]interface{}{
			{1, 2, 3, 4, 5, 6}, {6, 5, 4, 3, 2, 1}}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7}, [][]interface{}{
			{1, 2, 3, 4, 5, 6, 7}, {7, 6, 5, 4, 3, 2, 1}}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8},
			[][]interface{}{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{8, 7, 6, 5, 4, 3, 2, 1}}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[][]interface{}{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{9, 8, 7, 6, 5, 4, 3, 2, 1}}},
	}
	for _, c := range cases {
		got := GeneratePermutationsSorted(c.in)
		n := len(got) - 1
		if !reflect.DeepEqual(got[0], c.want[0]) {
			t.Errorf("GeneratePermutationsSorted(%v) == [0]%v, want[0] %v", c.in, got[0], c.want[0])
		}
		if !reflect.DeepEqual(got[n], c.want[1]) {
			t.Errorf("GeneratePermutationsSorted(%v) == [n]%v, want[1] %v", c.in, got[n], c.want[1])
		}
	}
}

// Package permutations is library of functions to create permutations
//
// Ralf Poeppel 2018
package permutations

import (
	"fmt"
)

// EnumerateInterval enumerate the integers in an intervall, s. sicp chapter 2.2.3
func EnumerateInterval(low, high int) []int {
	len := high - low + 1
	res := make([]int, 0, len)
	for i := low; i <= high; i++ {
		res = append(res, i)
	}
	return res
}

// Faculty calculates the value of faculty(n)
func Faculty(n int) int64 {
	n64 := int64(n)
	res := int64(1)
	if n < 0 {
		panic("underflow for n=" + fmt.Sprint(n) + ", out of range of definition")
	}

	if n == 0 {
		return res
	}
	maxval := int64(^uint64(0)>>1) / n64
	for i := res; i <= n64; i++ {
		if res < maxval {
			res *= i
		} else {
			panic("overflow at i=" + fmt.Sprint(i) + ", out of range of implementation")
		}
	}
	return res
}

// Permutations calculates the permutations of an array of values
func Permutations(values []interface{}) [][]interface{} {
	n := len(values)
	m := Faculty(n)
	n64 := int64(n)

	// example from https://golang.org/doc/effective_go.html#two_dimensional_slices
	// Allocate the top-level slice, the same as before.
	res := make([][]interface{}, m) // One row per permutation
	// Allocate one large slice to hold all the permutation values.
	vals := make([]interface{}, m*n64) // Has type []interface{} even though picture is [][]interface{}.
	// Loop over the rows, slicing each row from the front of the remaining values slice.
	for i := range res {
		res[i], vals = vals[:n], vals[n:]
	}

	permutate(values, res, m, n64)
	return res
}

func permutate(values []interface{}, res [][]interface{}, m, n int64) {
	if n == 1 {
		res[0][0] = values[0]
		return
	}

	mprev := m / n
	nprev := n - 1
	permutate(values, res, mprev, nprev)

	// set last value to current at all rows of inner permutation
	current := values[nprev]
	zero := int64(0) // zero as 64 bit integer
	for i := zero; i < mprev; i++ {
		res[i][nprev] = current
	}

	// creat new rows where current is inserted at all positions except the last
	r := mprev
	for i := zero; i < mprev; i++ {
		for j := zero; j < nprev; j++ {
			l := zero
			for k := zero; k < nprev; k++ {
				if j == l {
					res[r][l] = current
					l++
				}
				res[r][l] = res[i][k]
				l++
			}
			r++
		}
	}
}

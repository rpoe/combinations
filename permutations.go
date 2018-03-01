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

// GeneratePermutations generates the permutations of an array
// of values. The rows of the array are not sorted.
// This is the fastest algorithm in this package.
func GeneratePermutations(values []interface{}) [][]interface{} {
	n := len(values)
	m := Faculty(n)
	n64 := int64(n)

	// example from https://golang.org/doc/effective_go.html#two_dimensional_slices
	// Allocate the top-level slice, the same as before.
	res := make([][]interface{}, m) // One row per permutation
	// Allocate one large slice to hold the whole arrayof all the permutation values.
	vals := make([]interface{}, m*n64) // Has type []interface{} even though array is [][]interface{}.
	// Loop over the rows, slicing each row from the front of the remaining values slice.
	for i := range res {
		res[i], vals = vals[:n], vals[n:]
	}

	permutate(values, res, m, n64)
	return res
}

// permutate generates m permutations of an array of size n.
// It works recursively. Each recursive call generates the
// previous permutation of values[:n-1].
// The previous permutation is generated in the result array
// at the final position, so there is no more operation on
// its values.
func permutate(values []interface{}, res [][]interface{}, m, n int64) {
	// check for last permutation
	if n <= 1 {
		res[0][0] = values[0]
		return
	}

	// generate inner permutation
	mprev := m / n
	nprev := n - 1
	permutate(values, res, mprev, nprev)

	// set last value to current at all rows of inner permutation
	current := values[nprev]
	zero := int64(0) // zero as 64 bit integer
	for i := zero; i < mprev; i++ {
		res[i][nprev] = current
	}

	// generate remaining rows by inserting current at
	// all positions except the last position
	// see Donal E. Knuth, The art of computer programming
	// Volume 4, Fascicle 2, 7.2.1.2, (3), p. 41, 2005
	r := mprev                      // row index begins with first after inner permutation
	for i := zero; i < mprev; i++ { // source is each row of inner permutation
		for j := zero; j < nprev; j++ { // loop over insertions
			l := zero                       // insertion position
			for k := zero; k < nprev; k++ { // loop over all values of row
				if j == l { // insert
					res[r][l] = current
					l++
				}
				// copy from inner
				res[r][l] = res[i][k]
				l++
			}
			r++ // advance result row
		}
	}
}

// GeneratePermutationsSorted generate an array with all permutations
// of an array of values and maintain the sort order of the array.
func GeneratePermutationsSorted(values []interface{}) [][]interface{} {
	n := len(values)
	m := Faculty(n)
	n64 := int64(n)

	// example from https://golang.org/doc/effective_go.html#two_dimensional_slices
	// Allocate the top-level slice, the same as before.
	res := make([][]interface{}, m) // One row per permutation
	// Allocate one large slice to hold all the permutation values.
	vals := make([]interface{}, m*n64) // Has type []interface{} even though array is [][]interface{}.
	// Loop over the rows, slicing each row from the front of the remaining values slice.
	for i := range res {
		res[i], vals = vals[:n], vals[n:]
	}

	addPermutationsOfSubArray(values, res, m, n64)
	return res
}

// addPermutationsOfSubArray add the permutationsnof the sub array to res
func addPermutationsOfSubArray(values []interface{}, res [][]interface{}, m, n int64) {
	if n == 1 {
		res[0][0] = values[0]
		return
	}

	zero := int64(0) // zero as 64 bit integer

	mprev := m / n
	nprev := n - 1
	subres := make([][]interface{}, mprev) // One row per permutation
	// set reference of rows in res where subres should be stored
	for i, j := zero, m-mprev; i < mprev; i, j = i+1, j+1 {
		subres[i] = res[j][1:]
	}

	addPermutationsOfSubArray(values, subres, mprev, nprev)

	// set first value to current at all rows of inner permutation
	current := values[nprev]
	for j := m - mprev; j < m; j++ {
		res[j][0] = current
	}

	// Generate previous rows, a set for each previous value
	r := m - mprev
	r--
	k := m
	k--
	nprev--
	for nb := nprev; nb >= 0; nb-- { // loop over sets
		next := values[nb]
		for i := zero; i < mprev; i++ { // loop rows of block
			for l := zero; l < n; l++ {
				switch {
				case l == 0:
					res[r][l] = next
				case res[k][l] == next:
					res[r][l] = current
				default:
					res[r][l] = res[k][l]
				}
			}
			r--
			k--
		}
		current = next
	}
}

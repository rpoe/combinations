# permutations
Go functions for fast generation of the array of permutations of a set.
This are implementations having only n! assignments accessing
the memory and smallest number of calls to make.  All other operation
work on local variables. The implementation works recursively.
The result of a recursion is already at the right position, so
no further operation is needed.

There are two versions: one which should be theorethically be faster,
returning an unsorted array, and one which returns a sorted array.
For larger sets the sorted version shows better benchmark results.

<a href="https://godoc.org/github.com/rpoe/permutations"><img src="https://godoc.org/github.com/rpoe/permutations?status.svg" alt="GoDoc"></a>

This project is licensed under the terms of the MIT license.

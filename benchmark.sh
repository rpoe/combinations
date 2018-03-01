#!/bin/sh
# run benchmark test
go test -v -run TestGeneratePermutationsLast -bench . -count 10
go test -v -run TestGeneratePermutationsSortedLast -bench . -count 10

package random

func flipAndInvertImage(A [][]int) [][]int {
	for ri := 0; ri < len(A); ri++ {
		// swap the each value from the beginning and end
		for ci := 0; ci < len(A[ri])/2; ci++ {
			A[ri][ci], A[ri][len(A[ri])-1-ci] = 1-A[ri][len(A[ri])-1-ci], 1-A[ri][ci]
		}
		// Flip the bit if we have an odd length of values in the row
		if len(A[ri])%2 != 0 {
			A[ri][len(A[ri])/2] = 1 - A[ri][len(A[ri])/2]
		}
	}
	return A
}

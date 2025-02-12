package sliceutil

import "slices"

func rotateNew(s []int) {
	slices.Reverse(s) //  go doc slices.Reverse: Reverse reverses the elements of the slice in place.
}

func rotateOld(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

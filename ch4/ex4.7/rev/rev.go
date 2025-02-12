// Exercise 4.7

package rev

import "slices"

func reverseNew(s []int) {
	slices.Reverse(s)
}

func reverseOld(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

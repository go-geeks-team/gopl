package sliceutil

import "slices"

func RemDup(s []string) {
	for i, elem := range s {
		if i < len(s)-1 && elem == s[i+1] {
			slices.Delete(s, i, i+1)
		}
	}
}

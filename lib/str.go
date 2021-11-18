package lib

import (
	"fmt"
)

func createPrefix(needle string) []uint16 {
	var prefix []uint16 = make([]uint16, len(needle))
	var k uint16
	for i := 1; i < len(prefix); i++ {
		k = prefix[i-1]

		for k > 0 && needle[i] != needle[k] {
			k = prefix[k-1]
		}
		if needle[i] == needle[k] {
			k += 1
		}
		prefix[i] = k
	}
	return prefix
}

func strStr(haystack string, needle string) int {
	// Search needle in haystack
	var prefix []uint16 = createPrefix(needle)
	if len(needle) == 0 {
		return 0
	}
	var i uint16
	var k uint16 = 0
	var pos int = -1

	for i = 0; i < uint16(len(haystack)); i++ {
		for k > 0 && needle[k] != haystack[i] {
			k = prefix[k-1]
		}
		if needle[k] == haystack[i] {
			k += 1
		}
		if k == uint16(len(needle)) {
			pos = int(i - k)
			break
		}
	}

	if k == uint16(len(needle)) {
		return int(i - k + 1)
	}

	return pos
}

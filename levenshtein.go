package levenshtein

import (
	"unicode/utf8"
)

func Get(str1 string, str2 string) int {
	str1Len := utf8.RuneCountInString(str1)
	str2Len := utf8.RuneCountInString(str2)

	// Base cases
	if str1Len == 0 {
		return str2Len
	}

	if str2Len == 0 {
		return str1Len
	}

	if str1 == str2 {
		return 0
	}

	// Two rows
	var curCol int
	var nextCol int
	var i int
	var j int
	var tmp int
	prevRow := make([]int, str2Len+1)
	str2Char := make([]rune, str2Len+1)
	runes1 := []rune(str1)
	runes2 := []rune(str2)

	// initialise previous row
	for i = 0; i < str2Len; i++ {
		prevRow[i] = i
		str2Char[i] = runes2[i]
	}
	prevRow[str2Len] = str2Len

	var strCmp bool
	for i = 0; i < str1Len; i++ {
		nextCol = i + 1
		for j = 0; j < str2Len; j++ {
			curCol = nextCol

			// substution
			strCmp = runes1[i] == str2Char[j]

			nextCol = prevRow[j]
			// Go doesn't have ternaries :(
			if !strCmp {
				nextCol++
			}

			// insertion
			tmp = curCol + 1
			if nextCol > tmp {
				nextCol = tmp
			}

			// deletion
			tmp = prevRow[j+1] + 1
			if nextCol > tmp {
				nextCol = tmp
			}

			// copy current col value into previous (in preparation for next iteration)
			prevRow[j] = curCol
		}

		// copy last col value into previous (in preparation for next iteration)
		prevRow[j] = nextCol
	}
	return nextCol
}

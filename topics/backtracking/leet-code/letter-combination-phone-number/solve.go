package solve

import (
	"strings"
)

var digitToLetter = map[string][]string{
	"1": {},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func Run(digitsStr string) []string {
	if digitsStr == "" {
		return []string{}
	}

	return solve(0, "", strings.Split(digitsStr, ""), []string{})
}

func solve(step int, str string, digits []string, result []string) []string {
	if len(str) == len(digits) {
		return append(result, str)
	}

	digit := digits[step]

	for i := 0; i < len(digitToLetter[digit]); i++ {
		str += digitToLetter[digit][i]
		result = solve(step+1, str, digits, result)
		str = str[:len(str)-1]
	}

	return result
}

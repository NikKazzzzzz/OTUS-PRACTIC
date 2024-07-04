package unpacksrt

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(input string) (string, error) {
	var result strings.Builder
	escape := false

	for i, r := range input {
		if unicode.IsDigit(r) && !escape {
			if i == 0 {
				return "", errors.New("invalid string")
			}

			count, _ := strconv.Atoi(string(r))
			lastRune := result.String()[result.Len()-1]
			result.WriteString(strings.Repeat(string(lastRune), count-1))
		} else if r == '\\' && !escape {
			escape = true
		} else {
			result.WriteRune(r)
			escape = false
		}
	}
	return result.String(), nil
}

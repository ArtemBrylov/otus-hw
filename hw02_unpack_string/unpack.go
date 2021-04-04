package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sb strings.Builder

	for i, val := range s {
		if unicode.IsDigit(rune(val)) {
			if i == 0 || (i >= 1 && unicode.IsDigit(rune(s[i-1]))) {
				return "", ErrInvalidString
			}
		} else {
			if i+1 == len(s) || !unicode.IsDigit(rune(s[i+1])) {
				sb.WriteString(string(val))
			} else if unicode.IsDigit(rune(s[i+1])) {
				intValue, err := strconv.Atoi(string(s[i+1]))
				if err != nil {
					fmt.Println("Error during cast Atoi: " + err.Error())
				}
				sb.WriteString(strings.Repeat(string(val), intValue))
			}
		}
	}
	return sb.String(), nil
}

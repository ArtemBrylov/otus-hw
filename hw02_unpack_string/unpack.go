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
		switch {
		case unicode.IsDigit(val) && i == 0:
			return "", ErrInvalidString
		case unicode.IsDigit(val) && i >= 1 && unicode.IsDigit(rune(s[i-1])):
			return "", ErrInvalidString
		case unicode.IsLetter(val) && i+1 == len(s):
			sb.WriteString(string(val))
		case unicode.IsLetter(val) && !unicode.IsDigit(rune(s[i+1])):
			sb.WriteString(string(val))
		case unicode.IsLetter(val) && unicode.IsDigit(rune(s[i+1])):
			intValue, err := strconv.Atoi(string(s[i+1]))
			if err != nil {
				fmt.Println("Error during cast Atoi: " + err.Error())
			}
			sb.WriteString(strings.Repeat(string(val), intValue))
		}
	}
	return sb.String(), nil
}

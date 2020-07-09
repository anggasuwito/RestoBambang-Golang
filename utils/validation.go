package utils

import (
	"errors"
	"strings"
)

//ValidateInputNotNil ValidateInputNotNil
func ValidateInputNotNil(data ...interface{}) error {
	for _, value := range data {
		switch value {
		case "":
			return errors.New("Data Input Cannot Empty")
		case 0:
			return errors.New("Data Input Cannot 0")
		case nil:
			return errors.New("Input cannot be nil")
		}
	}
	return nil
}

//ValidateInputNotSymbol ValidateInputNotSymbol
func ValidateInputNotSymbol(data ...interface{}) error {
	for _, value := range data {
		if strings.ContainsAny(value.(string), "!@#$%^&*()_+{}|:?><\"") { // containsany yaitu mengecek symbol yg kita tidak mau
			return errors.New("Input cannot containsAny symbol")
		}
	}
	return nil
}

//ValidateInputLenCharacter ValidateInputLenCharacter
func ValidateInputLenCharacter(min, max int, data ...interface{}) error {
	for _, value := range data {
		if len(value.(string)) >= min && len(value.(string)) <= max {
			return errors.New("Input cannot more 20 character")
		}
	}
	return nil
}

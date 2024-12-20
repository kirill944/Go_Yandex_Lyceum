package calculation

import (
	"errors"
	"strings"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrInvalidSyntax  = errors.New("invalid syntax")
)

func Handler_err(exp string) (string, error) {

	if exp == "" {
		return "", ErrInvalidSyntax
	}

	if strings.Count(exp, "(") != strings.Count(exp, ")") {
		return "", ErrInvalidSyntax
	}

	sub := [16]string{
		"*-", "*+", "**", "*/",
		"/-", "/+", "/*", "//",
		"--", "-+", "-*", "-/",
		"+-", "++", "+*", "+/"}

	for _, s := range sub {
		if strings.Count(exp, s) > 0 {
			return "", ErrInvalidSyntax
		}
	}

	if (exp[0] < 48 || exp[0] > 57) &&
		string(exp[0]) != "(" && string(exp[0]) != ")" {
		return "", ErrInvalidSyntax
	}

	if (exp[len(exp)-1] < 48 || exp[len(exp)-1] > 57) &&
		string(exp[len(exp)-1]) != "(" && string(exp[len(exp)-1]) != ")" {
		return "", ErrInvalidSyntax
	}

	for i := 0; i < len(exp); i++ {
		if string(exp[i]) != "*" && string(exp[i]) != "/" && string(exp[i]) != "-" && string(exp[i]) != "+" && string(exp[i]) != "(" && string(exp[i]) != ")" && (exp[i] < 48 || exp[i] > 57) {
			return "", ErrInvalidSyntax
		}
	}

	return exp, nil
}

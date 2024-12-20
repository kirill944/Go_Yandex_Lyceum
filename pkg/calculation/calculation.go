package calculation

import (
	"strconv"
)

func A_move_B(expression string) (float64, error) {
	for i := 0; i < len(expression); i++ {
		if expression[i] < 48 || expression[i] > 57 {
			a, _ := strconv.ParseFloat(string(expression[:i]), 64)
			b, _ := strconv.ParseFloat(string(expression[i+1:]), 64)
			if string(expression[i]) == "+" {
				return a + b, nil
			}
			if string(expression[i]) == "-" {
				return a - b, nil
			}
			if string(expression[i]) == "*" {
				return a * b, nil
			}
			if string(expression[i]) == "/" {
				if b == 0 {
					return 0.0, ErrDivisionByZero
				}
				return a / b, nil
			}
		}
	}
	return 0.0, nil
}

func Calculate(expression string) (float64, error) {
	flag := true
	for flag {
		flag = false
		f := 0
		c := 0
		start := 0
		for i := 0; i < len(expression); i++ {
			if string(expression[i]) == "(" {
				c += 1
				if f == 0 {
					start = i
					f = 1
				}
			}
			if string(expression[i]) == ")" {
				c -= 1
				f = 1
			}
			if c < 0 {
				return 0.0, ErrInvalidSyntax
			}
			if c == 0 && f == 1 {
				exp, err := Calculate(expression[start+1 : i])
				if err != nil {
					return 0.0, err
				}
				expression = expression[:start] + strconv.FormatFloat(exp, 'f', 20, 64) + expression[i+1:]
				flag = true
				break
			}
		}
	}
	flag = true

	for flag {
		flag = false
		for i := 0; i < len(expression); i++ {
			if string(expression[i]) == "*" || string(expression[i]) == "/" {
				start := 0
				for j := i - 1; j >= 0; j-- {
					if (expression[j] > 57 || expression[j] < 48) && string(expression[j]) != "." {
						start = j + 1
						break
					}
				}
				end := len(expression)
				for j := i + 1; j < len(expression); j++ {
					if (expression[j] > 57 || expression[j] < 48) && string(expression[j]) != "." {
						end = j
						break
					}
				}
				res, err := A_move_B(expression[start:end])
				if err != nil {
					return 0.0, err
				}
				expression = expression[:start] + strconv.FormatFloat(res, 'f', 20, 64) + expression[end:]
				flag = true
				break
			}
		}
	}
	flag = true
	for flag {
		flag = false
		for i := 0; i < len(expression); i++ {
			if string(expression[i]) == "-" || string(expression[i]) == "+" {
				start := 0
				for j := i - 1; j >= 0; j-- {
					if (expression[j] > 57 || expression[j] < 48) && string(expression[j]) != "." {
						start = j + 1
						break
					}
				}
				end := len(expression)
				for j := i + 1; j < len(expression); j++ {
					if (expression[j] > 57 || expression[j] < 48) && string(expression[j]) != "." {
						end = j
						break
					}
				}
				res, err := A_move_B(expression[start:end])
				if err != nil {
					return 0.0, err
				}
				expression = expression[:start] + strconv.FormatFloat(res, 'f', 20, 64) + expression[end:]
				flag = true
				break
			}
		}
	}
	res, _ := strconv.ParseFloat(expression, 64)
	return res, nil
}

func Calc(expression string) (float64, error) {
	expression, err := Handler_err(expression)
	if err != nil {
		return 0, err
	}
	return Calculate(expression)
}

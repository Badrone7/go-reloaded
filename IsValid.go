package reloaded

func IsHex(str string) bool {
	if len(str) < 4 {
		return false
	}
	return str == "(hex)" || str == "(hex)\n" || str == "(hex)\t" || str == "(hex) "
}

func IsBin(str string) bool {
	if len(str) < 4 {
		return false
	}
	return str == "(bin)" || str == "(bin)\n" || str == "(bin)\t" || str == "(bin) "
}

func IsUp(str string) bool {
	if len(str) < 4 {
		return false
	}
	runes := []rune(str)
	if runes[len(runes)-1] != ')' {
		runes = runes[:len(runes)-1]
		if runes[len(runes)-1] != ')' {
			runes = runes[:len(runes)-1]
		}
	}
	tester := runes[:4]
	if string(runes) == "(up)" {
		return true
	}
	if string(tester) == "(up," {
		tester = runes[4 : len(runes)-1]
		if tester[0] == ' ' {
			for _, ch := range tester {
				if !((ch >= '0' && ch <= '9') || ch == ' ' || ch == ')') {
					return false
				}
			}
			return true
		}
	}
	return false
}

func IsLow(str string) bool {
	if len(str) < 4 {
		return false
	}
	runes := []rune(str)
	if runes[len(runes)-1] != ')' {
		runes = runes[:len(runes)-1]
		if runes[len(runes)-1] != ')' {
			runes = runes[:len(runes)-1]
		}
	}
	tester := runes[:5]
	if string(runes) == "(low)" {
		return true
	}
	if string(tester) == "(low," {
		tester = runes[5 : len(runes)-1]
		if tester[0] == ' ' {
			for _, ch := range tester {
				if !((ch >= '0' && ch <= '9') || ch == ' ' || ch == ')') {
					return false
				}
			}
			return true
		}
	}
	return false
}

func IsCap(str string) bool {
	if len(str) < 4 {
		return false
	}
	runes := []rune(str)
	if runes[len(runes)-1] != ')' {
		runes = runes[:len(runes)-1]
		if runes[len(runes)-1] != ')' {
			runes = runes[:len(runes)-1]
		}
	}
	tester := runes[:5]
	if string(runes) == "(cap)" {
		return true
	}
	if string(tester) == "(cap," {
		tester = runes[5 : len(runes)-1]
		if tester[0] == ' ' {
			for _, ch := range tester {
				if !((ch >= '0' && ch <= '9') || ch == ' ' || ch == ')') {
					return false
				}
			}
			return true
		}
	}
	return false
}

func IsValid(str string) bool {
	if len(str) < 4 {
		return false
	}
	if IsHex(str) || IsBin(str) || IsUp(str) || IsCap(str) || IsLow(str) {
		return true
	}
	return false
}

package reloaded

import (
	"strconv"
)

func Detect(txt string) int {
	char := txt[:3]
	numstr := ""
	if txt[len(txt)-1] != ')' {
		txt = txt[:len(txt)-1]
		if txt[len(txt)-1] != ')' {
			txt = txt[:len(txt)-1]
		}
	}
	if txt == "(up)" || txt == "(up) " || txt == "(up)\n" || txt == "(up)\t" || txt == "(low)" || txt == "(low) " || txt == "(low)\n" || txt == "(low)\t" || txt == "(cap)" || txt == "(cap) " || txt == "(cap)\n" || txt == "(cap)\t" {
		return 1
	}
	if char == "(up" {
		if txt[len(txt)-1] == ')' {
			char = txt[4 : len(txt)-1]
		} else {
			char = txt[4 : len(txt)-2]
		}
		for _, ch := range char {
			if ch >= '0' && ch <= '9' {
				numstr += string(ch)
				continue
			}
			if ch != ' ' {
				return 0
			}
		}
		num, err := strconv.Atoi(numstr)
		if err != nil {
			return 0
		}
		if num > 0 {
			return num
		}
		return 0
	}
	if char == "(lo" || char == "(ca" {
		if txt[len(txt)-1] == ')' {
			char = txt[5 : len(txt)-1]
		} else {
			char = txt[5 : len(txt)-2]
		}
		for _, ch := range char {
			if ch >= '0' && ch <= '9' {
				numstr += string(ch)
				continue
			}
			if ch != ' ' {
				return 0
			}
		}
		num, err := strconv.Atoi(numstr)
		if err != nil {
			return 0
		}
		if num > 0 {
			return num
		}
		return 0
	}
	return 0
}

package reloaded

import (
	"strconv"
)

func Hex(str string) string {
	newstr := str
	final := ""
	ch := ""
	///////////////////////////////////
	if !(str[0] >= '0' && str[0] <= '9') && !(str[0] >= 'a' && str[0] <= 'f') && !(str[0] >= 'A' && str[0] <= 'F') {
		newstr = str[0:]
		final += string(str[0])
	}
	if !(str[len(str)-1] >= '0' && str[len(str)-1] <= '9') && !(str[len(str)-1] >= 'a' && str[len(str)-1] <= 'f') && !(str[len(str)-1] >= 'A' && str[len(str)-1] <= 'F') {
		newstr = str[:len(str)-1]
		ch = string(str[len(str)-1])
	}
	dec, err := strconv.ParseInt(newstr, 16, 64)
	if err != nil {
		return str
	}
	final += strconv.FormatInt(dec, 10)
	if ch != "" {
		final += ch
	}
	return final
}

func Bin(str string) string {
	newstr := str
	final := ""
	ch := ""
	if !(str[0] >= '0' && str[0] <= '9') && !(str[0] >= 'a' && str[0] <= 'f') && !(str[0] >= 'A' && str[0] <= 'F') {
		newstr = str[0:]
		final += string(str[0])
	}
	if !(str[len(str)-1] >= '0' && str[len(str)-1] <= '9') && !(str[len(str)-1] >= 'a' && str[len(str)-1] <= 'f') && !(str[len(str)-1] >= 'A' && str[len(str)-1] <= 'F') {
		newstr = str[:len(str)-1]
		ch = string(str[len(str)-1])
	}
	dec, err := strconv.ParseInt(newstr, 2, 64)
	if err != nil {
		return str
	}
	final += strconv.FormatInt(dec, 10)
	if ch != "" {
		final += ch
	}
	return final
}

func Cap(str string) string {
	final := ""
	runes := []rune(str)
	if str[0] == '\'' {
		final += string(str[0])
		runes = []rune(str[1:])
	}
	first := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' && first == 0 {
			final += string(runes[i])
			first = 1
			continue
		}
		if (runes[i] >= 'a' && runes[i] <= 'z') && first == 0 {
			final += string(runes[i] - 32)
			first = 1
			continue
		} else if runes[i] >= 'A' && runes[i] <= 'Z' && first == 1 {
			final += string(runes[i] + 32)
		} else {
			final += string(runes[i])
		}
	}
	return final
}

func Low(str string) string {
	final := ""
	runes := []rune(str)
	if str[0] == '\'' {
		final += string(str[0])
		runes = []rune(str[1:])
	}
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			final += string(runes[i] + 32)
			continue
		}
		final += string(runes[i])
	}
	return final
}

func Up(str string) string {
	final := ""
	runes := []rune(str)
	if str[0] == '\'' {
		final += string(str[0])
		runes = []rune(str[1:])
	}
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			final += string(runes[i] - 32)
			continue
		}
		final += string(runes[i])
	}
	return final
}

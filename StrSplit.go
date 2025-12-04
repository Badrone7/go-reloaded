package reloaded

func StringPlitter(text []rune) []string {
	words := []string{}
	word := ""
	count := 0
	for i := 0; i < len(text); i++ {
		if word == "" && (text[i] == ' ' || text[i] == '\t' || text[i] == '\n' || text[i] == '\r') {
			continue
		}
		if text[i] == '(' {
			word += string(text[i])
			count = 1
			continue
		}
		if count == 1 {
			if text[i] == ')' {
				word += string(text[i])
				count = 0
				continue
			}
			word += string(text[i])
			continue
		}
		if text[i] != ' ' && text[i] != '\t' && text[i] != '\n' && text[i] != '\r' {
			word += string(text[i])
			continue
		}
		word += string(text[i])
		if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

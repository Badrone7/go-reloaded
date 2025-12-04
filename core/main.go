package main

import (
	"fmt"
	"io"
	"os"

	"reloaded"
)

/*
	Our Program is meant to Correct a TEXT file from any extra spaces or syntax errors
	like quotes and punctuations, also it does the transformations if necessary
*/

func QuotesFixer(text []rune) string {
	tempcount := 0
	tempstr := ""
	for i := 0; i < len(text); i++ {
		if i == len(text)-1 && tempcount == 1 {
			tempstr += "'"
			break
		}
		if text[i] == '\'' && i == len(text)-1 {
			break
		}
		if text[i] == '\'' && tempcount == 0 && text[i+1] == '\n' {
			tempstr += "'"
			continue
		}
		if i != len(text)-1 && text[i] == '\'' && text[i+1] == '\'' {
			continue
		}
		if text[i] == '\n' && tempcount == 1 {
			tempstr += "'\n"
			tempcount = 0
			continue
		}
		if text[i] == '\'' && (i != 0 && ((text[i-1] >= 'a' && text[i-1] <= 'z') || (text[i-1] >= 'A' && text[i-1] <= 'Z'))) && (i != len(text)-1 && ((text[i+1] >= 'a' && text[i+1] <= 'z') || (text[i+1] >= 'A' && text[i+1] <= 'Z'))) {
			tempstr += string(text[i])
			continue
		}
		if text[i] == '\'' && tempcount == 0 {
			if i != 0 && text[i-1] != ' ' && text[i-1] != '\t' && text[i-1] != '\n' && text[i-1] != '\r' {
				tempstr += " "
			}
			tempcount = 1
			tempstr += string(text[i])
			continue
		}
		if (text[i] == ' ' || text[i] == '\t' || text[i] == '\n' || text[i] == '\r') && tempcount == 1 {
			tempcount = 2
			continue
		} else if !(text[i] == ' ' || text[i] == '\t' || text[i] == '\n' || text[i] == '\r') && tempcount == 1 {
			tempstr += string(text[i])
			tempcount = 2
			continue
		}
		if i+1 != len(text) && text[i+1] == '\'' && tempcount == 2 && (text[i] == ' ' || text[i] == '\t' || text[i] == '\n' || text[i] == '\r') {
			tempcount = 0
			tempstr += "'"
			i++
			if text[i+1] != ' ' && text[i+1] != '\t' && text[i+1] != '\n' && text[i+1] != '\r' {
				tempstr += " "
			}
			continue
		}
		tempstr += string(text[i])
	}
	return tempstr
}

func PonctuationFixer(text []rune) string {
	tempstr := ""
	for i := 0; i < len(text); i++ {
		if i+1 != len(text) && (text[i+1] == '.' || text[i+1] == ',' || text[i+1] == ';' || text[i+1] == ':' || text[i+1] == '!' || text[i+1] == '?') && (text[i] == ' ' || text[i] == '\t' || text[i] == '\r') {
			continue
		}
		if i+1 != len(text) && (text[i] == '.' || text[i] == ',' || text[i] == ';' || text[i] == ':' || text[i] == '!' || text[i] == '?') && !(text[i+1] == ' ' || text[i+1] == '\t' || text[i+1] == '\n' || text[i+1] == '\r') && !(text[i+1] == '.' || text[i+1] == ',' || text[i+1] == ';' || text[i+1] == ':' || text[i+1] == '!' || text[i+1] == '?') {
			tempstr += string(text[i])
			tempstr += " "
			continue
		}
		tempstr += string(text[i])
	}
	return tempstr
}

// our main function

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not Enough Args")
		return
	}
	if len(os.Args) == 2 {
		fmt.Println("MIssing the Result File Name")
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("Too Many Args")
		return
	}
	for i := len(os.Args[1]) - 1; i >= 0; i-- {
		if os.Args[1][i] == '.' {
			if i+1 < len(os.Args[1]) && os.Args[1][i+1:] != "txt" {
				fmt.Println("Input File must be a .txt file")
				return
			}
			break
		}
		if i == 0 {
			fmt.Println("Input File must be a .txt file")
			return
		}
	}
	for i := len(os.Args[2]) - 1; i >= 0; i-- {
		if os.Args[2][i] == '.' {
			if i+1 < len(os.Args[2]) && os.Args[2][i+1:] != "txt" {
				fmt.Println("Resulted File must be a .txt file")
				return
			}
			break
		}
		if i == 0 {
			fmt.Println("Resulted File must be a .txt file")
			return
		}
	}
	file, ferr := os.Open(os.Args[1])
	if ferr != nil {
		fmt.Println("Error opening file:", ferr)
		return
	}
	defer file.Close()
	temptext, terr := io.ReadAll(file)
	if terr != nil {
		fmt.Println("Error reading file:", terr)
		return
	}
	tempstr := string(temptext)
	text := []rune(tempstr)
	tempstr = ""
	tempcount := 0
	// here we are going to remove any extra newline or pace or tabulation
	for i := 0; i < len(text); i++ {
		if text[i] != ' ' && text[i] != '\t' && text[i] != '\n' && text[i] != '\r' {
			tempstr += string(text[i])
			tempcount = 0
			continue
		}
		if tempcount == 0 {
			tempstr += string(text[i])
			tempcount++
		}
		if i == len(text)-1 && text[i] != '\n' && text[i+1] != '\n' {
			tempstr += string(text[i])
			tempstr += "\n"
		}
	}
	text = []rune(tempstr)
	tempstr = ""
	tempcount = 0
	// here we are going to replace if necessary
	txt := reloaded.StringPlitter(text)
	final := []string{}
	vld := 0
	for i := 0; i < len(txt); i++ {
		count := 0
		if reloaded.IsValid(txt[i]) {
			vld++
			if txt[i] == "" {
				continue
			}
			if txt[i][len(txt[i])-1] == '\'' || txt[i][len(txt[i])-2] == '\'' {
				final = append(final, "'")
			}
			// here we would do the transformation
			if reloaded.IsHex(txt[i]) {
				final[i-vld] = reloaded.Hex(final[i-vld])
				continue
			}
			if reloaded.IsBin(txt[i]) {
				final[i-vld] = reloaded.Bin(final[i-vld])
				continue
			}
			if reloaded.IsUp(txt[i]) {
				count = reloaded.Detect(txt[i])
				if count == 0 {
					continue
				}
				if count == 1 {
					final[i-vld] = reloaded.Up(final[i-vld])
					continue
				}
				for j := 0; j < count; j++ {
					final[i-j-vld] = reloaded.Up(final[i-j-vld])
				}
				continue
			}
			if reloaded.IsLow(txt[i]) {
				count = reloaded.Detect(txt[i])
				if count == 0 {
					continue
				}
				if count == 1 {
					final[i-vld] = reloaded.Low(final[i-vld])
					continue
				}
				for j := 0; j < count; j++ {
					final[i-j-vld] = reloaded.Low(final[i-j-vld])
				}
				continue
			}
			if reloaded.IsCap(txt[i]) {
				count = reloaded.Detect(txt[i])
				if count == 0 {
					continue
				}
				if count == 1 {
					final[i-vld] = reloaded.Cap(final[i-vld])
					continue
				}
				for j := 0; j < count; j++ {
					final[i-j-vld] = reloaded.Cap(final[i-j-vld])
				}
				continue
			}
		}
		final = append(final, txt[i])
	}
	FinalResult := ""
	for i := 0; i < len(final); i++ {
		FinalResult += final[i]
	}
	text = []rune(FinalResult)
	FinalResult = ""
	// here we are going to fix the a and an problem
	for i := 0; i < len(text); i++ {
		if i+2 < len(text) && (text[i] == 'a' || text[i] == 'A') && text[i+1] == ' ' {
			if text[i+2] == 'a' || text[i+2] == 'e' || text[i+2] == 'i' || text[i+2] == 'o' || text[i+2] == 'u' || text[i+2] == 'A' || text[i+2] == 'E' || text[i+2] == 'I' || text[i+2] == 'O' || text[i+2] == 'U' {
				FinalResult += "an "
				i++
				continue
			}
		}
		FinalResult += string(text[i])
	}
	// here we are going to process the quotes
	FinalResult = QuotesFixer([]rune(FinalResult))
	// here we are going to process the punctuations
	FinalResult = PonctuationFixer([]rune(FinalResult))
	err := os.WriteFile(os.Args[2], []byte(FinalResult), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Print("File Processed Successfully\n")
}

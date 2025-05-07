package BA

import (
	"os"
	"strings"
)

func LoadFont(style string) map[rune][]string {
	var StyleLocation string
	var lines []string
	if style == "standard" {
		StyleLocation = "../internal/templates/standard.txt"
	} else if style == "shadow" {
		StyleLocation = "../internal/templates/shadow.txt"
	} else if style == "thinkertoy" {
		StyleLocation = "../internal/templates/thinkertoy.txt"
	} else {
		println("\033[31;1m ERROR: provided template is not correct!\033[0m")
		os.Exit(0)
	}

	data, err := os.ReadFile(StyleLocation)
	if err != nil {
		println("\033[31;1m ERROR: reading the file!\033[0m")
		os.Exit(0)
	}
	if style == "thinkertoy" {
		lines = strings.Split(string(data), "\r\n")
	} else {
		lines = strings.Split(string(data), "\n")
	}

	fontMap := make(map[rune][]string)
	startChar := 32 // space character
	for i := 1; i+7 < len(lines); i += 9 {
		char := rune(startChar)
		fontMap[char] = lines[i : i+8]
		startChar++
	}
	return fontMap
}

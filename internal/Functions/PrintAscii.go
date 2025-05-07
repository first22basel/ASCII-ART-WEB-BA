package BA

import (
	"errors"
	"strings"
)

func PrintAscii(input string, fontMap map[rune][]string) (string, error) {
	var result strings.Builder

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range line {
				if ch == '\r' {
					continue // Ignore carriage returns
				}
				if art, ok := fontMap[ch]; ok {
					result.WriteString(art[row])
				} else {
					return "", errors.New("unsupported character: '" + string(ch) + "'")
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

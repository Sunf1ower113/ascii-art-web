package utils

import (
	"fmt"
	"regexp"
)

func ValidString(str string) ([]string, error) {
	temp := ""
	var result []string
	var err error
	if len(str) == 0 {
		return nil, fmt.Errorf("input error: %w", ArtError.InputError)
	}
	r2 := regexp.MustCompile(string(byte(13)))
	str = r2.ReplaceAllString(str, "")
	r1 := regexp.MustCompile(`\n`)
	str = r1.ReplaceAllString(str, "\\n")
	r := regexp.MustCompile(`\\n`)
	newLines := r.Split(str, len(str))
	for _, str := range newLines {
		for i := 0; i < len(str); i++ {
			if str[i] == '\\' {
				if i < len(str)-1 {
					if str[i+1] == '!' && i != 0 {
						temp += "!"
						i++
						continue
					} else if str[i+1] == 'a' ||
						str[i+1] == 'b' ||
						str[i+1] == 't' ||
						str[i+1] == 'v' ||
						str[i+1] == 'f' ||
						str[i+1] == 'r' ||
						str[i+1] == '0' {
						i++
						continue
					}
				}
			}
			if str[i] >= 32 && str[i] <= 126 {
				temp += string(str[i])
			} else {
				err = fmt.Errorf("input error: %w", ArtError.InputError)
				return nil, err
			}
		}
		result = append(result, temp)
		temp = ""
	}
	return result, err
}

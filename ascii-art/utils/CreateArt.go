package utils

import (
	"fmt"
	"strings"
)

const (
	AlphLen      = 95
	LetterHeight = 8
)

func CreateAlph(symbols []byte) ([][]string, error) {
	alph := make([][]string, 0, AlphLen*64)
	temp := ""
	for _, ch := range symbols {
		if ch >= ' ' && ch <= '~' || ch == '\n' {
			temp += string(ch)
		}
	}
	symbols = []byte(temp)
	tempAlph := strings.Split(string(symbols), "\n\n")
	tempArr := make([]string, 0, 64)

	if len(tempAlph) != AlphLen {
		return [][]string{}, fmt.Errorf("validateAlphLen: %w", ArtError.BannerError1)
	}
	for _, ch := range tempAlph {
		tempArr = strings.Split(ch, "\n")
		count := 0
		for i, ch := range tempArr {
			if i == 0 {
				count = len(ch)
			} else if count != len(ch) {
				return [][]string{}, fmt.Errorf("validateAlphLen: %w", ArtError.BannerError2)
			}
		}
		if len(tempArr) != LetterHeight {
			return [][]string{}, fmt.Errorf("validateAlphLen: %w", ArtError.BannerError3)
		}
		alph = append(alph, tempArr)
	}
	return alph, nil
}

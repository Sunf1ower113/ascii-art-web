package asciiart

import (
    "ascii-art-web/ascii-art/utils"
    "io/ioutil"
)

type Art struct {
    Text []byte
    Banner string `default:"standard"`
    AsciiArt string
}
func newArt(text, banner string, ascii [][]string, str []string, i int) *Art {
    p := new(Art)
    p.Text = []byte(text)
    p.Banner = banner
    p.AsciiArt = createArt(ascii, str, i, "")
    return p
}

func createArt(ascii [][]string, str []string, i int, art string) string {
    for j := 0; j < utils.LetterHeight; j++ {
        if len(str[i]) == 0 {
            art += string('\n')
            j = 0
            break
        }
        for _, ch := range str[i] {
            art += ascii[ch-32][j]
        }
        art += string('\n')
    }
    if i < len(str)-1 {
        art = createArt(ascii, str, i+1, art)
    }
    return art
}
func AsciiArt(text, banner string) (*Art, error) {
	str, err := utils.ValidString(text)
	if err != nil {
		return nil, err
	}
	if len(str) == 1 {
		if str[0] == "" {
			return nil, nil
		}
	}
	symbols, err := ioutil.ReadFile(banner)
	if err != nil {
		return nil, err
	}

	alph, err := utils.CreateAlph(symbols)
	if err != nil {
		return nil, err
	}
    return newArt(text, banner, alph, str, 0), nil
}

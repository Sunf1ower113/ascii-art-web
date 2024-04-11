package utils

import (
	"fmt"
	"io/ioutil"
)

func PrintArt(art string) {
	fmt.Print(art)
}

func WriteArt(filename, art string) error {
	err := ioutil.WriteFile(filename, []byte(art), 0644)
	if err != nil {
		return err
	}
	return nil
}

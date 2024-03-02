package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encrypt(key int, plainText string) (result string) {

	hasLetter := hashLetterFn(key, originalLetter)
	var hashString = ""
	findOne := func(r rune) rune {
		pos := strings.Index(originalLetter, string([]rune{r}))
     
		 if pos!= -1 {
			
		 }

return r
	}

	strings.Map(findOne, plainText)

}

func hashLetterFn(key int, letter string) (result string) {

	runes := []rune(letter)

	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOverLetters := string(runes[0 : len(letter)-key])

	return fmt.Sprintf(`%s%s`, lastLetterKey, leftOverLetters)

}

func decrypt(key int, encryptedText string) (result string) {

}

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain Text : ", plainText)
	encrypted := encrypt(5, plainText)

	fmt.Printf("Encrypted Text", encrypted)

	decrypted := decrypt(5, encrypted)
	fmt.Printf("Decrypted : ", decrypted)
}

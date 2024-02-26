package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encrypt(key int, plainText string) (result string) {

	hasLetter := hashLetterFn(key, originalLetter)
	var hashString = ""

	strings.Map(findOne,plainText)

}

func hashLetterFn(key int, letter string) (result string) {}

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

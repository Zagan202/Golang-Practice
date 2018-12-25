/***************************************************************************************
*
	 Title: Ceasar Cipher
*    Original Author: Golang Programs
*	 Altered Ver Author: Lawrence Lawson
*    Availability: http://www.golangprograms.com/cryptography/what-is-cryptography.html
*
***************************************************************************************/
package main

import (
	"fmt"
	"unicode"
)

// Cipher encrypts and decrypts a string.
type CeaserCipher interface {
	CeaserEncryption(string) string
	CeaserDecryption(string) string
}

// Cipher holds the key used to encrypts and decrypts messages.
type Ccipher []int

// cipherAlgorithm encodes a letter based on some function.
func (c Ccipher) CeasercipherAlgorithm(letters string, shift func(int, int) int) string {
	shiftedText := ""
	for _, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue
		}
		shiftDist := c[len(shiftedText)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		switch {
		case s < 'a':
			s += 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}

// Encryption encrypts a message.
func (c *Ccipher) CeaserEncryption(plainText string) string {
	return c.CeasercipherAlgorithm(plainText, func(a, b int) int { return a + b })
}

// Decryption decrypts a message.
func (c *Ccipher) CeaserDecryption(cipherText string) string {
	return c.CeasercipherAlgorithm(cipherText, func(a, b int) int { return a - b })
}

// NewCaesar creates a new Caesar shift cipher.
func NewCaesar(key int) CeaserCipher {
	return NewShift(key)
}

// NewShift creates a new Shift cipher.
func NewShift(shift int) CeaserCipher {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	c := Ccipher([]int{shift})
	return &c
}

func runCeaser() {
	c := NewCaesar(1)
	fmt.Println("Encrypt Key(01) abcd =>", c.CeaserEncryption("abcd"))
	fmt.Println("Decrypt Key(01) bcde =>", c.CeaserDecryption("bcde"))
	fmt.Println()

	c = NewCaesar(10)
	fmt.Println("Encrypt Key(10) abcd =>", c.CeaserEncryption("abcd"))
	fmt.Println("Decrypt Key(10) klmn =>", c.CeaserDecryption("klmn"))
	fmt.Println()

	c = NewCaesar(15)
	fmt.Println("Encrypt Key(15) abcd =>", c.CeaserEncryption("abcd"))
	fmt.Println("Decrypt Key(15) pqrs =>", c.CeaserDecryption("pqrs"))
}

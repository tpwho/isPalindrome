package isPalindrome

import (
	"regexp"
	"strings"
)

func removeNonAlpha(s string) string {
	reg, err := regexp.Compile("[^a-z0-9]+")
	if err != nil {
		panic(err)
	}
	result := reg.ReplaceAllString(s, "")
	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//Check will test to see if a given string is a palindrome or not. Returns true/false
func Check(sentence string) bool {
	//We don't care about case, so make it all lower
	sentence = strings.ToLower(sentence)
	strippedSentence := removeNonAlpha(sentence)
	reversedSentence := reverseString(strippedSentence)
	if strippedSentence == reversedSentence {
		return true
	}
	return false
}

func isPalindrome() {}

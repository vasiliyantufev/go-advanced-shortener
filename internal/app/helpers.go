package app

import (
	"math/rand"
	"regexp"
	"strings"
)

//const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func shorting(inputValue string) string {

	b := make([]byte, 5)
	validValue := getString(inputValue)

	for i := range b {

		index := rand.Intn(len(validValue))
		b[i] = validValue[index]
	}

	check := ExistElement(string(b))
	if check {
		shorting(inputValue)
	}
	return string(b)
}

func getString(input string) (validStr string) {
	re := regexp.MustCompile("[0-9A-Za-z]+")
	validValue := re.FindAllString(input, -1)
	validStr = strings.ReplaceAll(strings.Join(validValue, " "), " ", "")
	return
}

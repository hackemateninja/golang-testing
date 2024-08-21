package iteration

import "strings"

var builder strings.Builder

const repeatCount = 5

func Repeat(value string) string {
	for i := 0; i < 5; i++ {
		builder.WriteString(value)
	}
	return builder.String()
}

func RepeatTraditional(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

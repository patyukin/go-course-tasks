package main

import (
	"fmt"

	"patyukin/go-course-tasks/hw-3/message_formatter/formatter"
)

func main() {
	text := "Hello, world!"

	plainText := formatter.PlainTextFormatter{}
	bold := formatter.BoldFormatter{}
	code := formatter.CodeFormatter{}
	italic := formatter.ItalicFormatter{}

	chainFormatter := formatter.ChainFormatter{}
	chainFormatter.AddFormatter(plainText)
	chainFormatter.AddFormatter(bold)
	chainFormatter.AddFormatter(code)
	chainFormatter.AddFormatter(italic)

	formattedText := chainFormatter.Format(text)
	fmt.Println("Formatted Text:", formattedText)
}

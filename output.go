package ejikit

import (
	"fmt"
	"log"
)

func OnLineEmojiOutPut() {
	output := ""
	emojiMap, err := GenerateJson(false)
	if err != nil {
		log.Fatalln(err)
	}

	for key, value := range emojiMap {
		output += key + ":\t" + value + "\n"
	}

	fmt.Print(output)
}

func EmojiOutPut() {
	output := ""
	emojiMap, err := GenerateJson(false)
	if err != nil {
		log.Fatalln(err)
	}

	index := 0
	space := " "
	oneLine := ""
	for _, value := range emojiMap {
		if index <= 15 {
			oneLine += value + space
			index += 1
		} else {
			output += "\t" + oneLine + "\t" + "\n"
			oneLine = ""
			index = 0
		}
	}
	fmt.Printf("%s", output)
}

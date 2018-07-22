package ejikit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const gemojiDBJsonURL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"

func GenerateJson(permitEscape bool) (map[string]string, error) {
	res, err := http.Get(gemojiDBJsonURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	emojiFile, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var gs []GemojiEmoji
	if err := json.Unmarshal(emojiFile, &gs); err != nil {
		return nil, err
	}

	emojiCodeMap := make(map[string]string)
	for _, gemoji := range gs {
		for _, a := range gemoji.Aliases {
			if permitEscape {
				emojiCodeMap[a] = fmt.Sprintf("%+q", gemoji.Emoji)
			} else {
				emojiCodeMap[a] = gemoji.Emoji
			}
		}
	}

	return emojiCodeMap, nil
}

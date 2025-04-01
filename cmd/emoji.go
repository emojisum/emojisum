package cmd

import (
	esum "github.com/emojisum/emojisum/emoji"
)

// EmojiFromBytes parses the bytes buffer for colon notation and returns the
// corresponding emoji
func EmojiFromBytes(buf []byte) string {
	var ret string
	for _, b := range buf {
		for _, e := range esum.Map(b) {
			// use the first colon notation word and continue
			if esum.IsColonNotation(e) {
				ret = ret + e
				break
			}
		}
	}
	return ret
}

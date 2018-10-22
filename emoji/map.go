package emoji

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// Map returns the emoji at the provided position.
// This list is from 0-255
func Map(b byte) Words {
	return mapGen.EmojiWords[int(b)]
}

// Version returns the version of the emojisum document currently compiled
// against
func Version() string {
	return mapGen.Version
}

var mapGen VersionedMap

// VersionedMap is the structure used for the `emojimap.json` document
type VersionedMap struct {
	Description string `json:"description"`
	Version     string `json:"version"`
	// these are an ordered list, referened by a byte (each byte of a checksum digest)
	EmojiWords []Words `json:"emojiwords"`
}

// Words are a set of options to represent an emoji.
// Possible options could be the ":colon_notation:", a "U+26CF" style
// codepoint, or the unicode value itself.
type Words []string

// IsColonNotation checks for whether a word is the :colon_notation: of emoji
func IsColonNotation(word string) bool {
	return strings.HasPrefix(word, ":") && strings.HasSuffix(word, ":")
}

// IsCodepoint checks for whether a word is the "U+1234" codepoint style of emoji. Codepoints can sometimes be a combo, like flags
func IsCodepoint(word string) bool {
	return strings.HasPrefix(strings.ToUpper(word), "U+")
}

var unicodeURL = `http://www.unicode.org/emoji/charts/full-emoji-list.html`

// UnicodeLinkURL returns a link to unicode.org list for CodePoint, or just the
// full list if not a codepoint
func UnicodeLinkURL(word string) string {
	if !IsCodepoint(word) {
		return unicodeURL
	}

	return fmt.Sprintf("%s#%s", unicodeURL, strings.Join(strings.Split(strings.TrimPrefix(strings.ToLower(word), "u+"), "u+"), "_"))
}

// CodepointToUnicode takes a "U+26CF" style word and returns the `\U00026CF` formated unicode string
func CodepointToUnicode(word string) string {
	if !IsCodepoint(word) {
		return word
	}

	var ret string

	for _, chunk := range strings.Split(strings.TrimPrefix(strings.ToUpper(word), "U+"), "U+") {
		c, err := strconv.ParseInt(chunk, 16, 64)
		if err != nil {
			return ret
		}
		ret = fmt.Sprintf("%s%c", ret, c)
	}
	return ret
}

// FromHexString parses string s as two character byte of hexadecimal into
// Unicode Codepoint
func FromHexString(s string) (string, error) {
	d, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	var ret string
	for _, b := range d {
		for _, e := range Map(b) {
			// use the first colon notation word and continue
			if IsCodepoint(e) {
				ret = ret + CodepointToUnicode(e)
				break
			}
		}
	}
	return ret, nil
}

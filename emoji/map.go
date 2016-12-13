//go:generate go run map_json.go -in ./emojimap.json -out ./map_gen.go

package emoji

// Map returns the emoji at the provided position.
// This list is from 0-255
func Map(b byte) string {
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
	EmojiWords []string `json:"emojiwords"`
}

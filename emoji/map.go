//go:generate go run map_json.go

package emoji

// Map returns the emoji at the provided position.
// This list is from 0-255
func Map(b byte) string {
	return sumList[int(b)]
}

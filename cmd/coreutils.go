package cmd

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kyokomi/emoji"
)

/*
ParseCoreUtils expects some input like:
```shell
$ sha256sum ./tmp.efLuko
f18bd8b680e834ab8097a66deb0255821195d9624e39da6b65903ff6a09a01bb  ./tmp.efLuko
```
*/
func ParseCoreUtils(line string) (filename string, sum []byte, err error) {
	chunks := strings.SplitN(strings.TrimRight(line, "\n"), "  ", 2)
	if len(chunks) != 2 {
		return "", nil, ErrNotCoreUtilsLine
	}
	sum, err = hex.DecodeString(chunks[0])
	if err != nil {
		return "", nil, err
	}
	return chunks[1], sum, nil
}

// ErrNotCoreUtilsLine when the line to parse is not formated like a coreutils checksum line
var ErrNotCoreUtilsLine = errors.New("not a coreutils checksum line")

// PrintCoreUtils reads in the content like from `sha25sum ...` and returns
// in likeness but with emojisum instead.
// TODO(vb) return a buffer that the caller can choose to print out themselves
func PrintCoreUtils(rdr io.Reader) error {
	buf := bufio.NewReader(rdr)
	for {
		line, err := buf.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		name, sum, err := ParseCoreUtils(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %q\n", err, line)
			continue
		}
		str := EmojiFromBytes(sum)
		fmt.Printf("%x  %s\n", sum, name)
		fmt.Printf("%s  %s\n", str, name)
		emoji.Print(str)
		fmt.Printf("  %s\n", name)
	}
	return nil
}

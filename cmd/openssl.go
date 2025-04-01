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
ParseOpenSSL expects some input like:
```shell
$> openssl sha256 tmp.efLuko
SHA256(tmp.efLuko)= f18bd8b680e834ab8097a66deb0255821195d9624e39da6b65903ff6a09a01bb
```
*/
func ParseOpenSSL(line string) (hash, filename string, sum []byte, err error) {
	if !strings.Contains(line, "(") {
		return "", "", nil, ErrNotOpenSSLLine
	}
	chunks := strings.SplitN(strings.TrimRight(line, "\n"), ")= ", 2)
	if len(chunks) != 2 {
		return "", "", nil, ErrNotOpenSSLLine
	}
	chunksprime := strings.SplitN(chunks[0], "(", 2)
	if len(chunks) != 2 {
		return "", "", nil, ErrNotOpenSSLLine
	}
	sum, err = hex.DecodeString(chunks[1])
	if err != nil {
		return "", "", nil, err
	}
	return chunksprime[0], chunksprime[1], sum, nil
}

// ErrNotOpenSSLLine when the line to parse is not formated like an OpenSSL checksum line
var ErrNotOpenSSLLine = errors.New("not an openssl checksum line")

// PrintOpenSSL reads in the content like from `openssl sha25 ...` and returns
// in likeness but with emojisum instead.
// TODO(vb) return a buffer that the caller can choose to print out themselves
func PrintOpenSSL(rdr io.Reader) error {
	buf := bufio.NewReader(rdr)
	for {
		line, err := buf.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		hash, name, sum, err := ParseOpenSSL(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %q\n", err, line)
			continue
		}
		str := EmojiFromBytes(sum)
		fmt.Printf("%s(%s)= %x\n", hash, name, sum)
		fmt.Printf("%s(%s)= %s\n", hash, name, str)
		fmt.Printf("%s(%s)= ", hash, name)
		emoji.Println(str)
	}
	return nil
}

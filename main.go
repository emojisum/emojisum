package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kyokomi/emoji"
	esum "github.com/vbatts/emojisum/emoji"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var (
	flParseOpenSSL   = flag.Bool("parse-openssl", false, "parse the output of OpenSSL style checksums on stdin (`openssl sha256 ./foo | emojisum -parse-openssl`)")
	flParseCoreUtils = flag.Bool("parse-coreutils", false, "parse the output of CoreUtils style checksums on stdin (`sha256sum ./foo | emojisum -parse-coreutils`)")
)

func run() error {
	flag.Parse()

	if *flParseOpenSSL {
		buf := bufio.NewReader(os.Stdin)
		for {
			line, err := buf.ReadString('\n')
			if err != nil && err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}
			hash, name, sum, err := parseOpenSSL(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %q\n", err, line)
				continue
			}
			str := emojiFromBytes(sum)
			fmt.Printf("%s(%s)= %x\n", hash, name, sum)
			fmt.Printf("%s(%s)= %s\n", hash, name, str)
			fmt.Printf("%s(%s)= ", hash, name)
			emoji.Println(str)
		}
		// never gets here because of the return on EOF or err
	}
	if *flParseCoreUtils {
		buf := bufio.NewReader(os.Stdin)
		for {
			line, err := buf.ReadString('\n')
			if err != nil && err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}
			name, sum, err := parseCoreUtils(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %q\n", err, line)
				continue
			}
			str := emojiFromBytes(sum)
			fmt.Printf("%x  %s\n", sum, name)
			fmt.Printf("%s  %s\n", str, name)
			emoji.Print(str)
			fmt.Printf("  %s\n", name)
		}
		// never gets here because of the return on EOF or err
	}

	// Otherwise do the checksum ourselves

	if flag.NArg() == 0 {
		sum, err := Sum(os.Stdin)
		if err != nil {
			return err
		}
		str := emojiFromBytes(sum)
		fmt.Printf("SHA1(-)= %x\n", sum)
		fmt.Printf("SHA1(-)= %s\n", str)
		fmt.Printf("SHA1(-)= ")
		emoji.Println(str)
		return nil
	}

	for _, arg := range flag.Args() {
		fh, err := os.Open(arg)
		if err != nil {
			return err
		}
		defer fh.Close()

		sum, err := Sum(fh)
		if err != nil {
			return err
		}
		str := emojiFromBytes(sum)
		fmt.Printf("SHA1(%s)= %x\n", arg, sum)
		fmt.Printf("SHA1(%s)= %s\n", arg, str)
		fmt.Printf("SHA1(%s)= ", arg)
		emoji.Println(str)
	}
	return nil
}

/*
openssl sum:
```
$> openssl sha256 tmp.efLuko
SHA256(tmp.efLuko)= f18bd8b680e834ab8097a66deb0255821195d9624e39da6b65903ff6a09a01bb
```
*/
func parseOpenSSL(line string) (hash, filename string, sum []byte, err error) {
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

/*
coreutils output:
```
$ sha256sum ./tmp.efLuko
f18bd8b680e834ab8097a66deb0255821195d9624e39da6b65903ff6a09a01bb  ./tmp.efLuko
```
*/
func parseCoreUtils(line string) (filename string, sum []byte, err error) {
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

// Sum is a basic wrapper around crypto/sha1
func Sum(r io.Reader) ([]byte, error) {
	h := sha1.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	sum := h.Sum(nil)
	return sum[:], nil
}

func emojiFromBytes(buf []byte) string {
	var ret string
	for _, b := range buf {
		ret = ret + esum.Map(b)[0]
	}
	return ret
}

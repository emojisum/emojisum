package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emojisum/emojisum/cmd"
	"github.com/kyokomi/emoji"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var (
	flParseOpenSSL   = flag.Bool("pb", false, "parse the output of BSD/OpenSSL style checksums on stdin (`openssl sha256 ./foo | emojisum -pb`)")
	flParseCoreUtils = flag.Bool("pg", false, "parse the output of GNU/CoreUtils style checksums on stdin (`sha256sum ./foo | emojisum -pg`)")
)

func run() error {
	flag.Parse()

	if *flParseOpenSSL {
		return cmd.PrintOpenSSL(os.Stdin)
	}
	if *flParseCoreUtils {
		return cmd.PrintCoreUtils(os.Stdin)
	}

	// Otherwise do the checksum ourselves

	if flag.NArg() == 0 {
		sum, err := cmd.Sum(os.Stdin)
		if err != nil {
			return err
		}
		str := cmd.EmojiFromBytes(sum)
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

		sum, err := cmd.Sum(fh)
		if err != nil {
			return err
		}
		str := cmd.EmojiFromBytes(sum)
		fmt.Printf("SHA1(%s)= %x\n", arg, sum)
		fmt.Printf("SHA1(%s)= %s\n", arg, str)
		fmt.Printf("SHA1(%s)= ", arg)
		emoji.Println(str)
	}
	return nil
}

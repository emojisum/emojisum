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
		return nil
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
		return nil
	}

	// Otherwise do the checksum ourselves

	if flag.NArg() == 0 {
		sum, err := Sum(os.Stdin)
		if err != nil {
			return err
		}
		str := emojiFromBytes(sum)
		fmt.Printf("SHA1(-)=\t%x\n", sum)
		fmt.Printf("SHA1(-)=\t%s\n", str)
		fmt.Printf("SHA1(-)=\t")
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
		fmt.Printf("SHA1(%s)=\t%x\n", arg, sum)
		fmt.Printf("SHA1(%s)=\t%s\n", arg, str)
		fmt.Printf("SHA1(%s)=\t", arg)
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

var ErrNotCoreUtilsLine = errors.New("not a coreutils checksum line")

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
		ret = ret + emojiFromRune(b)
	}
	return ret
}

func emojiFromRune(b byte) string {
	//fmt.Printf("%#v\n", b)
	return emojiMap[int(b)]
}

// these are an ordered list, referened by a byte (each byte of a checksum digest)
var emojiMap = []string{
	":+1:",
	":8ball:",
	":airplane:",
	":alien:",
	":anchor:",
	":angel:",
	":angry:",
	":ant:",
	":apple:",
	":art:",
	":baby:",
	":baby_bottle:",
	":back:",
	":bamboo:",
	":banana:",
	":barber:",
	":bathtub:",
	":beer:",
	":bell:",
	":bicyclist:",
	":bird:",
	":birthday:",
	":blossom:",
	":blue_car:",
	":boar:",
	":bomb:",
	":boom:",
	":bow:",
	":boy:",
	":broken_heart:",
	":bulb:",
	":bus:",
	":cactus:",
	":calendar:",
	":camera:",
	":candy:",
	":cat:",
	":cherries:",
	":children_crossing:",
	":chocolate_bar:",
	":clap:",
	":cloud:",
	":clubs:",
	":cn:",
	":coffee:",
	":construction:",
	":cookie:",
	":copyright:",
	":corn:",
	":cow:",
	":crescent_moon:",
	":crown:",
	":cry:",
	":crystal_ball:",
	":curly_loop:",
	":dancers:",
	":dash:",
	":de:",
	":diamonds:",
	":dog:",
	":doughnut:",
	":dragon:",
	":dvd:",
	":ear:",
	":eggplant:",
	":elephant:",
	":end:",
	":envelope:",
	":es:",
	":eyes:",
	":facepunch:",
	":family:",
	":ferris_wheel:",
	":finnadie:",
	":fire:",
	":fireworks:",
	":floppy_disk:",
	":football:",
	":fork_and_knife:",
	":four_leaf_clover:",
	":fr:",
	":fries:",
	":frog:",
	":fu:",
	":full_moon:",
	":game_die:",
	":gb:",
	":gem:",
	":girl:",
	":goat:",
	":grimacing:",
	":grin:",
	":guardsman:",
	":guitar:",
	":gun:",
	":hamburger:",
	":hammer:",
	":hamster:",
	":hear_no_evil:",
	":heart:",
	":heart_eyes_cat:",
	":hearts:",
	":heavy_check_mark:",
	":moyai:",
	":izakaya_lantern:",
	":helicopter:",
	":hocho:",
	":honeybee:",
	":horse:",
	":horse_racing:",
	":hourglass:",
	":house:",
	":hurtrealbad:",
	":icecream:",
	":imp:",
	":it:",
	":jack_o_lantern:",
	":japanese_goblin:",
	":jp:",
	":key:",
	":kiss:",
	":kissing_cat:",
	":koala:",
	":kr:",
	":lemon:",
	":lipstick:",
	":lock:",
	":lollipop:",
	":man:",
	":maple_leaf:",
	":mask:",
	":metal:",
	":microscope:",
	":moneybag:",
	":monkey:",
	":mount_fuji:",
	":muscle:",
	":mushroom:",
	":musical_keyboard:",
	":musical_score:",
	":nail_care:",
	":new_moon:",
	":no_entry:",
	":nose:",
	":notes:",
	":nut_and_bolt:",
	":o:",
	":ocean:",
	":ok_hand:",
	":on:",
	":package:",
	":palm_tree:",
	":panda_face:",
	":paperclip:",
	":partly_sunny:",
	":passport_control:",
	":paw_prints:",
	":peach:",
	":penguin:",
	":phone:",
	":pig:",
	":pill:",
	":pineapple:",
	":pizza:",
	":point_left:",
	":point_right:",
	":poop:",
	":poultry_leg:",
	":pray:",
	":princess:",
	":purse:",
	":pushpin:",
	":rabbit:",
	":rainbow:",
	":raised_hand:",
	":recycle:",
	":red_car:",
	":registered:",
	":ribbon:",
	":rice:",
	":rocket:",
	":roller_coaster:",
	":rooster:",
	":ru:",
	":sailboat:",
	":santa:",
	":satellite:",
	":satisfied:",
	":saxophone:",
	":scissors:",
	":see_no_evil:",
	":sheep:",
	":shell:",
	":shoe:",
	":ski:",
	":skull:",
	":sleepy:",
	":smile:",
	":smiley_cat:",
	":smirk:",
	":smoking:",
	":snail:",
	":snake:",
	":snowflake:",
	":soccer:",
	":soon:",
	":space_invader:",
	":spades:",
	":speak_no_evil:",
	":star:",
	":stars:",
	":statue_of_liberty:",
	":steam_locomotive:",
	":sunflower:",
	":sunglasses:",
	":sunny:",
	":sunrise:",
	":surfer:",
	":swimmer:",
	":syringe:",
	":tada:",
	":tangerine:",
	":taxi:",
	":tennis:",
	":tent:",
	":thought_balloon:",
	":tm:",
	":toilet:",
	":tongue:",
	":tophat:",
	":tractor:",
	":trolleybus:",
	":trollface:",
	":trophy:",
	":trumpet:",
	":turtle:",
	":two_men_holding_hands:",
	":two_women_holding_hands:",
	":uk:",
	":umbrella:",
	":unlock:",
	":us:",
	":v:",
	":vhs:",
	":violin:",
	":warning:",
	":watermelon:",
	":wave:",
	":wavy_dash:",
	":wc:",
	":wheelchair:",
	":woman:",
	":x:",
	":yum:",
	":zap:",
	":zzz:",
}

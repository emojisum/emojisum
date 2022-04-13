# Emojisum
[![Build Status](https://travis-ci.org/emojisum/emojisum.svg?branch=master)](https://travis-ci.org/emojisum/emojisum)

:pray: :paperclip: An easier way to compare hashes /fingerprints, when dealing human weak link :link: :tada:

A curated list of 256 emojis that are not entirely similar.
Using http://www.webpagefx.com/tools/emoji-cheat-sheet/ and http://www.unicode.org/emoji/charts/full-emoji-list.html to compare them.
With 256 as that is 8bit/1byte, and the hexadecimal output that is 2 hex characters.

So 2 hex positions are 1 emoji!
Would you rather compare 60 hexadecimal characters or only 30 emoji?! :smiley_cat:

## Specification

For information on the draft for broader practice, see [the draft](./emoji/README.md).
Perma-URL:
- [Draft](https://github.com/emojisum/emojisum/blob/master/emoji/README.md)
- [JSON doc](https://raw.githubusercontent.com/emojisum/emojisum/master/emoji/emojimap.json)

## Tool

### build

```bash
go get github.com/emojisum/emojisum
```

### usage

This uses [`github.com/kyokomi/emoji`](https://github.com/kyokomi/emoji) to print to the console, but also gives the string ouptut for easy pasting to github/markdown.

```bash
$> emojisum main.go 
SHA1(main.go)=  14b09535217ca8f5f47f4665e2266e686f0728b4
SHA1(main.go)=  :bird::red_car::on::crystal_ball::calendar::lemon::pray::warning::violin::lollipop::facepunch::hearts::tm::children_crossing::hourglass::heavy_plus_sign::house::ant::clap::rocket:
SHA1(main.go)=  🐦 🚗 🔛 🔮 📆 🍋 🙏 ⚠️🎻 🍭 👊 ♥️™️🚸 ⌛️➕ 🏠 🐜 👏 🚀 
```

Like so!

SHA1(main.go)=  :bird::red_car::on::crystal_ball::calendar::lemon::pray::warning::violin::lollipop::facepunch::hearts::tm::children_crossing::hourglass::heavy_plus_sign::house::ant::clap::rocket:


### pass in the checksums

Rather than relying on this simple tool to do the checksum itself, you will likely want to rely on OpenSSL or coreutils for checksumming.
`emojisum` can just take those formats on stdin:

```bash
$> sha1sum main.go | emojisum -pg
7656835947b4c6da272023c56b6f2529511bf88b  main.go
:jp::gb::metal::goat::family::rocket::smiley_cat::swimmer::chocolate_bar::cactus::candy::smile::honeybee::house::cherries::cloud::fries::bow::wavy_dash::musical_score:  main.go
🇯🇵 🇬🇧 🤘 🐐 👪 🚀 😺 🏊 🍫 🌵 🍬 😄 🐝 🏠 🍒 ☁️🍟 🙇 〰️ 🎼   main.go
```

Like so: 

🇯🇵 🇬🇧 🤘 🐐 👪 🚀 😺 🏊 🍫 🌵 🍬 😄 🐝 🏠 🍒 ☁️🍟 🙇 〰️ 🎼   main.go


```bash
$> openssl sha1 main.go |emojisum -pb
SHA1(main.go)= 7656835947b4c6da272023c56b6f2529511bf88b
SHA1(main.go)= :jp::gb::metal::goat::family::rocket::smiley_cat::swimmer::chocolate_bar::cactus::candy::smile::honeybee::house::cherries::cloud::fries::bow::wavy_dash::musical_score:
SHA1(main.go)= 🇯🇵 🇬🇧 🤘 🐐 👪 🚀 😺 🏊 🍫 🌵 🍬 😄 🐝 🏠 🍒 ☁️🍟 🙇 〰️ 🎼 
```

And like so:

SHA1(main.go)= 🇯🇵 🇬🇧 🤘 🐐 👪 🚀 😺 🏊 🍫 🌵 🍬 😄 🐝 🏠 🍒 ☁️🍟 🙇 〰️ 🎼 


## Library

Use the golang library to access the mapped emoji words:

```golang
package main

import (
  "fmt"

  "github.com/emojisum/emojisum/emoji"
)

func main() {
  for i := 0; i < 255; i++ {
    fmt.Printf("%d (%2.2x):\n", i, i)
    for _, word := range emoji.Map(byte(i)) {
      fmt.Printf("  - %s\n", emoji.CodepointToUnicode(word))
    }
  }
}
```

## Contributed Implementations

### Rust

In `./contrib/emojisum-rs/` you'll find a rust library for emojisum.

#### Build

```shell
cd ./contrib/emojisum-rs/
cargo build
```

#### Test

```shell
cargo test
```

Emoji Working Group

Category: Informational


# emoji checksum mapping

## Status of This Document

This document provides information for the community. This document does not
specify a standard of any kind.  It is open to suggestions and discussion for
improvements.  This document is presently a draft and will apply versioning of
the documents as needed.  Distribution of this document is unlimited.


## Notices

Permission is granted to copy and distribute this document for any purpose and
without charge, including translations into other languages and incorporation
into compilations, provided that the copyright notice and this notice are
preserved, and that any substantive changes or deletions from the original are
clearly marked.

A pointer to the latest version of the canonical JSON is the URL: [http://emoji.thisco.de/draft/emojimap.json](http://emoji.thisco.de/draft/emojimap.json)
A pointer to the latest version of this spec the URL: [http://emoji.thisco.de/draft/](http://emoji.thisco.de/draft/)
Related documentation in can be found at the URL: [http://emoji.thisco.de/](http://emoji.thisco.de/)

## Abstract

This document specifies a practice of mapping an 8bit byte to one of a
corresponding list of 256 emoji strings.

The [`emojimap.json`](./emojimap.json) JSON is the authority of ordering.
While this directory contains golang source that is is importable by golang
projects, other languages can fetch the ordered list of the JSON document and
do their own emojisum comparison or rendering.

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD",
"SHOULD NOT", "RECOMMENDED",  "MAY", and "OPTIONAL" in this document are to be
interpreted as described in [RFC 2119](https://tools.ietf.org/html/rfc2119).


## Introduction

### Purpose

The purpose of this practice is provide simplified way to convey [checksums](https://en.wikipedia.org/wiki/Checksum) for visual comparison.


### Intended Audience

This is intended for use by implementors of software to convey checksums or validate conveyed checksums.


## Details

By operating on an 8bit byte, this provides the oportunity for 256 permutations.
Most checksums convey in a hexadecimal notation, there showing a par of case-insensitive hexadecimal characters per byte (`16*16 = 256`).
Having a mapping of 256 emojis this thereby reduces the number of characters (or emojis) needed to convey the checksum.
In example, a [SHA1](https://en.wikipedia.org/wiki/SHA-1) checksum is 40 hexadecimal characters long, whereas an SHA1-emojisum is only 20 emojis.

## References

* Unicode Technical Report #51 - http://www.unicode.org/reports/tr51/
* http://www.webpagefx.com/tools/emoji-cheat-sheet/

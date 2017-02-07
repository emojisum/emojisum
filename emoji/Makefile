
default: map_gen.go list.md

.PHONY: map_gen.go
map_gen.go: map_json.go map.go emojimap.json
	rm -f $@
	go run map_json.go -template map_gen -in ./emojimap.json -out ./$@

.PHONY: list.md
list.md: map_json.go map.go emojimap.json
	rm -f $@
	go run map_json.go -template markdown_gen -in ./emojimap.json -out ./$@


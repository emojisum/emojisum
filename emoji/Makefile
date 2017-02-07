

.PHONY: map_gen.go
map_gen.go: map_json.go map.go emojimap.json map_gen.tmpl
	rm -f map_gen.go
	go run map_json.go -template map_gen.tmpl -in ./emojimap.json -out ./map_gen.go


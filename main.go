package main

import "github.com/nikitasmall/detector/detector"

const (
	INPUT_FOLDER  = "input"
	OUTPUT_FOLDER = "result"
)

func main() {
	songs := detector.SelectSongs(INPUT_FOLDER)
	detector.CategorizeFolder(OUTPUT_FOLDER, songs)
}

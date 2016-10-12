package detector

import (
	"fmt"
	"io/ioutil"
	"log"
)

func SelectSongs(inputFolder string) []string {
	files, err := ioutil.ReadDir(inputFolder)
	if err != nil {
		log.Panicf("Cannot load input directory: %s", err)
	}

	songs := make([]string, 0)
	for _, file := range files {
		if file.Name() == ".gitignore" {
			continue
		}

		songs = append(songs, fmt.Sprintf("%s/%s", inputFolder, file.Name()))
	}

	return songs
}

func CategorizeFolder(outputFolder string, songs []string) {
	done := make(chan bool, len(songs))

	for _, song := range songs {
		go processSong(outputFolder, song, done)
	}

	results := make([]bool, 0)
	for len(results) < len(songs) {
		results = append(results, <-done)
	}
}

func processSong(outputFolder, song string, done chan bool) {
	b, err := getTrailingBytes(song, 128)
	if err != nil {
		log.Printf("Cannot parse end of file: %s", err)
		done <- false
		return
	}

	tags, err := mp3ID3v1(b).Parse()
	if err != nil {
		log.Printf("Cannot parse tags for song: %s", err)
		done <- false
		return
	}

	err = moveFileByTags(outputFolder, song, tags)
	if err != nil {
		log.Printf("Cannot move song by tags: %s", err)
		done <- false
		return
	}

	done <- true
}

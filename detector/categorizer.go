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
	for _, song := range songs {
		b, err := getTrailingBytes(song, 128)
		if err != nil {
			log.Printf("Cannot parse end of file: %s", err)
			continue
		}

		tags, err := mp3ID3v1(b).Parse()
		if err != nil {
			log.Printf("Cannot parse tags for song: %s", err)
			continue
		}

		moveFileByTags(outputFolder, song, tags)
	}
}

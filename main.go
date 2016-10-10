package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

//getTrailingBytes opens a file and reads the last n bytes
func getTrailingBytes(filename string, n int) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = f.Seek(-int64(n), os.SEEK_END)
	if err != nil {
		return nil, err
	}

	b := make([]byte, n)
	_, err = f.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//Tag represents anything that can produce a list of details
type Tag interface {
	//Parse returns the complete list of all data found in the tag
	Parse() map[string]string
}

//mp3ID3v1 is a specific kind of tagging
type mp3ID3v1 []byte

//Parse decodes the ID3v1 tag
//According to wikipedia, track number is in here somewhere too
//http://en.wikipedia.org/wiki/ID3#Layout
func (mp3 mp3ID3v1) Parse() (map[string]string, error) {
	m := make(map[string]string, 8)
	if string(mp3[:3]) != "TAG" {
		return nil, errors.New("Could not parse tags!")
	}
	m["title"] = string(bytes.Trim(mp3[3:33], "\x00"))
	m["artist"] = string(bytes.Trim(mp3[33:63], "\x00"))
	m["album"] = string(bytes.Trim(mp3[63:93], "\x00"))
	m["year"] = string(mp3[93:97])
	m["comment"] = string(mp3[97:126])
	return m, nil
}

func main() {
	b, err := getTrailingBytes("music/track_01.mp3", 128)
	if err != nil {
		log.Fatal(err)
	}

	tags, err := mp3ID3v1(b).Parse()
	if err == nil {
		artist := strings.Replace(tags["artist"], " ", "_", -1)
		album := strings.Replace(tags["album"], " ", "_", -1)
		title := strings.Replace(tags["title"], " ", "_", -1)

		err = os.MkdirAll(
			fmt.Sprintf(
				"%s/%s/%s",
				"result",
				artist,
				album,
			),
			0777,
		)
		log.Println(err)

		err = os.Rename("music/track_01.mp3",
			fmt.Sprintf(
				"%s/%s/%s/%s.mp3",
				"result",
				artist,
				album,
				title,
			),
		)
		log.Println(err)
	} else {
		log.Print(err)
	}
}

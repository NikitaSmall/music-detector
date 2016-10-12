package detector

import (
	"bytes"
	"errors"
)

var wrongTagFormatError = errors.New("Wrong tag format for file!")

//mp3ID3v1 is a specific kind of tagging
type mp3ID3v1 []byte

//Tag represents anything that can produce a list of details
type Tag interface {
	//Parse returns the complete list of all data found in the tag
	Parse() map[string]string
}

//Parse decodes the ID3v1 tag
//According to wikipedia, track number is in here somewhere too
//http://en.wikipedia.org/wiki/ID3#Layout
func (mp3 mp3ID3v1) Parse() (map[string]string, error) {
	m := make(map[string]string, 8)
	if string(mp3[:3]) != "TAG" {
		return nil, wrongTagFormatError
	}
	m["title"] = string(bytes.Trim(mp3[3:33], "\x00"))
	m["artist"] = string(bytes.Trim(mp3[33:63], "\x00"))
	m["album"] = string(bytes.Trim(mp3[63:93], "\x00"))
	m["year"] = string(bytes.Trim(mp3[93:97], "\x00"))
	m["comment"] = string(bytes.Trim(mp3[97:126], "\x00"))
	return m, nil
}

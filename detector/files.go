package detector

import (
	"fmt"
	"os"
)

// getTrailingBytes opens a file and reads the last n bytes
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

func moveFileByTags(outputFolder, originalSong string, songTags map[string]string) error {
	err := createDirStructure(outputFolder, songTags)
	if err != nil {
		return fmt.Errorf("Error during creating dir structure by tags: %s", err)
	}

	err = moveSongToTagStructure(outputFolder, originalSong, songTags)
	if err != nil {
		return fmt.Errorf("Error during track moving: %s", err)
	}

	return nil
}

func createDirStructure(outputFolder string, songTags map[string]string) error {
	return os.MkdirAll(
		fmt.Sprintf(
			"%s/%s/%s",
			outputFolder,
			songTags["artist"],
			songTags["album"],
		),
		0777,
	)
}

func moveSongToTagStructure(outputFolder, originalSong string, songTags map[string]string) error {
	return os.Rename(originalSong,
		fmt.Sprintf(
			"%s/%s/%s/%s.mp3",
			outputFolder,
			songTags["artist"],
			songTags["album"],
			songTags["title"],
		),
	)
}

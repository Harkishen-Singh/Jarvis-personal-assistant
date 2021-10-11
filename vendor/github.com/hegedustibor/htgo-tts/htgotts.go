package htgotts

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

/**
 * Required:
 * - mplayer
 *
 * Use:
 *
 * speech := htgotts.Speech{Folder: "audio", Language: "en"}
 */

// Speech struct
type Speech struct {
	Folder   string
	Language string
}

// Speak downloads speech and plays it using mplayer
func (speech *Speech) Speak(text string) error {

	fileName := speech.Folder + "/" + text + ".mp3"

	var err error
	if err = speech.createFolderIfNotExists(speech.Folder); err != nil {
		return err
	}
	if err = speech.downloadIfNotExists(fileName, text); err != nil {
		return err
	}

	return speech.play(fileName)
}

/**
 * Create the folder if does not exists.
 */
func (speech *Speech) createFolderIfNotExists(folder string) error {
	dir, err := os.Open(folder)
	if os.IsNotExist(err) {
		return os.MkdirAll(folder, 0700)
	}

	dir.Close()
	return nil
}

/**
 * Download the voice file if does not exists.
 */
func (speech *Speech) downloadIfNotExists(fileName string, text string) error {
	f, err := os.Open(fileName)
	if err != nil {
		url := fmt.Sprintf("http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s", url.QueryEscape(text), speech.Language)
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		output, err := os.Create(fileName)
		if err != nil {
			return err
		}

		_, err = io.Copy(output, response.Body)
		return err
	}

	f.Close()
	return nil
}

/**
 * Play voice file.
 */
func (speech *Speech) play(fileName string) error {
	mplayer := exec.Command("mplayer", "-cache", "8092", "-", fileName)
	return mplayer.Run()
}

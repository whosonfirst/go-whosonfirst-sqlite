package tts

import (
	"errors"
	"github.com/whosonfirst/go-writer-tts/speakers"
)

func NewSpeakerForEngine(engine string, options ...interface{}) (speakers.Speaker, error) {

	if engine == "osx" {
		return speakers.NewOSXSpeaker()
	}

	return nil, errors.New("Unknown or unsupported text to speech engine")
}

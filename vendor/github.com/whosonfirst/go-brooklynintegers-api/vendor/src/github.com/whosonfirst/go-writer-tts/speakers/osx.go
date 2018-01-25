package speakers

import (
	"github.com/everdev/mack"
	"io"
	"io/ioutil"
	"strings"
)

type OSXSpeaker struct {
	Speaker
}

func NewOSXSpeaker() (*OSXSpeaker, error) {

	s := OSXSpeaker{}
	return &s, nil
}

func (s *OSXSpeaker) Read(reader io.Reader) error {

	tee := io.TeeReader(reader, s)
	_, err := ioutil.ReadAll(tee)
	return err
}

func (s *OSXSpeaker) WriteString(text string) (int64, error) {
	r := strings.NewReader(text)
	return r.WriteTo(s)
}

func (s *OSXSpeaker) Write(p []byte) (int, error) {

	var text string
	text = string(p[:])

	mack.Say(text)

	count := len(text)
	return count, nil
}

func (s *OSXSpeaker) Close() error {
	return nil
}

package speakers

import (
	"io"
)

type Speaker interface {
	Read(io.Reader) error
	WriteString(text string) (int64, error)
	Write(p []byte) (int, error)
	Close() error
}

package gonfs

import (
	"io"
)

type xdrState struct {
	// err != nil means error state
	err error

	// reader != nil means we are reading
	reader io.Reader

	// writer != nil means we are writing
	writer io.Writer
}

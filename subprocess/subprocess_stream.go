package wzlib_subprocess

import (
	"io/ioutil"
	"os"
	"strings"
)

// ProcessStream object
type ProcessStream struct {
	filePipe *os.File
}

// NewProcessStream creates a ProcessStream instance. Management of the pipe file is solely on module caller.
func NewProcessStream(fname string) *ProcessStream {
	var err error
	zs := new(ProcessStream)
	ioutil.WriteFile(fname, []byte(""), 0644)
	zs.filePipe, err = os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return zs
}

// Write data to the underlying pipe file
func (zs *ProcessStream) Write(data []byte) (n int, err error) {
	line := strings.TrimSpace(string(data)) + "\n"
	zs.filePipe.WriteString(line)
	return len(data), nil
}

// Close stream
func (zs *ProcessStream) Close() error {
	return zs.filePipe.Close()
}

package wzlib_subprocess

import (
	"bufio"
	"bytes"
	"io"
)

const (
	// stdoutBufSize is the size of the buffers given to a sub-process stdout
	stdoutBufSize = 16384
)

type BufferedCmd struct {
	*Cmd

	Stdin  io.WriteCloser
	Stdout *bufio.Reader
	Stderr *bufio.Reader
}

func (bc *BufferedCmd) buf2String(reader io.Reader) string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (bc *BufferedCmd) StdoutString() string {
	return bc.buf2String(bc.Stdout)
}

func (bc *BufferedCmd) StderrString() string {
	return bc.buf2String(bc.Stderr)
}

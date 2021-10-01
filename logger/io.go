package logger

import (
	"fmt"
	"io"
)

type IOLogger struct{
	nextreader io.Reader
	nextwriter io.Writer
	Enabled bool ///if you want to leave this logger in place in the prod version and be able to enable/disable it
}


/**
reader - [optional] The real io.Reader implementation
writer - [optional] The real io.Writer implementation

Note: if either reader or writer is nil and something tries to read/write using this
		object, it will panic
	In other words, reader/writer is optional provided you don't try to use them
 */
func NewIOLogger(reader io.Reader, writer io.Writer)*IOLogger{
	return &IOLogger{
		nextreader: reader,
		nextwriter: writer,
		Enabled: true,
	}
}

func (p *IOLogger) Read(data []byte) (n int, err error) {
	if p.Enabled {
		fmt.Println(data)
	}
	return p.nextreader.Read(data)
}

func (p *IOLogger) Write(data []byte) (n int, err error) {
	if p.Enabled {
		fmt.Println(data)
	}
	return p.nextwriter.Write(data)
}

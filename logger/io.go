package logger

import (
	"fmt"
	"io"
)

type IOLogger struct{
	nextreader io.Reader
	nextwriter io.Writer
	Enabled bool ///if you want to leave this logger in place in the prod version and be able to enable/disable it
	MaxLogSize int
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
		MaxLogSize: 2000,
	}
}

func (p *IOLogger)log(data []byte){
	if p.Enabled {
		if len(data) > p.MaxLogSize{
			fmt.Println(string(data[:p.MaxLogSize]))
		}else {
			fmt.Println(string(data))
		}
	}
}

func (p *IOLogger) Read(data []byte) (n int, err error) {
	n, err = p.nextreader.Read(data)
	p.log(data)
	return n, err
}

func (p *IOLogger) Write(data []byte) (n int, err error) {
	p.log(data)
	return p.nextwriter.Write(data)
}

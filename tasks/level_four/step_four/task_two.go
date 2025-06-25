// "БОЛЬШИЕ СТРОКИ"

package main

import (
	"strings"
)

type Writer interface {
	Write(p []byte) int
}

type Reader interface {
	Read() []byte
}

type ReaderWriter interface {
	Writer
	Reader
}

type UpperReaderWriter struct {
	UpperString string
}

func (s *UpperReaderWriter) Write(p []byte) int {
	s.UpperString = strings.ToUpper(string(p))
	return len(p)
}

func (s *UpperReaderWriter) Read() []byte {
	return []byte(s.UpperString)
}

package main

import "fmt"

type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Write(string) (int, error)
}

type ReaderWriter interface {
	Reader
	Writer
}
type HelloReader struct {
}

func (r *HelloReader) Read() (string, error) {
	return "Hello there", nil
}

type HelloWriter struct {
}

func (w *HelloWriter) Write(message string) (int, error) {
	n, err := fmt.Println(message)
	return n, err
}

type RWStruct struct {
	*HelloReader
	*HelloWriter
}

func NewRWStruct() *RWStruct {
	return &RWStruct{&HelloReader{}, &HelloWriter{}}
}

func main() {
	n := NewRWStruct()
	v, _ := n.Read()
	fmt.Println("Read using Using NewRWStruct", v)
	n.Write("Using NewRWStruct")
}

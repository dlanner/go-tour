package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (myreader MyReader) Read(b []byte) (int, error) {
	return copy(b, []byte{'A'}), nil
}

func main() {
	reader.Validate(MyReader{})
}

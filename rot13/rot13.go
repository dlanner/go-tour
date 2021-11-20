package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = reader.r.Read(p)
	var rot13 []byte
	for i := 0; i < len(p) && p[i] > 0; i++ {
		x := p[i]
		if x >= 'A' && x <= 'Z' {
			x = ((x - 'A' + 13) % 26) + 'A'
		} else if x >= 'a' && x <= 'z' {
			x = ((x - 'a' + 13) % 26) + 'a'
		}
		rot13 = append(rot13, x)
	}
	n, err = reader.r.Read(rot13)
	fmt.Printf("%s\n", rot13)
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n,err = rot.r.Read(p)
	for i := range p{
		if ( 'A' <= p[i]) && (p[i] <= 'Z') {
			p[i] = 65 + (((p[i] - 65) + 13) % 26)
		} else if ('a' <= p[i]) && (p[i] <= 'z') {
			p[i] = 97 + (((p[i] - 97) + 13) % 26)
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func (rot rot13Reader) Read(char []byte) (int, error) {
	n, err := rot.r.Read(char)
	for i := 0; i < n; i++ {
		char[i], _ = deCipher(isAlpha(char[i]))
	}
	return n, err
}

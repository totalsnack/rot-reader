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

func deCipher(char byte, register byte) (byte, byte) {
	if register == 0 {
		return char, 0
	}
	return register + (char-register+13)%26, 1
}

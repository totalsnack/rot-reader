package rot_reader

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(char []byte) (int, error) {
	n, err := rot.r.Read(char)
	for i := 0; i < n; i++ {
		char[i], _ = deCipher(isAlpha(char[i]))
	}
	return n, err
}

func isAlpha(char byte) (byte, byte) {
	if char >= 'A' && char <= 'Z' {
		return char, 'A'
	} else if char >= 'a' && char <= 'z' {
		return char, 'a'
	} else {
		return char, 0
	}
}

func deCipher(char byte, register byte) (byte, byte) {
	if register == 0 {
		return char, 0
	}
	return register + (char-register+13)%26, 1
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

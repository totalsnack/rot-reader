func (rot rot13Reader) Read(char []byte) (int, error) {
n, err := rot.r.Read(char)
for i := 0; i < n; i++ {
char[i], _ = deCipher(isAlpha(char[i]))
}
return n, err
}, 0
}

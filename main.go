func isAlpha(char byte) (byte, byte) {
	if char >= 'A' && char <= 'Z' {
		return char, 'A'
	} else if char >= 'a' && char <= 'z' {
		return char, 'a'
	} else {
		return char, 0
	}
}

package runes

func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func IsWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\r' || r == '\t'
}

func IsLeftBrace(r rune) bool {
	return r == '{'
}

func IsRightBrace(r rune) bool {
	return r == '}'
}

func IsColon(r rune) bool {
	return r == ':'
}

func IsComma(r rune) bool {
	return r == ','
}

func IsLeftBracket(r rune) bool {
	return r == '['
}

func IsRightBracket(r rune) bool {
	return r == ']'
}

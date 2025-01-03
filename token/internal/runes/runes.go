package runes

func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func IsHex(r rune) bool {
	return IsDigit(r) || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F')
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

func MaybeTrue(data []rune) bool {
	return 4 <= len(data) && string(data[:4]) == "true"
}

func MaybeFalse(data []rune) bool {
	return 5 <= len(data) && string(data[:5]) == "false"
}

func MaybeNull(data []rune) bool {
	return 4 <= len(data) && string(data[:4]) == "null"
}

func MaybeString(data []rune) bool {
	return 2 <= len(data) && data[0] == '"'
}

func MaybeNumber(r rune) bool {
	return IsDigit(r) || r == '-'
}

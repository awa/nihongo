package nihongo

import (
	"bytes"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

// Normalize japanese text which will convert with NFKC normalization.
// Hankaku-Kana -> Zenkaku-Kana
// Zenkaku special chars -> Hankaku special chars
func Normalize(text string) string {
	return norm.NFKC.String(text)
}

// ContainsHiragana returns true when text contains hiragana
func ContainsHiragana(text string) bool {
	for _, r := range text {
		if unicode.In(r, unicode.Hiragana) {
			return true
		}
	}
	return false
}

// ContainsKatakana returns true when text contains katakana
func ContainsKatakana(text string) bool {
	for _, r := range text {
		if unicode.In(r, unicode.Katakana) {
			return true
		}
	}
	return false
}

// ToHiragana converts all katakana text to hiragana.
// You should normalize text before converting.
func ToHiragana(text string) string {
	var buf bytes.Buffer
	for _, r := range text {
		if unicode.In(r, unicode.Katakana) {
			// Convert to hiragana
			r -= 0x60
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

// ToKatakana converts all hiragana text to katakana.
// You should normalize text before converting.
func ToKatakana(text string) string {
	var buf bytes.Buffer
	//buf := bytes.NewBuffer(make([]byte, len(text)))
	for _, r := range text {
		if unicode.In(r, unicode.Hiragana) {
			// Convert to hiragana
			r += 0x60
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

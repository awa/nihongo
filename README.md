NihonGo is japanese utility for Golang
====

[![Build Status](https://travis-ci.org/dogenzaka/nihongo.svg?branch=master)](https://travis-ci.org/dogenzaka/nihongo)
[![Coverage Status](https://coveralls.io/repos/dogenzaka/nihongo/badge.svg)](https://coveralls.io/r/dogenzaka/nihongo)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/dogenzaka/nihongo/blob/master/LICENSE)

NihonGo is japanese utility on Go.

```
go get github.com/dogenzaka/nihongo
```

Featuers
--

- Katakana / Hiragana conversion
- Unicode normalization
- Detect katakana / hiragana string
- Simple tokenizer ported TinySegmenter

Examples
--

```go
import (
  "fmt"
  "github.com/dogenzaka/nihongo"
)

func TestNormalize() {
  normalized := nihongo.Normalize("テストﾃｽﾄ＋＝")
  fmt.Println(normalized) // テストテスト+=
}

func TestToHiragana() {
  hira := nihongo.ToHiragana("テストてすと")
  fmt.Println(hira) // てすとてすと
}

func TestToKatakana() {
  kana := nihongo.ToKatakana("テストてすと")
  fmt.Println(kana) // テストテスト
}

func TestTokenize() {
  words := nihongo.Tokenize("私は人間です")
  fmt.Println(words) // ["私" "は" "人間" "です"]
}

func TestContainsHiragana() {
  nihongo.ContainsHiragana("ひらがな") // true
  nihongo.ContiansHiragana("日本語") // false
}

func TestContainsKatakana() {
  nihongo.ContainsKatakana("カタカナ") // true
  nihongo.ContiansKatakana("日本語") // false
}
```

License
--
ISC


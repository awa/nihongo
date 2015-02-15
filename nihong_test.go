package nihongo

import (
	"testing"
)

func TestNormalize(t *testing.T) {

	norm := Normalize("テストてすと")
	if norm != "テストてすと" {
		t.Errorf("Converted strings which should not be converted %v", norm)
	}

	norm = Normalize("テストﾃｽﾄ／＋")
	if norm != "テストテスト/+" {
		t.Errorf("Converted strings which should not be converted %v", norm)
	}

}

func TestToHiragana(t *testing.T) {

	hira := ToHiragana("テスト")
	if hira != "てすと" {
		t.Errorf("Not converted to Hiragana %v", hira)
	}

	hira = ToHiragana("テスト混合てすと")
	if hira != "てすと混合てすと" {
		t.Errorf("Not converted to Hiragana %v", hira)
	}

	hira = ToHiragana("Englishテスト混合")
	if hira != "Englishてすと混合" {
		t.Errorf("Not converted to Hiragana %v", hira)
	}

	hira = ToHiragana("アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ")
	if hira != "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽ" {
		t.Errorf("Not converted to Hiragana %v", hira)
	}
}

func TestToKatakana(t *testing.T) {

	kana := ToKatakana("てすと")
	if kana != "テスト" {
		t.Errorf("Not converted to Katakana %v", kana)
	}

	kana = ToKatakana("てすと混合テスト")
	if kana != "テスト混合テスト" {
		t.Errorf("Not converted to Katakana %v", kana)
	}

	kana = ToKatakana("てすと混合English")
	if kana != "テスト混合English" {
		t.Errorf("Not converted to Katakana %v", kana)
	}

	kana = ToKatakana("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽ")
	if kana != "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ" {
		t.Errorf("Not converted to Katakana %v", kana)
	}
}

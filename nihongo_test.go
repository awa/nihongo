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

func TestContainsHiragana(t *testing.T) {
	checkContainsHiragana(t, "テスト", false)
	checkContainsHiragana(t, "English", false)
	checkContainsHiragana(t, "てすと", true)
	checkContainsHiragana(t, "テストてすと", true)
}

func TestContainsKatakana(t *testing.T) {
	checkContainsKatakana(t, "テスト", true)
	checkContainsKatakana(t, "English", false)
	checkContainsKatakana(t, "てすと", false)
	checkContainsKatakana(t, "テストてすと", true)
}

func checkContainsHiragana(t *testing.T, text string, expected bool) {
	contains := ContainsHiragana(text)
	if contains == expected {
		return
	}
	if contains {
		t.Errorf("ContainsHiragana detected hiragana on %v", text)
	} else {
		t.Errorf("ContainsHiragana did not detect hiragana on %v", text)
	}
}

func checkContainsKatakana(t *testing.T, text string, expected bool) {
	contains := ContainsKatakana(text)
	if contains == expected {
		return
	}
	if contains {
		t.Errorf("ContainsKatakana detected katakana on %v", text)
	} else {
		t.Errorf("ContainsKatakana did not detect katakana on %v", text)
	}
}

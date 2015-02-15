package nihongo

import (
	"testing"
)

func TestSegmenter(t *testing.T) {

	testEqual(
		t,
		Tokenize(""),
		[]string{},
	)

	testEqual(
		t,
		Tokenize("私は人間です"),
		[]string{"私", "は", "人間", "です"},
	)

	testEqual(
		t,
		Tokenize("私はロボットです"),
		[]string{"私", "は", "ロボット", "です"},
	)

	testEqual(
		t,
		Tokenize("This is an English sentence."),
		[]string{"This", " ", "is", " ", "an", " ", "English", " ", "sentence", "."},
	)

}

func testEqual(t *testing.T, a1 []string, a2 []string) {
	if len(a1) != len(a2) {
		t.Errorf("Segments should be equals %q <-> %q", a1, a2)
		return
	}
	for i, str := range a1 {
		if str != a2[i] {
			t.Errorf("Index %d is not equals for %q <-> %q", i, a1, a2)
		}
	}
}

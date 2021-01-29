package soundex

import (
	"testing"
)

func TestOrigin(t *testing.T) {

	snd, err := New("origin")
	if err != nil {
		t.Fatal("origin algorithm not found")
	}

	tF := func(src string, res string, err error) {
		sr, se := snd.Code(src)
		if sr != res || se != err {
			t.Fatalf("origin test failed src=%s wait=%s ret=%s we=%v re=%v", src, res, sr, err, se)
		}
	}

	tF("ammonium", "A555", nil)
	tF("implementation", "I514", nil)
	tF("robert", "R163", nil)
	tF("rupert", "R163", nil)
	tF("rubin", "R150", nil)
	tF("ashcraft", "A261", nil)
	tF("ashcroft", "A261", nil)
	tF("tymczak", "T522", nil)
	tF("morphs", "M612", nil)
	tF("a", "A000", nil)
	tF("pa", "P000", nil)
	tF("pap", "P100", nil)
	tF("", "", ErrEmptyString)
	tF("13123", "", ErrInvalidCharacter)
}

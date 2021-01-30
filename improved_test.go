package soundex

import (
	"testing"
)

func TestImproved(t *testing.T) {

	snd, err := New(AlgoImproved)
	if err != nil {
		t.Fatal("origin algorithm not found")
	}

	tF := func(src string, res string, err error) {
		sr, se := snd.Code(src)
		if sr != res || se != err {
			t.Fatalf("improved test failed src=%s wait=%s ret=%s we=%v re=%v", src, res, sr, err, se)
		}
	}

	tF("morphs", "M913", nil)
	tF("aquaman", "A588", nil)
	tF("goprog", "G194", nil)
	tF("traderworld", "T969976", nil)
	tF("", "", ErrEmptyString)
	tF("13123", "", ErrInvalidCharacter)
}

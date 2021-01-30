package soundex

import (
	"testing"
)

func TestSoundex(t *testing.T) {
	_, err := New(41231234)
	if err != ErrUnsupportedAlgorithm {
		t.Fatal("error expected")
	}

	_, err = New(41231234, 2134)
	if err != ErrUnsupportedAlgorithm {
		t.Fatal("error expected")
	}
}

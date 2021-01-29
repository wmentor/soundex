package soundex

import (
	"testing"
)

func TestSoundex(t *testing.T) {
	_, err := New("unsupported")
	if err != ErrUnsupportedAlgorithm {
		t.Fatal("error expected")
	}
}

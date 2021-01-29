package soundex

import (
	"errors"
)

var (
	ErrUnsupportedAlgorithm error = errors.New("unsupported algorithm")
	ErrInvalidCharacter     error = errors.New("invalid character")
	ErrEmptyString          error = errors.New("empty string")
)

type Soundex interface {
	Code(string) (string, error)
}

func New(algo string) (Soundex, error) {

	switch algo {
	case "origin", "":
		return &soundexOrigin{}, nil

	default:
		return nil, ErrUnsupportedAlgorithm
	}
}

package soundex

import (
	"errors"
)

var (
	ErrUnsupportedAlgorithm error = errors.New("unsupported algorithm")
	ErrInvalidCharacter     error = errors.New("invalid character")
	ErrEmptyString          error = errors.New("empty string")
)

const (
	AlgoOrigin uint64 = 1 << iota
	AlgoImproved
)

type Soundex interface {
	Code(string) (string, error)
}

func New(algorithm ...uint64) (Soundex, error) {

	alg := AlgoOrigin

	if len(algorithm) > 1 {
		return nil, ErrUnsupportedAlgorithm
	} else if len(algorithm) == 1 {
		alg = algorithm[0]
	}

	switch alg {
	case AlgoOrigin:
		return &soundexOrigin{}, nil

	case AlgoImproved:
		return &soundexImproved{}, nil

	default:
		return nil, ErrUnsupportedAlgorithm
	}
}

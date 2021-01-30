package soundex

import (
	"unicode"
)

type soundexImproved struct {
}

var (
	improvedAllowed map[rune]bool
	improvedMap     map[rune]rune

	_ Soundex = (*soundexImproved)(nil)
)

func init() {
	improvedAllowed = make(map[rune]bool, 64)
	for _, r := range "QWERTYUIOPASDFGHJKLZXCVBNM" {
		improvedAllowed[r] = true
	}

	improvedMap = make(map[rune]rune, 64)

	for i, str := range []string{"BP", "FV", "CKS", "GJ", "QXZ", "DT", "L", "MN", "R"} {
		for _, r := range str {
			improvedMap[r] = '1' + rune(i)
		}
	}

}

func (s *soundexImproved) Code(src string) (string, error) {

	if len(src) == 0 {
		return "", ErrEmptyString
	}

	list := make([]rune, 0, 64)

	for _, r := range src {
		ur := unicode.ToUpper(r)
		if improvedAllowed[ur] {
			list = append(list, ur)
		} else {
			return "", ErrInvalidCharacter
		}
	}

	j := 0

	lr := rune(0)

	for i, r := range list {
		if i == 0 {
			j++
		} else {
			if r != 'H' && r != 'W' {
				if nr := improvedMap[list[i]]; nr != 0 {
					if lr != nr {
						list[j] = nr
						j++
						lr = nr
					}
				} else {
					lr = 0
				}
			}
		}
	}

	list = list[:j]

	return string(list), nil
}

package soundex

import (
	"unicode"
)

type soundexOrigin struct {
}

var (
	originAllowed map[rune]bool
	originMap     map[rune]rune

	_ Soundex = (*soundexOrigin)(nil)
)

func init() {
	originAllowed = make(map[rune]bool, 64)
	for _, r := range "QWERTYUIOPASDFGHJKLZXCVBNM" {
		originAllowed[r] = true
	}

	originMap = make(map[rune]rune, 64)

	for i, str := range []string{"BFPV", "CGJKQSXZ", "DT", "L", "MN", "R"} {
		for _, r := range str {
			originMap[r] = '1' + rune(i)
		}
	}

}

func (s *soundexOrigin) Code(src string) (string, error) {

	if len(src) == 0 {
		return "", ErrEmptyString
	}

	list := make([]rune, 0, 64)

	for _, r := range src {
		ur := unicode.ToUpper(r)
		if originAllowed[ur] {
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
				if nr := originMap[list[i]]; nr != 0 {
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

	if j > 4 {
		j = 4
	}

	list = list[:j]

	for i := j; i < 4; i++ {
		list = append(list, '0')
	}

	return string(list), nil
}

package acronym

import (
	"strings"
	"unicode"
)

func parseSymbol(s string) (bool, int) {
	for i, v := range s {
		if !unicode.IsLetter(v) {
			return true, i
		}
	}
	return false, -1
}

func splitter(s, del string) string {
	slist := strings.Split(s, del)
	var data string
	if len(slist) == 0 {
		return ""
	}
	for i := 0; i < len(slist); i++ {
		if len(slist[i]) > 1 {
			data += string(slist[i][0])

		}
	}
	return data
}

// Abbreviate given string
func Abbreviate(s string) string {
	list := strings.Fields(s)
	var data string
	for _, v := range list {
		if ok, index := parseSymbol(v); !ok {
			data += string(v[0])
		} else {
			// v = strings.Trim(v, string(v[0]))
			v := splitter(v, string(v[index]))
			if v != "" {
				data += v
			}
		}
	}
	return strings.ToUpper(data)
}

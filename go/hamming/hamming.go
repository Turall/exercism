package hamming

import "errors"

// Distance check distance beetween two DNA
func Distance(a, b string) (int, error) {
	//return error if len is different
	if len(a) != len(b) {
		return 0, errors.New("Length of DNA are different")
	}

	count := 0
	for i, v := range a {
		if rune(b[i]) != v {
			count++
		}
	}
	// another way
	// for i := 0; i < len(a); i++ {
	// 	if b[i] != a[i] {
	// 		count++
	// 	}
	// }
	return count, nil

}

package raindrops

import (
	"strconv"
)

// Convert convert a number into a string that contains raindrop sounds
func Convert(num int) string {
	// create array with integer keys because map is not ordered
	n := [3]int{3, 5, 7}
	m := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	var s string
	for _, v := range n {
		if num%v == 0 {
			s += m[v]
		}
	}
	if s != "" {
		return s
	}
	s = strconv.Itoa(num)
	return s
}

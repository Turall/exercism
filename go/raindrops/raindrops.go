package raindrops

import (
	"strconv"
)

// Convert convert a number into a string that contains raindrop sounds
func Convert(num int) string {
	var s string
	if num%3 == 0 {
		s += "Pling"
	}
	if num%5 == 0 {
		s += "Plang"
	}
	if num%7 == 0 {
		s += "Plong"
	}
	if s != "" {
		return s
	}
	return strconv.Itoa(num)

}

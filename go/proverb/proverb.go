package proverb

// Proverb func.
func Proverb(rhyme []string) []string {
	list := []string{}
	if len(rhyme) == 0 {
		return list
	}
	if len(rhyme) >= 2 {
		for i := 0; i < len(rhyme); i++ {
			temp := rhyme[i+1:]
			for j := 0; j < len(temp); j++ {
				s := "For want of a " + rhyme[i] + " the " + temp[j] + " was lost."
				list = append(list, s)
				break
			}
		}
	}

	last := "And all for the want of a " + rhyme[0] + "."
	list = append(list, last)
	return list
}

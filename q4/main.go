package main

import "fmt"

func main() {
	arr := []string{":)", ";(", ";}", ":-D"}
	fmt.Println(CountSmileys(arr))
}

// Constraint
// 1. must have eyes      : or ;
// 2. optional nose       - or ~
// 3. must have mouth     ) or D
// 4. No additional characters are allowed except for those mentioned.
func CountSmileys(smiles []string) int {
	count := 0
	for _, smile := range smiles {
		// valid length of smiley face must be 2(without nose) or 3(with nose)
		if len(smile) != 2 && len(smile) != 3 {
			continue
		}

		// check valid according to length of the smile string
		switch len(smile) {
		case 2:
			eyes := smile[0]
			mouth := smile[1]
			// check eye and mouth
			if (eyes != ':' && eyes != ';') || (mouth != ')' && mouth != 'D') {
				// skip invalid smile
				continue
			}
		case 3:
			eyes := smile[0]
			nose := smile[1]
			mouth := smile[2]
			// check eye, nose and mouth
			if (eyes != ':' && eyes != ';') || (nose != '-' && nose != '~') || (mouth != ')' && mouth != 'D') {
				// skip invalid smile
				continue
			}
		}
		count++
	}
	return count
}

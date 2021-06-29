package kata

import (
  "fmt"
  "strings"
)

func Decode(roman string) int {
	var x int = 0
	words := strings.Split(roman, "")
	var nextLetter string
	var prevLetter string
	for i, letter := range words {
		fmt.Println("it's X!", x)
		fmt.Println("================================")
		fmt.Println("Len words", len(words))
		fmt.Println((i + 1) >= (len(words)))

		if (i + 1) >= (len(words)) {
			fmt.Println("(i + 1)  > (len(words)) condition")
			nextLetter = words[i+0]
		} else {
			nextLetter = words[i+1]
		}

		if !((i - 1) < 0) {
			prevLetter = words[i-1]
		}

		fmt.Println("prevLetter =>", prevLetter)
		fmt.Println(strings.Contains("VXLCDM", letter))

		fmt.Println("nextLetter =>", nextLetter)
		fmt.Println(strings.Contains("VXLCDM", nextLetter))

		fmt.Println("first condition handle left I", (letter == "I" && strings.Contains("VXLCDM", nextLetter)))
		fmt.Println("second condition handle left I", (prevLetter == "I" && strings.Contains("VXLCDM", letter)))
		fmt.Println("both condition handle left I", (letter == "I" && strings.Contains("VXLCDM", nextLetter) || prevLetter == "I" && strings.Contains("VXLCDM", letter)))

		if letter == "I" && strings.Contains("VXLCDM", nextLetter) || prevLetter == "I" && strings.Contains("VXLCDM", letter) {
			x = leftRomanNumber(letter, x)
			fmt.Println("--I'm inside if--")
			
		} else {
			x = rightRomanNumber(letter, x)
			fmt.Println("--I'm inside else--")
			
		}

	}

	fmt.Println("Final result ", x)
	return x
}

func leftRomanNumber(letter string, x int) int {
		switch letter {
			case "I":
				x += 1
			case "V":
				x = 5 - x
			case "X":
				x = 10 - x
			case "L":
				x = 50 - x
			case "C":
				x = 100 - x
			case "D":
				x = 500 - x
			case "M":
				x = 1000 - x
			default:
				x += 0
		}
	return x
}

func rightRomanNumber(letter string, x int) int {
		switch letter {
			case "I":
				x += 1
			case "V":
				x += 5
			case "X":
				x += 10
			case "L":
				x += 50
			case "C":
				x += 100
			case "D":
				x += 500
			case "M":
				x += 1000
			default:
				x += 0
		}
	return x
}


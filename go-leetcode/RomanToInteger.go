// 13. Roman to Integer

package main

import (
	"strings"
)

func romanToInt(s string) int {
	charArr := strings.Split(s, "")

	// for key, value := range charArr {
	// 	fmt.Println(strconv.Itoa(key) + value)
	// }

	sum := 0

	var prev = ""

	for _, value := range charArr {
		switch value {
		case "I":
			sum += 1
			break
		case "V":
			if prev == "I" {
				sum += 3
			} else {
				sum += 5
			}
			break
		case "X":
			if prev == "I" {
				sum += 8
			} else {
				sum += 10
			}
			break
		case "L":
			if prev == "X" {
				sum += 30
			} else {
				sum += 50
			}
			break
		case "C":
			if prev == "X" {
				sum += 80
			} else {
				sum += 100
			}
			break
		case "D":
			if prev == "C" {
				sum += 300
			} else {
				sum += 500
			}
			break
		case "M":
			if prev == "C" {
				sum += 800
			} else {
				sum += 1000
			}
			break
		}
		prev = value
	}

	return sum
}

// func main() {
// 	fmt.Println(romanToInt("MCMXCIV"))
// }

// 14. Longest Common Prefix
package goleetcode

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	var result string
	strin := strs[0]

	charArr := strings.Split(strin, "")
	j := 1
	for _, c := range charArr {
		i := 0
		for _, str := range strs {
			if strings.HasPrefix(str, strin[0:j]) {
				i++
			}
			if i == len(strs) {
				result = result + c
			}
			j++
		}
	}

	return result
}

func main() {

	strs := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}

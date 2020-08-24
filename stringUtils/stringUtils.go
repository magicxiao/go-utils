package stringUtils

import "strings"

func Contains(s string, substr ...string) bool {
	var flag bool
	for _, str := range substr {
		flag = strings.Contains(s, str)
		if flag {
			break
		}
	}

	return flag
}

package main

func gcdOfStrings2(str1 string, str2 string) string {
	i := 0
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}
	// find
	for i = len(str1); i >= 1; i-- {
		sub := str1[:i]
		// skip
		if len(str2)%len(sub)+len(str1)%len(sub) != 0 {
			continue
		}
		flag := true
		for j := 0; j < len(str2); j += len(sub) {
			if sub != str2[j:j+len(sub)] {
				flag = false
				break
			}
		}
		for k := 0; k < len(str1); k += len(sub) {
			if sub != str1[k:k+len(sub)] {
				flag = false
				break
			}
		}
		if flag {
			return sub
		}
	}
	return ""
}

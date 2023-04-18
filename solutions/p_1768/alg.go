package main

func mergeAlternately(word1 string, word2 string) string {
	var i int
	result := make([]byte, 0, len(word1)+len(word2))
	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			result = append(result, word1[i])
		}
		if i < len(word2) {
			result = append(result, word2[i])
		}
		i++
	}
	return string(result)
}

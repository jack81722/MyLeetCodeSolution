package main

func convertToTitle(columnNumber int) string {
	var bs []byte
	for columnNumber > 0 {
		bs = append(bs, byte((columnNumber-1)%26+'A'))
		columnNumber -= 1
		columnNumber /= 26
	}
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
	}
	return string(bs)
}

package main

func encrypt(txt string) string {
	runes := []rune(txt)
	for index, letter := range runes {
		runes[index] = letter*'1' - '1'
	}
	return string(runes)
}

package reversestring

func ReverseString(word string) string {
	result := ""
	for _, v := range word {
		result = string(v) + result
	}
	return result
}
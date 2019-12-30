package reverse

// Reverse a string
func Reverse(in string) (revd string) {
	revd = ""
	for _, ch := range in {
		revd = string(ch) + revd
	}
	return revd
}

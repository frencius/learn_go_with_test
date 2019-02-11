package iteration

// Repeat and print the char argument
func Repeat(char string, n int) string {
	var repeatedChar string
	for i := 0; i < n; i++ {
		repeatedChar += char
	}
	return repeatedChar
}

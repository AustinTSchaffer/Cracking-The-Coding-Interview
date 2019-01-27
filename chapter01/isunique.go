package chapter01

func LetterHistogram(someString string) map[rune]int {
	letterFreq := make(map[rune]int)
	for _, runeValue := range someString {
		letterFreq[runeValue]++
	}
	return letterFreq
}

func IsUnique(someString string) bool {

	return false
}

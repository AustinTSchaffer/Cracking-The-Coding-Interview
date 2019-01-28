package main

import (
	"fmt"

	"github.com/AustinTSchaffer/CrackingTheCodingInterview/chapter01"
)

func main() {
	fmt.Println("Testing IsUnique")
	fmt.Println(chapter01.IsUnique("test"))

	fmt.Println("Testing StringsArePermutations")
	fmt.Println(chapter01.StringsArePermutations("tes", "ste"))
	fmt.Println(chapter01.StringsArePermutations("tes", "test"))
	fmt.Println(chapter01.StringsArePermutations("test", "test"))

	fmt.Println("Testing URLIfy")
	fmt.Println(chapter01.URLIfy("Luke, I am your father       "))

	fmt.Println("Testing IsAPermutationOfAPalindrome")
	fmt.Println(chapter01.IsAPermutationOfAPalindrome("aabbcccddee"))
	fmt.Println(chapter01.IsAPermutationOfAPalindrome("aabbCcddddEeFfGgGgGGggggg"))
	fmt.Println(chapter01.IsAPermutationOfAPalindrome("aabbCcddddEeFfgGgGGgggggX"))
	fmt.Println(chapter01.IsAPermutationOfAPalindrome("aabbCcddddEeFfgGgGGgggggXYYY"))

	fmt.Println("Testing OneAway")
	fmt.Println(chapter01.OneAway("asdf", "asdf"))
	fmt.Println(chapter01.OneAway("adf", "asdf"))
	fmt.Println(chapter01.OneAway("asddf", "asdf"))
	fmt.Println(chapter01.OneAway("addf", "asdf"))
}

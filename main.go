package main

import (
	"fmt"

	"github.com/AustinTSchaffer/CrackingTheCodingInterview/chapter01"
)

func main() {
	fmt.Println(chapter01.IsUnique("test"))

	fmt.Println(chapter01.StringsArePermutations("tes", "ste"))
	fmt.Println(chapter01.StringsArePermutations("tes", "test"))
	fmt.Println(chapter01.StringsArePermutations("test", "test"))

	fmt.Println(chapter01.URLIfy("Luke, I am your father       "))
}

package main

import (
	"fmt"

	c01 "github.com/AustinTSchaffer/CrackingTheCodingInterview/chapter01"
)

func main() {
	fmt.Println("Testing IsUnique")
	fmt.Println(c01.IsUnique("test"))

	fmt.Println("Testing StringsArePermutations")
	fmt.Println(c01.StringsArePermutations("tes", "ste"))
	fmt.Println(c01.StringsArePermutations("tes", "test"))
	fmt.Println(c01.StringsArePermutations("test", "test"))

	fmt.Println("Testing URLIfy")
	fmt.Println(c01.URLIfy("Luke, I am your father       "))

	fmt.Println("Testing IsAPermutationOfAPalindrome")
	fmt.Println(c01.IsAPermutationOfAPalindrome("aabbcccddee"))
	fmt.Println(c01.IsAPermutationOfAPalindrome("aabbCcddddEeFfGgGgGGggggg"))
	fmt.Println(c01.IsAPermutationOfAPalindrome("aabbCcddddEeFfgGgGGgggggX"))
	fmt.Println(c01.IsAPermutationOfAPalindrome("aabbCcddddEeFfgGgGGgggggXYYY"))

	fmt.Println("Testing OneAway")
	fmt.Println(c01.OneAway("asdf", "asdf"))
	fmt.Println(c01.OneAway("adf", "asdf"))
	fmt.Println(c01.OneAway("asddf", "asdf"))
	fmt.Println(c01.OneAway("addf", "asdf"))

	fmt.Println("Testing MatrixRotate")

	// Wow this is awful. Definitely should be using openCV for image creation and manipulation.

	image1 := [][]c01.Pixel {
		{ c01.IntPixel{Value: 1}, c01.IntPixel{Value: 2}, c01.IntPixel{Value: 3}, c01.IntPixel{Value: 4} },
		{ c01.IntPixel{Value: 5}, c01.IntPixel{Value: 6}, c01.IntPixel{Value: 7}, c01.IntPixel{Value: 8} },
		{ c01.IntPixel{Value: 9}, c01.IntPixel{Value: 10}, c01.IntPixel{Value: 11}, c01.IntPixel{Value: 12} },
		{ c01.IntPixel{Value: 13}, c01.IntPixel{Value: 14}, c01.IntPixel{Value: 15}, c01.IntPixel{Value: 16} },
	}

	newImage1 := c01.ImageRotate(image1)
	c01.InPlaceImageRotate(image1)
	fmt.Println(newImage1)

	image2 := [][]c01.Pixel {
		{c01.IntPixel{Value: 1}, c01.IntPixel{Value: 2}, c01.IntPixel{Value: 3}, c01.IntPixel{Value: 4}, c01.IntPixel{Value: 5}},
		{c01.IntPixel{Value: 6}, c01.IntPixel{Value: 7}, c01.IntPixel{Value: 8}, c01.IntPixel{Value: 9}, c01.IntPixel{Value: 10}},
		{c01.IntPixel{Value: 11}, c01.IntPixel{Value: 12}, c01.IntPixel{Value: 13}, c01.IntPixel{Value: 14}, c01.IntPixel{Value: 15}},
		{c01.IntPixel{Value: 16}, c01.IntPixel{Value: 17}, c01.IntPixel{Value: 18}, c01.IntPixel{Value: 19}, c01.IntPixel{Value: 20}},
		{c01.IntPixel{Value: 21}, c01.IntPixel{Value: 22}, c01.IntPixel{Value: 23}, c01.IntPixel{Value: 24}, c01.IntPixel{Value: 25}},
	}

	newImage2 := c01.ImageRotate(image2)
	c01.InPlaceImageRotate()
	fmt.Println(newImage2)
}

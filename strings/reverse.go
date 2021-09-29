package strings

import (
	"unicode/utf8"
)

func ReverseStrings(a *[]string) {
	// initialize start and end index pointer.
	i := 0
	j := len(*a) - 1

	for i < j {

		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
		// move the cursor toward the middle.
		i++
		j--
	}
}

func ReverseStrings2(a *[]string) {

	n := len(*a) - 1

	for i := 0; i < n/2; i++ {
		(*a)[n-i], (*a)[i] = (*a)[i], (*a)[n-i]
	}
}


func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Reverse1(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Reverse2(s string) string {
	runes := []rune(s)
	n := len(runes) - 1
	for i := 0; i < n/2; i++ {
		runes[n-i], runes[i] = runes[i], runes[n-i]
	}
	return string(runes)
}

func Reverse3(str string) string {
	var size int

	tail := len(str)
	buf := make([]byte, tail)
	s := buf

	for len(str) > 0 {
		_, size = utf8.DecodeRuneInString(str)
		tail -= size
		s = append(s[:tail], []byte(str[:size])...)
		str = str[size:]
	}

	return string(buf)
}
  

func ReverseSimple(value string) string {
	// Convert string to rune slice.
	// ... This method works on the level of runes, not bytes.
	data := []rune(value)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// Return new string.
	return string(result)
}
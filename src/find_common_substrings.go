package src

import (
	"fmt"
	"index/suffixarray"
	"regexp"
	"strings"
	// "github.com/flanglet/kanzi-go"
)

func FindCommonSubstrings(text1 string, text2 string, minLen int) {
	const TEXT_SEP string = "\x00"
	const PAGE_SEP string = "@@@"
	texts := []string{text1, text2}
	textCombined := "\x00" + strings.Join(texts, TEXT_SEP) + "\x00"
	sa := suffixarray.New([]byte(textCombined))
	fmt.Println(sa)
	match, err := regexp.Compile("\x00[^\x00]*")
	if err != nil {
		panic(err)
	}
	ms := sa.FindAllIndex(match, -1)

	for _, m := range ms {
		start, end := m[0], m[1]
		fmt.Printf("match = %q\n", textCombined[start+1:end])
	}
	// fmt.Println(textCombined)
}

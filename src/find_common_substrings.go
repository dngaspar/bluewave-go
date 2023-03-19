package src

import (
	"strings"

	lcppkg "github.com/nikitkocent/lcp"
)

type Result struct {
	subString string
	k1        int
	k2        int
	h         int
}

type LcpInfo struct {
	lcpSize   int
	subString string
	index1    int
	index2    int
}

var result []Result

func FindCommonSubstrings(text1 string, text2 string, minLen int) []Result {
	const TEXT_SEP string = "\x00"
	texts := []string{text1, text2}
	textCombined := strings.Join(texts, TEXT_SEP)
	lcp := lcppkg.NewLongestCommonPrefix(textCombined)
	text1Length := len(text1)
	text2Length := len(text2)
	var lcpInfo []LcpInfo
	for i := 0; i < text1Length; i++ {
		for j := 0; j < text2Length; j++ {
			item := int(lcp.Get(uint(i), uint(j+text1Length+1)))
			if item >= minLen {
				var temp LcpInfo
				temp.lcpSize = item
				temp.subString = textCombined[i : i+item]
				temp.index1 = i
				temp.index2 = j
				lcpInfo = append(lcpInfo, temp)
			}
		}
	}
	lenLcpSize := len(lcpInfo)
	for i := 1; i < lenLcpSize-1; i++ {
		prev := lcpInfo[i-1].lcpSize
		cur := lcpInfo[i].lcpSize
		next := lcpInfo[i+1].lcpSize
		var tempResult Result
		if i == 1 && prev > cur {
			tempResult.subString = lcpInfo[i-1].subString
			tempResult.h = lcpInfo[i-1].lcpSize
			tempResult.k1 = lcpInfo[i-1].index1
			tempResult.k2 = lcpInfo[i-1].index2
			result = append(result, tempResult)
		}
		if cur >= prev && cur >= next && !(prev == next && prev == cur) {
			tempResult.subString = lcpInfo[i].subString
			tempResult.h = lcpInfo[i].lcpSize
			tempResult.k1 = lcpInfo[i].index1
			tempResult.k2 = lcpInfo[i].index2
			result = append(result, tempResult)
		}
		if i == lenLcpSize-1 && cur < next {
			tempResult.subString = lcpInfo[i+1].subString
			tempResult.h = lcpInfo[i+1].lcpSize
			tempResult.k1 = lcpInfo[i+1].index1
			tempResult.k2 = lcpInfo[i+1].index2
			result = append(result, tempResult)
		}
	}

	return result
}

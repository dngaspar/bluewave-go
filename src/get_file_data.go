package src

import (
	"bluewave/pdf"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Block struct {
	Text           string    `json:"text"`
	Digits         string    `json:"digits"`
	Cum_len_text   int       `json:"cum_len_text"`
	Cum_len_digits int       `json:"cum_len_digits"`
	Bbox           []float32 `json:"bbox"`
	Page_num       int       `json:"page_num"`
	Block_num      int       `json:"block_num"`
}

type ImageHash struct {
}

type PdfInfo struct {
	Blocks      []Block     `json:blocks`
	Image_hashs []ImageHash `json:image_hashs`
	N_pages     int         `json:n_pages`
}

type Pdf struct {
	Version string  `json:"version"`
	Data    PdfInfo `json:"data"`
}

func GetMax(a []int) int {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func GetMin(a []int) int {
	min := a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}

func pageSkipConditions(pageText string) bool {
	conds := []string{
		"FORM FDA ",
		"Form FDA ",
		"PAPERWORK REDUCTION ACT",
		"PAYMENT IDENTIFICATION NUMBER",
		"For more assistance with Adobe Reader",
		"latest version of Adobe Reader",
		"..................................................................",
		"Safety Data Sheet",
		"SAFETY DATA SHEET",
		"Contains Nonbinding Recommendations",
	}

	for i := range conds {
		if strings.Contains(pageText, conds[i]) {
			return true
		}
	}
	return false
}

func blockSkipConditions(blockText string) bool {
	conds := []string{
		"510(k)",
		"New Hampshire Avenue",
		"ISO ",
		"IEC ",
		"..............",
		"Tel.:",
		"TEL:",
		"FAX:",
		"Fax:",
		"+86",
		"86-519",
	}
	for i := range conds {
		if strings.Contains(blockText, conds[i]) {
			return true
		}
	}
	return false
}

func isCompatible(vCurrent string, vCache string) bool {
	// vCurrent and vCache are strings like 1.1.1
	// compare first two digits
	strCur := strings.Join(strings.Split(vCurrent, ".")[:2], "")
	strCache := strings.Join(strings.Split(vCache, ".")[:2], "")
	intCur, _ := strconv.Atoi(strCur)
	intCache, _ := strconv.Atoi(strCache)
	// intCache := strconv.Atoi(strCur)
	return intCur == intCache

}

var white uint32 = 55535

func checkWhite(page image.Image, x int, y int) bool {
	r, g, b, _ := page.At(x, y).RGBA()
	if r >= white && g >= white && b >= white {
		return true
	} else {
		return false
	}
}

func getPageBlockAndHashes(fileName string, pageNum int) {
	// """Get page block and hashes"""
	// var imageHashes []ImageHash
	// var textBlocks []Block
	var pageImage image.Image

	pageImage = pdf.GetPageInfo(fileName, pageNum)
	// min := pageImage.Bounds().Min
	max := pageImage.Bounds().Max
	// fmt.Println(max)
	// r, g, b, a := pageImage.At(3, 4).RGBA()
	// fmt.Println(max, r, g, b, a)
	m := max.X
	n := max.Y
	// var blockarr [22222][22222]int
	type Data struct {
		color int
		x     int
		y     int
		st    int
		en    int
	}

	type Info struct {
		x int
		y int
		w int
		h int
	}
	var Array []Data

	for i := 0; i < n; i++ {
		var ok int
		minid := m
		maxid := -1
		for j := 0; j < m; j++ {
			// r, g, b, a := pageImage.At(i, j).RGBA()
			if !checkWhite(pageImage, j, i) {
				// fmt.Println(r, g, b, a)
				ok = 1
				minid = GetMin([]int{minid, j})
				maxid = GetMax([]int{maxid, j})
			}
		}
		var temp Data
		if ok == 1 {
			temp.color = 1
			temp.x = minid
			temp.y = i
			temp.st = minid
			temp.en = maxid
			Array = append(Array, temp)
		} else {
			temp.color = 0
			temp.x = 0
			temp.y = i
			temp.st = 0
			temp.en = m - 1
			Array = append(Array, temp)
		}
	}

	curid := -1
	curwidth := 0

	var info []Info
	for i := 0; i < n; i++ {
		if i == 0 {
			if Array[i].color == 1 {
				curid = 0
				curwidth = GetMax([]int{curwidth, Array[i].en - Array[i].st})
			}
			continue
		}
		if curid == -1 && Array[i].color == 1 {
			curid = i
			curwidth = 0
			curwidth = GetMax([]int{curwidth, Array[i].en - Array[i].st})
		} else if Array[i].color == 0 && Array[i-1].color == 1 {
			var temp Info
			temp.x = Array[curid].x
			temp.y = Array[curid].y
			temp.w = curwidth
			temp.h = i - curid
			info = append(info, temp)
			curid = -1
			curwidth = 0
			curwidth = GetMax([]int{curwidth, Array[i].en - Array[i].st})
		} else if Array[i].color == 1 {
			curwidth = GetMax([]int{curwidth, Array[i].en - Array[i].st})
		} else {
			curid = -1
			curwidth = 0
		}
	}

	if curid != -1 {
		var temp Info
		temp.x = Array[curid].x
		temp.y = Array[curid].y
		temp.w = curwidth
		temp.h = n - curid
		info = append(info, temp)
	}

}

func readBlocksAndHeaders(fileName string) {
	if fileName[len(fileName)-4:] != ".pdf" {
		panic("Fitz cannot read non-PDF file")
	}
	getPageBlockAndHashes(fileName, 0)
}

func GetFileData(fullFileName string, fileIndex int, regenCache bool, version string) {
	var cachedFileName string = fullFileName + ".jsoncached"
	_, err := os.Stat(cachedFileName)
	if err == nil && !regenCache {
		cachedFile, err := os.Open(cachedFileName)
		if err != nil {
			fmt.Println(cachedFile)
		}
		defer cachedFile.Close()
		byteValue, _ := ioutil.ReadAll(cachedFile)

		// var cached Cache
		var cached Pdf
		json.Unmarshal(byteValue, &cached)
		// fmt.Println(cached)
		var blocks []Block
		var imageHashes []ImageHash
		var nPages int
		lenBlocks := 0
		lenImageHashes := 0
		if isCompatible(version, cached.Version) {
			blocks = cached.Data.Blocks
			lenBlocks = len(blocks)
			imageHashes = cached.Data.Image_hashs
			nPages = cached.Data.N_pages
			lenBlocks = len(imageHashes)
			fmt.Println(blocks, imageHashes, nPages)
		}

		// if len(blocks) == 0 && len(image_hashes) && n_pages == 0 {

		// }
		if lenBlocks == 0 && lenImageHashes == 0 && nPages == 0 {
		}
	}
	readBlocksAndHeaders(fullFileName)

}

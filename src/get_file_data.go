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

func GetMinMax(a []int) int {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
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

func getPageBlockAndHashes(fileName string, pageNum int) {
	// """Get page block and hashes"""
	// var imageHashes []ImageHash
	// var textBlocks []Block
	var pageImage image.Image
	pageImage = pdf.GetPageInfo(fileName, pageNum)
	// min := pageImage.Bounds().Min
	max := pageImage.Bounds().Max
	fmt.Println(max)
	r, g, b, a := pageImage.At(3, 4).RGBA()
	fmt.Println(max, r, g, b, a)
	maxX := max.X
	maxY := max.Y
	var blockarr [][]int
	var white uint32 = 65535
	var blockNo int
	var xarr []int
	for y := 0; y < maxY; y++ {
		var ok int = 0
		// var firstline int = 0
		var len int = 0
		for x := 0; x < maxX; x++ {
			r, g, b, a := pageImage.At(x, y).RGBA()
			if r != white || g != white || b != white || a != white {
				ok = 1
				// if firstline == 0 {
				// 	firstline = 1
				// }
				xarr = append(xarr, x)
				len = x
			}
		}
		if ok == 0 {
			// if firstline == 1 {
			for i := 0; i < 4; i++ {
				blockarr[blockNo][i] = 0
			}
			blockarr[blockNo][1] = y
			if blockNo > 0 {
				blockarr[blockNo-1][3] = y - blockarr[blockNo-1][1]
			}
			// firstline = 0
			blockarr[blockNo][0] = GetMinMax(xarr)
			blockarr[blockNo][2] = len
			xarr = nil
			blockNo++
			// }
		}
	}
	// x := pageImage.At(2549, 3299)
	fmt.Println(blockarr)

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

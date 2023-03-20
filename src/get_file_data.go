package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type CacheDataMain struct {
	Text           string    `json:"text"`
	Digits         string    `json:"digits"`
	Cum_len_text   int       `json:"cum_len_text"`
	Cum_len_digits int       `json:"cum_len_digits"`
	Bbox           []float32 `json:"bbox"`
	Page_num       int       `json:"page_num"`
	Block_num      int       `json:"block_num"`
}

type CacheData struct {
	Blocks      []CacheDataMain `json:blocks`
	Image_hashs []string        `json:image_hashs`
	N_pages     int             `json:n_pages`
}

type Cache struct {
	Version string    `json:"version"`
	Data    CacheData `json:"data"`
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

func readBlocksAndHeaders(fileName string) {
	if fileName[len(fileName)-4:] != ".pdf" {

	}
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
		var cached Cache
		json.Unmarshal(byteValue, &cached)
		// fmt.Println(cached)
		var blocks []CacheDataMain
		var imageHashes []string
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
			readBlocksAndHeaders(fullFileName)
		}
	}

}

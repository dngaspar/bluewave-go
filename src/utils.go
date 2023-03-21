package src

import (
	"math"
	"unicode"
)

type Box struct {
	xmin   float64
	ymin   float64
	xmax   float64
	ymax   float64
	width  float64
	height float64
	area   float64
}

var box Box

func Intersection(box1 Box, box2 Box) (Box, error) {
	var temp Box
	temp.xmin = math.Min(box1.xmin, box2.xmin)
	temp.xmax = math.Max(box1.xmax, box2.xmax)
	temp.ymin = math.Min(box1.ymin, box2.ymin)
	temp.ymax = math.Max(box1.ymin, box2.ymin)
	width := temp.xmax - temp.xmin
	height := temp.ymax - temp.ymin
	if width > 0 && height > 0 {
		return temp, nil
	}
	var err error
	return temp, err
}

func Expand(box1 Box, box2 Box) Box {
	var temp Box
	temp.xmin = math.Min(box1.xmin, box2.xmin)
	temp.xmax = math.Max(box1.xmax, box2.xmax)
	temp.ymin = math.Min(box1.ymin, box2.ymin)
	temp.ymax = math.Max(box1.ymin, box2.ymin)
	return temp
}

func BoxDistance(box1 Box, box2 Box) float64 {
	newArea := Expand(box1, box2).area
	var unionArea float64
	intersection, err := Intersection(box1, box2)
	if err == nil {
		unionArea = box1.area + box2.area - intersection.area
	}

	d := (newArea / unionArea) - 1
	return math.Max(0, d)

}

// func asTuple(box1 Box) (float64, float64) {
// 	return
// }

func getModelsDirectory() string {
	currentDirectory := "models"
	return currentDirectory
}

// func IsNumeric(word string) bool {
// 	return regexp.MustCompile(`\d`).MatchString(word)
// }

func getDigits(text string) string {
	digits := ""
	nLetters := 0

	for _, char := range text {
		if unicode.IsDigit(char) {
			digits += string(char)
			nLetters++
		} else if unicode.IsLetter(char) {
			nLetters++
		}
	}

	if len(text) > 10 && float64(len(digits))/float64(nLetters) < 0.1 {
		return ""
	}

	return digits
}

// func logit(p float64) {
// 	return np.log(p / (1 - p))
// }

// func logistic(x float64) {
// 	return 1 / (np.exp(-x) + 1)
// }

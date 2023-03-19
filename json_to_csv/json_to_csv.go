package json_to_csv

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/exports"
)

func JsonToCsv(fileName string) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	type Pair struct {
		Type  string `json:"type"`
		Text  string `json:"string"`
		IsBad bool   `json:"is_bad"`
	}

	type PairsuspiciousPairs struct {
		Pairs []Pair `json:"suspicious_pairs"`
	}

	var pairsuspiciousPairs PairsuspiciousPairs

	json.Unmarshal(byteValue, &pairsuspiciousPairs)
	textDf := dataframe.NewDataFrame(dataframe.NewSeriesString("sp", nil), dataframe.NewSeriesString("text", nil), dataframe.NewSeriesString("is_bad", nil))
	digitsDf := dataframe.NewDataFrame(dataframe.NewSeriesString("sp", nil), dataframe.NewSeriesString("text", nil), dataframe.NewSeriesString("is_bad", nil))

	for i := 0; i < len(pairsuspiciousPairs.Pairs); i++ {
		out, err := json.Marshal(pairsuspiciousPairs.Pairs[i])
		sp := string(out)
		if err != nil {
			panic(err)
		}
		fmt.Println(pairsuspiciousPairs.Pairs[i].Text)
		fmt.Println(pairsuspiciousPairs.Pairs[i].Type == "Common text string")
		if pairsuspiciousPairs.Pairs[i].Type == "Common text string" {
			textDf.Append(nil, map[string]interface{}{
				"sp":     sp,
				"text":   pairsuspiciousPairs.Pairs[i].Text,
				"is_bad": nil,
			})
		}
		if pairsuspiciousPairs.Pairs[i].Type == "Common digit sequence" {
			digitsDf.Append(nil, map[string]interface{}{
				"sp":     sp,
				"text":   pairsuspiciousPairs.Pairs[i].Text,
				"is_bad": nil,
			})
		}
	}

	textFile, err := os.Create("data/text_training_data.csv")
	if err != nil {
		panic(err)
	}

	digitFile, err := os.Create("data/digits_training_data.csv")
	if err != nil {
		panic(err)
	}

	exports.ExportToCSV(context.Background(), textFile, textDf)
	exports.ExportToCSV(context.Background(), digitFile, digitsDf)

}

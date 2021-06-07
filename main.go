package main

import (
 	"fmt"
	"strings"
	common "ClinicalIndication/common"
	valid "github.com/asaskevich/govalidator"
	"github.com/tealeg/xlsx"
	"github.com/TwinProduction/go-color"
	"strconv"
)

var file *xlsx.File
var cell *xlsx.Cell
var sh 	*xlsx.Sheet
var fieldsToExtract = []string{"Cod", "Group", "Subgroup", "CID"}

func main() {
	println(color.Blue + "Start" + color.Reset)
	indications := common.ReadCsv()

	words := map[string]int{}
	wordscod := make(map[string]map[string]int)
	wordscid := make(map[string]map[string]int)
	wordsgrp := make(map[string]map[string]int)
	wordssub := make(map[string]map[string]int)

	for idx, client := range indications {
		f := common.FormatText(client.Indication)
		
		indications[idx].Formated = f
		res, _ := common.ExtractCid(f)
		// if err != nil  {
		// 	fmt.Printf(" Error: %s, %v;", res, err)
		// }
		indications[idx].CID = res

		for _, word := range strings.Fields(f) {
			if len(word) == 1 || valid.IsInt(word) || word[0] == '0' {
				continue
			}

			words[word] += 1

			if _, ok := wordscod[indications[idx].Cod]; !ok {
				wordscod[indications[idx].Cod] = make(map[string]int)
			}
			wordscod[indications[idx].Cod][word] += 1
			
			if _, ok := wordsgrp[indications[idx].Group]; !ok {
				wordsgrp[indications[idx].Group] = make(map[string]int)
			}
			wordsgrp[indications[idx].Group][word] += 1

			if _, ok := wordssub[indications[idx].Subgroup]; !ok {
				wordssub[indications[idx].Subgroup] = make(map[string]int)
			}
			wordssub[indications[idx].Subgroup][word] += 1

			for _, cid := range indications[idx].CID {
				if _, ok := wordscid[cid]; !ok {
					wordscid[cid] = make(map[string]int)
				}
				wordscid[cid][word] += 1
			}
		}

	}

	//fmt.Println(wordscod[30101069])
	//fmt.Println(wordscid["s611"])
	println(color.Yellow + "CIDs: " + strconv.Itoa(len(wordscid)) + color.Reset)

	file = xlsx.NewFile()

	//chanel
	common.WriteGeral(file, cell, words, "Geral")
	common.WriteRange(file, sh, cell, wordscid, "CID")
	common.WriteRange(file, sh, cell, wordscod, "COD")
	common.WriteRange(file, sh, cell, wordsgrp, "GRU")
	common.WriteRange(file, sh, cell, wordssub, "SUB")

	err := file.Save("CI.xls")
	if err != nil {
		fmt.Printf(err.Error())
	}
	println(color.Green + "End" + color.Reset)
}

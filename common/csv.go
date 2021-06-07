package common

import (
	"os"
	"github.com/gocarina/gocsv"
)

type Indication struct { 
	Cod      	string 	`csv:"cod_item"`
	Procedure	string	`csv:"descricao"`
	Group		string  `csv:"grupo"`
	Subgroup	string  `csv:"subgrupo"`
	Indication 	string	`csv:"ds_indicacao_clinica"`
	CID			[]string	`csv:"-"`
	Formated	string	`csv:"-"`
}

func ReadCsv() []*Indication {

	csvFile, err := os.OpenFile("IND.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	indications := []*Indication{}

	if err := gocsv.UnmarshalFile(csvFile, &indications); err != nil { 
		panic(err)
	}

	if _, err := csvFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	return indications
}
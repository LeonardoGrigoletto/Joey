package main

import (
	"fmt"
	"joey/joey"
)

func main() {
	dataframe, _ := joey.NewFromCsv("/home/leogrig/Documentos/repositorios/Curso-GOLANG/Joey/src/test.csv")
	dataframe.Show(12)
	dataframe.ShowTypes()

	dataframe.Convert("charge", "int")
	dataframe.ShowTypes()

	column, _ := dataframe.Column("charge")
	fmt.Println(column.Sum())

	index := column.FindFirst(int64(10))
	fmt.Println(index)

	// dataframe.Sum("column2", "column4")

}

// Em RemoveCol, adicionar maneiras de remover os ponteiros das cells removidas. Ponteiros de column para rows

package main

import "joey/joey"

func main() {
	dataframe := joey.NewFromCsv("/home/leogrig/Documentos/repositorios/Curso-GOLANG/Joey/src/test.csv")
	dataframe.Show(5)
	dataframe.ShowTypes()
	// dataframe.RemoveCol("column5")

	// dataframe.Sum("column2", "column4")

}

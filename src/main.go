package main

import (
	"joey/joey"
)

func main() {
	dataframe, _ := joey.NewFromCsv("/home/leogrig/Documentos/repositorios/Curso-GOLANG/Joey/src/test.csv")
	dataframe.Show()
	dataframe.ShowTypes()

}

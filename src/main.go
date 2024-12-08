package main

import (
	"joey/joey"
)

type Numbers interface {
	joey.StrCell | joey.IntCell
}

func main() {
	// dataframe, _ := joey.NewFromCsv("/home/leogrig/Documentos/repositorios/Curso-GOLANG/Joey/src/test.csv")
	// dataframe.Show()
	// dataframe.ShowTypes()
}

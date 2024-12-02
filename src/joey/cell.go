package joey

type Cell interface {
	// This function returns cell value in string format.
	GetFormattedData() string
	// This function returns cell value in its real type.
	GetRawData() interface{}
	// This functions returns cell value converted to float64 to use in arithmetic operations.
	GetNumber() float64
	// This function returns the length of its value when converted to string
	Length() int
	// This functions convert the Cell to another type, eg: StrCell, Int64Cell ...
	Convert(to string) (Cell, error)
}

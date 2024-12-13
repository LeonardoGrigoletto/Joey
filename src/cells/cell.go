package cells

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
	// This function returns the native type of a cell
	GetNativeType() interface{}
	// This function returns the artificial type of a cell, eg: StrCell, IntCell, Float32Cell...
	GetType() interface{}
	// This function adds a value to this cell. It will panic if the types are mismatched.
	Add(cell Cell)
	// This function subtract a value from this cell. It will panic if the types are mismatched.
	Subtract(cell Cell)
	// This function multiply a value with this cell. It will panic if the types are mismatched.
	Multiply(cell Cell)
}

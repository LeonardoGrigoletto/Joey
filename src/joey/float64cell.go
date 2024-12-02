package joey

import (
	"errors"
	"strconv"
	"strings"
)

type Float64Cell struct {
	data float64
}

func (f Float64Cell) GetFormattedData() string {
	return strconv.FormatFloat(f.data, 'f', 2, 64)
}

func (f Float64Cell) Length() int {
	return len(strconv.FormatFloat(f.data, 'f', 2, 64))
}

func (f Float64Cell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatFloat(f.data, 'f', 2, 64)
		return StrCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int32") {
		convertedData := int32(f.data)
		return Int32Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int64") {
		convertedData := int64(f.data)
		return Int64Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "float64") {
		return f, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (f Float64Cell) GetRawData() any {
	return f.data
}

func (f Float64Cell) GetNumber() float64 {
	return f.data
}

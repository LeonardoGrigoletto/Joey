package cells

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type StrCell struct {
	Data string
}

func (s *StrCell) Add(cell Cell) {
	data := cell.GetRawData()
	value, ok := data.(string)
	if !ok {
		panic("It is not possible to sum columns of different types")
	}
	s.Data += value
}

func (s StrCell) GetType() interface{} {
	return reflect.TypeOf(StrCell{})
}

func (s StrCell) GetNativeType() interface{} {
	return reflect.TypeOf(s.Data)
}

func (s StrCell) GetFormattedData() string {
	return s.Data
}

func (s StrCell) Length() int {
	return len(s.Data)
}

func (s StrCell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		return &s, nil
	}
	if strings.EqualFold(to, "int") {
		convertedData, err := strconv.ParseInt(s.Data, 10, 64)
		if err != nil {
			return nil, err
		}
		return &IntCell{Data: int(convertedData)}, nil
	}
	if strings.EqualFold(to, "int8") {
		convertedData, err := strconv.ParseInt(s.Data, 10, 64)
		if err != nil {
			return nil, err
		}
		return &Int8Cell{Data: int8(convertedData)}, nil
	}
	if strings.EqualFold(to, "int16") {
		convertedData, err := strconv.ParseInt(s.Data, 10, 64)
		if err != nil {
			return nil, err
		}
		return &Int16Cell{Data: int16(convertedData)}, nil
	}
	if strings.EqualFold(to, "int32") {
		convertedData, err := strconv.ParseInt(s.Data, 10, 64)
		if err != nil {
			return nil, err
		}
		return &Int32Cell{Data: int32(convertedData)}, nil
	}
	if strings.EqualFold(to, "int64") {
		convertedData, err := strconv.ParseInt(s.Data, 10, 64)
		if err != nil {
			return nil, err
		}
		return &Int64Cell{Data: int64(convertedData)}, nil
	}
	if strings.EqualFold(to, "float32") {
		convertedData, err := strconv.ParseFloat(s.Data, 64)
		if err != nil {
			return nil, err
		}
		return &Float32Cell{Data: float32(convertedData)}, nil
	}
	if strings.EqualFold(to, "float64") {
		convertedData, err := strconv.ParseFloat(s.Data, 64)
		if err != nil {
			return nil, err
		}
		return &Float64Cell{Data: convertedData}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (s StrCell) GetRawData() any {
	return s.Data
}

func (s StrCell) GetNumber() float64 {
	return float64(0)
}

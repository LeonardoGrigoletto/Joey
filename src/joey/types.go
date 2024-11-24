package joey

import (
	"errors"
	"strconv"
	"strings"
)

type Row struct {
	data []Cell
}

type Column struct {
	data *[]Cell
}

type Cell interface {
	GetData() any
	Length() int
	Convert(to string) (Cell, error)
}

type FloatCell struct {
	data float64
}

func (f FloatCell) GetData() any {
	return f.data
}

func (f FloatCell) Length() int {
	return len(strconv.FormatFloat(f.data, 'f', 2, 64))
}

func (f FloatCell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatFloat(f.data, 'f', 2, 64)
		return StrCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int") {
		convertedData := int64(f.data)
		return IntCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "float64") {
		return f, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

type IntCell struct {
	data int64
}

func (i IntCell) GetData() any {
	return i.data
}

func (i IntCell) Length() int {
	return len(strconv.Itoa(int(i.data)))
}

func (i IntCell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatInt(int64(i.data), 10)
		return StrCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int") {
		return i, nil
	}
	if strings.EqualFold(to, "float64") {
		return FloatCell{data: float64(i.data)}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

type StrCell struct {
	data string
}

func (s StrCell) GetData() any {
	return s.data
}

func (s StrCell) Length() int {
	return len(s.data)
}

func (s StrCell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		return s, nil
	}
	if strings.EqualFold(to, "int") {
		convertedData, err := strconv.ParseInt(s.data, 10, 64)
		if err != nil {
			return nil, err
		}
		return IntCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "float64") {
		convertedData, err := strconv.ParseFloat(s.data, 64)
		if err != nil {
			return nil, err
		}
		return FloatCell{data: convertedData}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

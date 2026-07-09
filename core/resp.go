package core

import (
	"errors"
)

func readSimpleString(data []byte) (interface{}, int, error) {
	pos := 1
	for ; data[pos] != '\r'; pos++ {
	}
	return string(data[1:pos]), pos + 2, nil
}

func readError(data []byte) (interface{}, int, error) {
	return readSimpleString(data)
}

func readInt64(data []byte) (interface{}, int, error) {
	pos := 1
	var value int64
	for ; data[pos] != '\r'; pos++ {
		value = value*10 + int64(data[pos]-'0')
	}
	return value, pos + 2, nil
}

func readBulkString(data []byte) (interface{}, int, error) {
	pos := 1
	len, delta := readLength(data[pos:])
	pos += delta+2
	return string(data[pos:(pos + len)]), pos + len + 2, nil
}

func readLength(data []byte) (int, int) {
	pos := 0
	var value int
	for ; data[pos] != '\r'; pos++ {
		value = value*10 + int(data[pos] - '0')
	}
	return value, pos
}

func readArray(data []byte) (interface{}, int, error) {
	pos := 1

	count, delta := readLength(data[pos:])
	pos += delta+2

	var elems []interface{} = make([]interface{}, count)

	for i := range elems {
		elem, delta, err := DecodeOne(data[pos:])
		if err != nil {
			return nil, 0, err
		}
		elems[i] = elem
		pos += delta
	}
	return elems, pos, nil
}

func DecodeOne(data []byte) (interface{}, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("no data")
	}
	switch data[0] {
	case '+':
		return readSimpleString(data)
	case '-':
		return readError(data)
	case ':':
		return readInt64(data)
	case '$':
		return readBulkString(data)
	case '*':
		return readArray(data)
	}
	return nil, 0, nil

}
func Decode(data []byte) (interface{}, error) {
	if len(data) == 0 {
		return nil, errors.New("no data")
	}
	value, _, err := DecodeOne(data)
	return value, err
}

package core

import (
	"testing"
	"fmt"
)

func TestReadSimpleString(t *testing.T) {
	testCases := map[string]string{
		"+OK\r\n": "OK",
	}

	for k, v := range testCases {
		value, _ := Decode([]byte(k))
		if value != v {
			t.Fail()
		}
	}
}

func TestReadError(t *testing.T) {
	testCases := map[string]string{
		"-Error Message\r\n": "Error Message",
	}

	for k, v := range testCases {
		value, _ := Decode([]byte(k))
		if value != v {
			t.Fail()
		}
	}
}

func TestReadInt64(t *testing.T) {
	testCases := map[string]int64{
		":0\r\n":    0,
		":1000\r\n": 1000,
	}
	for k, v := range testCases {
		value, _ := Decode([]byte(k))
		if value != v {
			t.Fail()
		}
	}
}

func TestReadBulkString(t *testing.T) {
	testCases := map[string]string{
		"$5\r\nhello\r\n": "hello",
		"$1\r\nA\r\n":     "A",
	}
	for k, v := range testCases {
		value, _ := Decode([]byte(k))
		if value != v {
			fmt.Println("value", value)
			fmt.Println("v", v)
			t.Fail()
		}
	}
}

func TestReadArray(t *testing.T) {
	testCases := map[string][]interface{}{
		"*0\r\n":                               {},
		"*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n": {"hello", "world"},
		"*3\r\n:1\r\n:2\r\n:3\r\n":             {int64(1), int64(2), int64(3)},
		"*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Hello\r\n-World\r\n": {[]int64{int64(1), int64(2), int64(3)}, []interface{}{"Hello", "World"}},
	}
	for k, v := range testCases {
		value, _ := Decode([]byte(k))
		array := value.([]interface{})
		if len(array) != len(v) {
			t.Fail()
		}
		for i := range array {
			if fmt.Sprintf("%v", v[i]) != fmt.Sprintf("%v", array[i]) {
				t.Fail()
			}
		}
	}
}

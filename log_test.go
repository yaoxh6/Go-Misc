package log

import (
	"fmt"
	"testing"
)

type EmptyStruct struct {

}
type TestStruct struct {
	Id int
	Name string
	Map map[string][]float64
	EmptyStruct EmptyStruct
	Address *string
	Reserver interface{}
}


func print() {
	fmt.Print("hello world")
}

func TestLog(t *testing.T) {
	tempString := "tempString"
	tempData := TestStruct{
		Id:       12,
		Name:     "tempData",
		Map: 	map[string][]float64{
			"test1":[]float64{12.3, 33.0},
			"test2":[]float64{88.456, 78.12},
		},
		EmptyStruct : EmptyStruct{},
		Address:  &tempString,
		Reserver: map[interface{}]interface{}{
			"Reserver": &tempString,
			"Func": print,
			"Nil": nil,
		},
	}
	LogTree("tempData", tempData)
}

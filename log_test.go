package log

import (
	"fmt"
	"testing"
)

type EmptyStruct struct {

}

type StringStruct struct {
	Str string
}
type TestStruct struct {
	Id int
	Name string
	FloatMap map[string][]float64
	StringMap map[int]StringStruct
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
		FloatMap: 	map[string][]float64{
			"test1":{12.3, 33.0},
			"test2":{88.456, 78.12},
		},
		StringMap: map[int]StringStruct{
			1:{Str: "hello"},
			2:{Str: "world"} ,
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

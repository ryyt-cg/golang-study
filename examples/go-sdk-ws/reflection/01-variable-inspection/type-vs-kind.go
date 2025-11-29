package main

import (
	"fmt"
	"reflect"
)

type MyString string

type MyStruct struct {
	Field1 int
	Field2 string
}

func main() {
	var s MyString = "hello"
	var i int = 10
	var i32 int32 = 20
	var ms MyStruct

	fmt.Printf("s: Type=%v, Kind=%v\n", reflect.TypeOf(s), reflect.ValueOf(s).Kind())
	fmt.Printf("i: Type=%v, Kind=%v\n", reflect.TypeOf(i), reflect.ValueOf(i).Kind())
	fmt.Printf("i32: Type=%v, Kind=%v\n", reflect.TypeOf(i32), reflect.ValueOf(i32).Kind())
	fmt.Printf("ms: Type=%v, Kind=%v\n", reflect.TypeOf(ms), reflect.ValueOf(ms).Kind())
}

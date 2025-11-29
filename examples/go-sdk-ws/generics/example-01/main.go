package main

import "fmt"

func print[T int | string | bool | person](data T) {
	fmt.Println(data)
}

type person struct {
	Name string
	Age  int
}

func main() {
	print("hello")
	print(123)
	print(true)
	print(person{
		Name: "Jay",
		Age:  23,
	})
}

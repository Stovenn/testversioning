package main

import (
	"github.com/stovenn/testversioning/pkg/mypkg"
)

func main() {
	s := mypkg.NewMyStruct("Hello", "World")
	MyFunc(s)
}

func MyFunc(data mypkg.MyStruct) bool {
	return data.CheckIfHelloWorld()
}

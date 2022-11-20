package main

import (
	"github.com/stovenn/testversioning/pkg/mypkg"
	"strings"
)

func main() {
	MyFunc(mypkg.MyStructV1{Foo: "Hello", Bar: "World"})
}

func MyFunc(data mypkg.MyStructV1) bool {
	return strings.Contains(data.Foo, "Hello") && strings.Contains(data.Bar, "World")
}

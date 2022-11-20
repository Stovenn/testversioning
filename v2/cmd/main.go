package main

import (
	"github.com/stovenn/testversioning/v2/pkg/mypkg"
	"log"
)

func main() {
	data := mypkg.MyStruct{}
	data.CheckIfHelloWorld()
	log.Println("testversioning")
}

package mypkg

type MyStruct struct {
	foo string
	bar string
}

func NewMyStruct(foo, bar string) MyStruct {
	return MyStruct{foo: foo, bar: bar}
}
func (ms MyStruct) CheckIfHelloWorld() bool {
	return ms.foo == "Hello" && ms.bar == "World"
}

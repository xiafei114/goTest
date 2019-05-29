package main

import (
	"fmt"

	"gotest/reflect/tools"
)

//你的结构体定义
type MyStruct struct {
	Name      string
	MaxHeight int
}

var s MyStruct

func main() {
	fmt.Println(tools.ProduceStructTag(s, "json"))
}

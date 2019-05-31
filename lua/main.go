package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func Double(L *lua.LState) int {
	lv := L.ToInt(1)            /* get argument */
	L.Push(lua.LNumber(lv * 2)) /* push result */
	return 1                    /* number of results */
}

func SayHello(L *lua.LState) int {
	name := L.ToString(1)
	fmt.Println("hello ", name)
	return 0
}

func GetDate(L *lua.LState) int {
	name := L.ToString(1)
	fmt.Println("date ", name)
	return 0
}

func main() {
	L := lua.NewState()
	defer L.Close()

	// L.SetGlobal("double", L.NewFunction(Double))
	// L.SetGlobal("sayHello", L.NewFunction(SayHello))

	L.SetGlobal("getDate", L.NewFunction(GetDate))

	L.DoString(`
	local date=os.date("%Y-%m-%d %H:%M:%S");
	getDate(date)
	`) // 8
	// L.DoString(`sayHello("wangjunsheng")`) // hello wangjunsheng

	//	L.DoFile("hello.lua") // from lua file
}

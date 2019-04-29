package main

import (
	// "encoding/json"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
	Age       int    `json:"age,omitempty"`
}

func main() {
	msg := Message{Sender: "test", Recipient: "aa", Age: 0}
	jsonStr, _ := jsoniter.Marshal(msg)
	fmt.Printf("%s", jsonStr)

	print()

	js := make(map[string]interface{})
	arrList := make([]interface{}, 0)

	arr1 := make([]interface{}, 0)
	arr1 = append(arr1, "data1")
	arr1 = append(arr1, 14)

	arr2 := make([]interface{}, 0)
	arr2 = append(arr2, "data2")
	arr2 = append(arr2, "red")

	arrList = append(arrList, arr1)
	arrList = append(arrList, arr2)

	userJSON := make(map[string]interface{})
	userJSON["a"] = "av1"
	userJSON["b"] = "bv1"

	roleJSONList := make([]interface{}, 0)
	roleJSON1 := make(map[string]interface{})
	roleJSON1["rid"] = "id1"
	roleJSON1["rname"] = "rName1"

	roleJSON2 := make(map[string]interface{})
	roleJSON2["rid"] = "id2"
	roleJSON2["rname"] = "rName2"
	roleJSONList = append(roleJSONList, roleJSON1)
	roleJSONList = append(roleJSONList, roleJSON2)

	js["params"] = arrList
	js["roles"] = roleJSONList

	js["int"] = 1
	js["string"] = "string"
	js["user"] = userJSON

	bt, _ := jsoniter.Marshal(js)

	fmt.Println(string(bt))

	print()

}

func print() {
	fmt.Println("---------------- 华丽的分割线  ---------------------")
}

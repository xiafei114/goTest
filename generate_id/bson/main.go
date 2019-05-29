package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("string  ID: %s\n", bson.NewObjectId().Hex())
	}
}

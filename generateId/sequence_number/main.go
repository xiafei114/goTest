package main

import (
	"fmt"
	"gotest/generateId/sequence_number/sequence"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "bindo:maxinfo@tcp(127.0.0.1:3306)/bindo?parseTime=true")
	if err != nil {
		fmt.Println(fmt.Sprintln(err))
	}

	defer db.Close()

	sequence.InitTable(db)

	arrs := [5]string{"test-1", "test-2", "test-3", "test-4", "test-5"}

	var mutex sync.Mutex

	for _, v := range arrs {
		fmt.Printf("start %s\n", v)
		go func(v string) {
			for n := 0; n < 100; n++ {
				doSeqNoGenerator(db, v, &mutex)
			}
			time.Sleep(2 * time.Second)
			fmt.Printf("end %s\n", v)
		}(v)

	}

	time.Sleep(2 * time.Second)
	fmt.Println("end")
}

func doSeqNoGenerator(db *gorm.DB, logicID string, mutex *sync.Mutex) {
	mutex.Lock()
	num, err := sequence.NewSeqNoGenerator(db, logicID).Next()
	if err != nil {
		fmt.Println(fmt.Sprintln(err))
	}
	fmt.Printf("%s --- %s\n", logicID, num)
	mutex.Unlock()
}

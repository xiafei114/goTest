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

	// seqno := sequence.NewSeqNoGenerator(db, "test-1")

	// seq, _ := seqno.Next()
	// fmt.Println(seq)
	arrs := [5]string{"test-1", "test-2", "test-3", "test-4", "test-5"}

	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	mutex.Lock()

	for _, v := range arrs {
		fmt.Printf("start %s\n", v)
		go func(v string) {
			for n := 0; n < 100; n++ {
				wait.Add(1)
				mutex.Lock()
				num, err := sequence.NewSeqNoGenerator(db, v).Next()
				if err != nil {
					fmt.Println(fmt.Sprintln(err))
				}
				fmt.Printf("%s --- %s\n", v, num)
				mutex.Unlock()

				defer wait.Done()
			}
			time.Sleep(2 * time.Second)
			fmt.Printf("end %s\n", v)
		}(v)

	}

	time.Sleep(2 * time.Second)
	fmt.Println("end")

	fmt.Println("Unlocked")
	mutex.Unlock()

	wait.Wait()

	// fmt.Println(fmt.Sprintln(db.Debug().Exec(sequence.MigrateSQL()).Error))

	// seqno := seqno.NewSeqNoGenerator(db, "test-logic-id").PrefixFormat("2006-01-02").Locker(seqno.NewRedisLocker(
	// 	&redis.Pool{
	// 		MaxIdle:     5,
	// 		IdleTimeout: 30 * time.Second,
	// 		Dial: func() (redis.Conn, error) {
	// 			return redis.Dial("tcp", "localhost:6379")
	// 		},
	// 	},
	// )).Step(1)

	// for n := 0; n < 100; n++ {
	// 	num, err := seqno.Next()
	// 	if err != nil {
	// 		fmt.Println(fmt.Sprintln(err))
	// 	}
	// 	fmt.Println(num)
	// }
}

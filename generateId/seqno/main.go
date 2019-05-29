package main

import (
	"fmt"
	"gotest/generateId/seqno/seqno"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "bindo:maxinfo@tcp(127.0.0.1:3306)/bindo?parseTime=true")
	if err != nil {
		fmt.Println(fmt.Sprintln(err))
	}
	// fmt.Println(fmt.Sprintln(db.Debug().Exec(seqno.MigrateSQL()).Error))

	seqno := seqno.NewSeqNoGenerator(db, "test-logic-id").PrefixFormat("2006-01-02").Locker(seqno.NewRedisLocker(
		&redis.Pool{
			MaxIdle:     5,
			IdleTimeout: 30 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", "localhost:6379")
			},
		},
	)).Step(1)

	for n := 0; n < 100; n++ {
		num, err := seqno.Next()
		if err != nil {
			fmt.Println(fmt.Sprintln(err))
		}
		fmt.Println(num)
	}
}

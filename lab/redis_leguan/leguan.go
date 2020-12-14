package redis_leguan

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func GetToken() {
	dial, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	i := 1
	k := "k" + strconv.Itoa(i)

	for {
		do, _ := redis.String(dial.Do("watch", k))
		fmt.Println("watch return = " + do)
		if do != "OK" {
			break
		}
		dial.Send("MULTI")
		fmt.Println("sleeping!!!!")
		for i := 5; i > 0; i-- {
			time.Sleep(1 * time.Second)
			fmt.Printf("%d second ...\n", i)
		}

		dial.Send("SET", k, 1)
		r, err := dial.Do("EXEC")
		if err != nil {
			fmt.Println(err)
			dial.Do("unwatch", k)
			break
		} else {
			fmt.Println("transition return", r)
			break
		}
	}

	dial.Close()

}

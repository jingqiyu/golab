package main

import (
	_ "github.com/go-sql-driver/mysql"

	"xormlearn/lab/xo"
)

func main() {
	xo.GetAll()
	/*var err error
	var engine *xorm.Engine
	engine, err = xorm.NewEngine("mysql", "root:@/test?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(engine)*/
	/*user := &UserMedal{
		MedalId:     2000000002,
		AwardStatus: 0,
	}
	fmt.Println(user)
	_ = user.Insert(engine)*/
	//	err = user.Update(engine, []string{"score"})

	/*err = DeleteByUid(engine, "222222")
	fmt.Println(err)*/

	//fmt.Println("---------------------GetMaxLevelMedalByUidMedalId----------------")
	//u1, has, _ := GetMaxLevelMedalByUidMedalId(engine, "1", 2000000001)
	//fmt.Println(u1, has)
	/*
		u, _ := GetMedalListByUid(engine, "1")
		fmt.Println("---------------------GetMedalListByUid----------------")
		for _, v := range u {
			fmt.Println(v)
		}

		//err = DeleteByUid(engine, "1")
		fmt.Println(err)
		///lab.Test2Channel()*/

	/*c, err := redigo.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
	do, err := redigo.StringMap(c.Do("HGETALL", "ABCD"))

	fmt.Println(do, err)*/
	//lab.TestJson()
	//auto_commenter.AutoCommenter("/Users/jingqiyu/dev/trpc/pick", false)

	//lab.AutoGen()
	/*fmt.Println([]byte("A"))

	var i = 0
	for i < 10 {
		id, _ := idtool.NextId()
		encode := base36.StdEncoding.Encode(id)
		fmt.Println(string(encode))
		i++
	}
	id, _ := idtool.NextId()
	fmt.Println(id)

	fmt.Println(len(base36.ConverToBianry(id)))

	encode := base36.StdEncoding.Encode(id)
	fmt.Println(string(encode))*/

}

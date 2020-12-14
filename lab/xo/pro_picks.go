package xo

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// ProPicks object
type ProPicks struct {
	Id         int64     `xorm:"id autoincr pk" json:"id"`
	IsMedia    int64     `xorm:"is_media" json:"is_media"`
	Uid        string    `xorm:"uid"  json:"uid"`                // 普通用户为用户id 、 媒体用户为媒体id
	ProLevel   int       `xorm:"pro_level" json:"pro_level"`     //proLevel 外包pro用户等级
	ExtId      string    `xorm:"ext_id" json:"ext_id"`           // 附加id 如果媒体号的时候把用户id保存上，备用
	ArticleId  string    `xorm:"article_id"  json:"article_id"`  //文章id
	PeekType   int       `xorm:"peek_type" json:"peek_type"`     //Peek类型
	PeekReason string    `xorm:"peek_reason" json:"peek_reason"` // Peek理由
	CreateTime time.Time `xorm:"create_time created" json:"create_time"`
	UpdateTime time.Time `xorm:"update_time updated" json:"update_time"`
}

// *ProPicks -> TableName
func (p *ProPicks) TableName() string {
	return "pro_picks"
}

func GetAll() {
	var engine *xorm.Engine
	engine, err := xorm.NewEngine("mysql", "root:@/test?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
	}
	engine.ShowSQL(true)
	picks := make([]ProPicks, 0)
	engine.Table("pro_picks").Select("*").In("article_id", []string{"1", "2"}).
		GroupBy("article_id").Find(&picks)

	fmt.Println(picks)

}

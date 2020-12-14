package vo

import (
	"fmt"
	"time"

	"xorm.io/xorm"
)

type UserMedal struct {
	Id          int64     `xorm:"id autoincr pk"`
	Uid         string    `xorm:"uid" `
	MedalId     int64     `xorm:"medal_id"`
	Level       int       `xorm:"level"`
	ReachTime   time.Time `xorm:"reach_time created"`
	ReachScore  float64   `xorm:"score"`        //达成目标时 用户的积分值
	AwardStatus int       `xorm:"award_status"` // 领奖状态 0 未领奖 1 已领奖
	CreateTime  time.Time `xorm:"create_time created"`
	UpdateTime  time.Time `xorm:"update_time updated"`
}

func (u *UserMedal) TableName() string {
	return "lite_user_medal"
}

//DB:新增
func (u *UserMedal) Insert(engine *xorm.Engine) (err error) {
	_, err = engine.InsertOne(u)
	return
}

func (u *UserMedal) Update(engine *xorm.Engine, cols []string) (err error) {
	if len(cols) == 0 {
		err = fmt.Errorf("[UpdateUser] cols is empty")
		return
	}
	if _, err = engine.ID(u.Id).Cols(cols...).Update(u); err != nil {
		err = fmt.Errorf("[UpdateUser] user update to db failed: %v", err)
		return
	}
	return
}

func GetMaxLevelMedalByUidMedalId(engine *xorm.Engine, uid string, medalId int64) (user *UserMedal, has bool, err error) {
	u := new(UserMedal)
	has, err = engine.Where("uid = ? and medal_id=?", uid, medalId).Desc("level").Get(u)
	return u, has, err
}

func GetMedalListByUid(engine *xorm.Engine, uid string) (user []*UserMedal, err error) {
	pEveryOne := make([]*UserMedal, 0)
	err = engine.Where("uid = ?", uid).Find(&pEveryOne)
	return pEveryOne, err
}

//DB:新增
func DeleteByUid(engine *xorm.Engine, uid string) (err error) {
	u := new(UserMedal)
	rows, err := engine.Where("uid=?", uid).Delete(u)
	fmt.Println(rows)
	return
}

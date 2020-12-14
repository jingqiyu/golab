package lab

import (
	"encoding/json"
	"fmt"
)

type SystemMessageCoinsUserItem struct {
	CoinsUserCoralUid string `json:"coins_user_coral_uid,omitempty"`
	CoinsUserUin      string `json:"coins_user_uin,omitempty"`
	CoinsUserNick     string `json:"coins_user_nick,omitempty"`
	CoinsUserHead     string `json:"coins_user_head,omitempty"`
	MediaId           string `json:"media_id,omitempty"`
}
type SystemMessageItem struct {
	//MessageType   string `json:"message_type"` // common(纯文字）coinsPost (打榜作品） coinsUser （打榜用户）
	BussType      string `json:"busstype"` // editprofile 用户资料编辑(纯文字）coinsPost (打榜文字） coinsUser （打榜用户）
	GoType        int64  `json:"goType"`
	PubTime       string `json:"pub_time"`
	ReplyContent  string `json:"reply_content"`
	ReportTime    string `json:"reportTime"`
	Titlefrom     string `json:"titlefrom"`
	ArticleTitle  string `json:"article_title"`
	ArticleId     string `json:"article_id"`
	Url           string `json:"url"` // 文章链接
	ArticleImgurl string `json:"article_imgurl"`

	CoinsUserInfo SystemMessageCoinsUserItem `json:"coins_user_info"`

	GoldNum   int    `json:"gold_num"`
	GoldH5Url string `json:"gold_h5_url"`
}

func Print() {
	var sl []SystemMessageItem
	t1 := SystemMessageItem{
		BussType:     "editprofile",
		GoType:       2,
		PubTime:      "1594648489",
		ReplyContent: "您的资料修改已通过审核",
		ReportTime:   "1594648489",
		Titlefrom:    "2",
	}
	t2 := SystemMessageItem{
		BussType:      "coinsPost",
		GoType:        2,
		PubTime:       "1594648489",
		ReplyContent:  "恭喜！{{user_nick}}为您打赏{{gold_num}}金币。查看金币》 ",
		ReportTime:    "1594648489",
		Titlefrom:     "2",
		ArticleTitle:  "这是一个测试文章",
		ArticleId:     "20180312V16YCY00",
		Url:           "http://view.inews.qq.com/a/20180312V16YCY00",
		ArticleImgurl: "http://inews.gtimg.com/newsapp_ls/0/3023911654_150120/0",
		CoinsUserInfo: SystemMessageCoinsUserItem{
			CoinsUserCoralUid: "1155000836",
			CoinsUserUin:      "ec1f38b69fc563dd2d1fcc66bcfb5222a6",
			CoinsUserNick:     "测试昵称",
			CoinsUserHead:     "http://thirdwx.qlogo.cn/mmopen/vi_32/DYAIOgq83eoRciaSbMCtS0h3Gm78DntzjWa5NuvDfeEGnlY4VQkHBqn663J9gRCkdYHZOpicpoqudHT188NpRywA/132",
		},
		GoldH5Url: "www.todo.com",
		GoldNum:   99,
	}
	t3 := SystemMessageItem{
		BussType:     "coinsPost",
		GoType:       2,
		PubTime:      "1594648489",
		ReplyContent: "恭喜！{{user_nick}}为您打赏{{gold_num}}金币。查看金币》 ",
		ReportTime:   "1594648489",
		Titlefrom:    "2",
		CoinsUserInfo: SystemMessageCoinsUserItem{
			CoinsUserCoralUid: "1155000836",
			CoinsUserUin:      "ec1f38b69fc563dd2d1fcc66bcfb5222a6",
			CoinsUserNick:     "测试昵称",
			CoinsUserHead:     "http://thirdwx.qlogo.cn/mmopen/vi_32/DYAIOgq83eoRciaSbMCtS0h3Gm78DntzjWa5NuvDfeEGnlY4VQkHBqn663J9gRCkdYHZOpicpoqudHT188NpRywA/132",
		},
		GoldH5Url: "www.todo.com",
		GoldNum:   99,
	}

	sl = append(sl, t1)
	sl = append(sl, t2)
	sl = append(sl, t3)

	marshal, _ := json.Marshal(sl)
	fmt.Println(string(marshal))
}

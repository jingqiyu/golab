package lab

import (
	"encoding/json"
	"fmt"

	"github.com/spaolacci/murmur3"
)

type GreyConfig struct {
	Bucket    int64    `json:"bucket"`
	Slot      int64    `json:"slot"`
	WhiteList []string `json:"white_list"`
}

type CommentGreyConfigMap map[string]GreyConfig

func HitToggle(buss string, uid string) bool {

	var CommentGreyConfigMap = make(map[string]GreyConfig)

	data := `{
"supportComment":
{
    "white_list":["1916000086559"],
    "bucket": 1000,
    "slot" : 600
}
}`

	err := json.Unmarshal([]byte(data), &CommentGreyConfigMap)

	if err != nil {
		fmt.Println(err)
	}

	if _, ok := CommentGreyConfigMap[buss]; !ok {
		return false
	}

	cfg := CommentGreyConfigMap[buss]

	for _, v := range cfg.WhiteList {
		if v == uid {
			return true
		}
	}

	sum64 := murmur3.Sum64([]byte(uid))
	slot := sum64 % uint64(cfg.Bucket)
	fmt.Printf("sum64=%d||slot=%d", sum64, slot)

	return slot < uint64(cfg.Slot)

}

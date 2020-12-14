package game

import (
	"fmt"
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

func TestExpUp(t *testing.T) {
	fmt.Println((ExpUp(7) + 1) * 1.5)

	type User struct {
		UserId string `yaml:"user_id"`
		Nick   string `yaml:"nick"`
		Tag    string `yaml:"tag"`
	}

	type A struct {
		WhiteUserList []User `yaml:"white_user_list"`
	}

	a := A{
		WhiteUserList: []User{
			{"58795977", "穿裤子的发糕", "Tag1"},
			{"35664", "jing", "Tag2"},
		},
	}
	marshal, _ := yaml.Marshal(a)
	fmt.Println(string(marshal))

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

}

package redis_tool

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type RedisConf struct {
	Name      string `yaml:"Name"`
	Target    string `yaml:"Target"`
	Timeout   int    `yaml:"Timeout"`
	PassWord  string `yaml:"PassWord"`
	NameSpace string `yaml:"NameSpace"`
}

type RedisCnf struct {
	IdleConnNum       int    `yaml:"IdleConnNum"`
	ActiveConnNum     int    `yaml:"ActiveConnNum"`
	MaxConsumeNum     int    `yaml:"MaxConsumeNum"`
	IdleTimeout       int64  `yaml:"IdleTimeout"`
	ConnectTimeout    int64  `yaml:"ConnectTimeout"`
	ReadTimeout       int64  `yaml:"ReadTimeout"`
	WriteTimeout      int64  `yaml:"WriteTimeout"`
	EnableConnPool    bool   `yaml:"EnableConnPool"`
	EnableAutoRetry   bool   `yaml:"EnableAutoRetry"`
	EnableExtendProto bool   `yaml:"EnableExtendProto"`
	AUTH              string `yaml:"Auth"`
}

type RedisZkname struct {
	Host string `yaml:"Host"`
}

type RedisConfOld struct {
	Zkname RedisZkname `yaml:"Zkname"`
	Conf   RedisCnf    `yaml:"Conf"`
}

func Convert() {
	file, err := ioutil.ReadFile("/Users/jingqiyu/godev/xormlearn/lab/redis_tool/a.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var OldConfList map[string]RedisConfOld

	var NewConfYaml = make(map[string]RedisConf)

	err = yaml.Unmarshal(file, &OldConfList)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range OldConfList {
		if k == "COMMENT_COMMON" {
			fmt.Println(v)
		}
		tmp := RedisConf{
			Name:      v.Zkname.Host,
			Target:    "polaris://" + v.Zkname.Host,
			Timeout:   int(v.Conf.ReadTimeout),
			PassWord:  v.Conf.AUTH,
			NameSpace: "Production",
		}
		NewConfYaml[k] = tmp
	}

	marshal, _ := yaml.Marshal(NewConfYaml)
	fmt.Println(string(marshal))

}

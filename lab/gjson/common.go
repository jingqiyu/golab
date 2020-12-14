package gjson

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func T() {
	rawJson := `{"count":20,"data":{"diffuse":[{"Wid":"90QI50J7G02DYTLB000","Time":1602587692.3,"Diffuser":"2227809669"},{"Wid":"90QI50CLG02DYSXQ000","Time":1602587660,"Diffuser":"2227809669"},{"Wid":"90QI4ZKLG02DYQAS000","Time":1602587074,"Diffuser":"2227809669"},{"Wid":"90QI4ZQAG02DYQRY000","Time":1602587068,"Diffuser":"2227809669"},{"Wid":"90QI4ZUXG02DYR78000","Time":1602587061,"Diffuser":"2227809669"},{"Wid":"90QI4ZZFD02DYRMX000","Time":1602587055,"Diffuser":"2227809669"},{"Wid":"90QI4YUYG02DYNZR000","Time":1602586089,"Diffuser":"2227809669"},{"Wid":"90QI4Z12G02DYOK4000","Time":1602586084,"Diffuser":"2227809669"},{"Wid":"90QI4WG1G02DYGKD000","Time":1602582477,"Diffuser":"2227809669"},{"Wid":"90QI4W7LD02DYFYH000","Time":1602582426,"Diffuser":"2227809669"},{"Wid":"90QI4WA8G02DYG69000","Time":1602582421,"Diffuser":"2227809669"},{"Wid":"90QI4WCRG02DYGCA000","Time":1602582411,"Diffuser":"2227809669"},{"Wid":"90QI4VPWG02DYENG000","Time":1602582060,"Diffuser":"2227809669"},{"Wid":"90QI4V3NG02DYCYB000","Time":1602581248,"Diffuser":"2227809669"},{"Wid":"90QI4VAEG02DYDI1000","Time":1602581241,"Diffuser":"2227809669"},{"Wid":"90QI4VCMG02DYDP7000","Time":1602581232,"Diffuser":"2227809669"},{"Wid":"90QI4TAUG02DY0C0000","Time":1602579058,"Diffuser":"2227809669"},{"Wid":"90QI4SGFG02DXQNC000","Time":1602577203,"Diffuser":"2227809669"},{"Wid":"90QI4QYUG02DXHSO000","Time":1602575831,"Diffuser":"2227809669"}],"user":[{"Wid":"90QI4SJDG02DXRHY000","Time":1602577273}]},"errMsg":"success","errNo":0,"offset_info":"%7B%22LastReadTime%22%3A1602575831%2C%22RefreshType%22%3A1%2C%22MediaExtend%22%3A%7B%22LastCursor%22%3A%22%22%7D%7D"}`
	data := gjson.GetMany(rawJson, "data.video", "data.article", "data.user", "data.diffuse")

	for _, v := range data {
		v.ForEach(func(k, v gjson.Result) bool {
			fmt.Println(v.Get("Wid").String())
			fmt.Println(v.Get("Time").Int())
			fmt.Printf("%f", v.Get("Time").Float())
			return true
		})
	}

}

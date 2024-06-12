package main

import (
	"encoding/json"
	zubao "github.com/springeye/zubao"
)

func main() {
	client := zubao.NewSDKClient("1511111111111", "adfjklajfklafjklasdf", "https://www.hzzszf.com/interface/")
	//电表详情
	detail, err := client.AmmeterInstall("deviceidxxxxxxxx")
	if err != nil {
		panic(err)
	}
	marshal, err := json.Marshal(detail)
	if err != nil {
		panic(err)
	}
	println(string(marshal))

}

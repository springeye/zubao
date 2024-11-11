package main

import (
	json "github.com/go-json-experiment/json"
	zubao "github.com/springeye/zubao"
)

func main() {
	client := zubao.NewSDKClient("15884421212", "0670d65e90e3762b42817e9e5102bc4c", "https://www.hzzszf.com/interface/")
	//电表详情
	detail, err := client.WatermeterDetail("959063418714")
	if err != nil {
		panic(err)
	}
	marshal, err := json.Marshal(detail)
	if err != nil {
		panic(err)
	}
	println(string(marshal))

}

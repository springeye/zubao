package main

import zubao "github.com/springeye/zubao"

func main() {
	client := zubao.NewSDKClient("1511111111111", "adfjklajfklafjklasdf", "https://www.hzzszf.com/interface/")
	//电表详情
	detail, err := client.AmmeterDetail("deviceidxxxxxxxx")
	if err != nil {
		return
	}
	println(detail)

}

package zubao

import (
	"testing"
)

func TestSDKClient_AmmeterDetail(t *testing.T) {

	client := NewSDKClient("1511111111111", "adfjklajfklafjklasdf", "https://www.hzzszf.com/interface/")
	//电表详情
	detail, err := client.AmmeterDetail("deviceidxxxxxxxx")
	if err != nil {
		return
	}
	println(detail)

	gasmeterDetail, err := client.GasmeterDetail("deviceidxxxxxxxx")
	if err != nil {
		return
	}
	println(gasmeterDetail)

	watermeterDetail, err := client.WatermeterDetail("deviceidxxxxxxxx")
	if err != nil {
		return
	}

}

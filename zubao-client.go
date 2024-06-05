package zubao

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Result struct {
	Result int    `json:"result"` // 1: success, 0: failed
	Msg    string `json:"msg"`    // error message
}

// AmmeterDetail 电表详情
type AmmeterDetail struct {
	Factory      string `json:"factory"`      // 厂家
	Device       string `json:"device"`       // 设备号
	Voltage      string `json:"voltage"`      // 电压
	Currents     string `json:"currents"`     // 电流
	Power        string `json:"power"`        // 功率
	Battery      string `json:"battery"`      // 用电量
	SwitchState  int    `json:"switchState"`  // 开关机返回状态 1开， 0关
	NetworkState int    `json:"networkState"` // 网络连接状态 1正常， 0断网
}

// WatermeterDetail 水表详情
type WatermeterDetail struct {
	Factory      string `json:"factory"`      // 厂家
	Device       string `json:"device"`       // 设备号
	Tonnage      string `json:"tonnage"`      // 吨位
	SwitchState  int    `json:"switchState"`  // 开关机返回状态 1开， 0关
	NetworkState int    `json:"networkState"` // 网络连接状态 1正常， 0断网
	AnomalyState int    `json:"anomalyState"` // 故障状态 3电池故障, 2阀门故障, 1磁干扰， 0正常
}

// GasmeterDetail 气表详情
type GasmeterDetail struct {
	Factory      string `json:"factory"`      // 厂家
	Device       string `json:"device"`       // 设备号
	Stere        string `json:"stere"`        // 气方数
	Leakage      int    `json:"leakage"`      // 报警 1漏气，0正常
	SwitchState  int    `json:"switchState"`  // 开关机返回状态 1开， 0关
	NetworkState int    `json:"networkState"` // 网络连接状态 1正常， 0断网
	AnomalyState int    `json:"anomalyState"` // 故障状态 3电池故障, 2阀门故障, 1磁干扰， 0正常
}

type SDKClient struct {
	http      *http.Client
	account   string
	authToken string
	host      string
}

func NewSDKClient(account, authToken, host string) *SDKClient {
	return &SDKClient{
		http:      &http.Client{},
		account:   account,
		authToken: authToken,
		host:      host,
	}
}

type P map[string]string

func (c *SDKClient) get(params P) ([]byte, error) {
	url := c.host
	if len(params) > 0 {
		url += "?"
		for key, val := range params {
			url += fmt.Sprintf("%s=%s&", key, val)
		}
		url = strings.TrimSuffix(url, "&")
	}
	res, err := c.http.Get(fmt.Sprintf(url))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, err

}

// AmmeterSwitch 电表详情
// device: 电表编号
// value: 开关状态("ON"或者“OFF”)
func (c *SDKClient) AmmeterSwitch(device, value string) (*Result, error) {
	do := "ammeterSwitch"
	bytes, err := c.get(P{
		"do":     do,
		"device": device,
		"switch": value,
	})
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// AmmeterDetail 电表详情
// device: 电表编号
func (c *SDKClient) AmmeterDetail(device string) (*AmmeterDetail, error) {
	do := "ammeterDetail"
	bytes, err := c.get(P{
		"do":     do,
		"device": device,
	})
	if err != nil {
		return nil, err
	}
	var result AmmeterDetail
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// AmmeterInstall 安装电表
// device: 电表编号
func (c *SDKClient) AmmeterInstall(device string) (*Result, error) {
	do := "ammeterInstall"
	bytes, err := c.get(P{
		"do":     do,
		"device": device,
	})
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// WatermeterSwitch 水表开关
// device: 水表编号
// value: 开关状态("ON"或者“OFF”)
func (c *SDKClient) WatermeterSwitch(device, value string) (*Result, error) {
	do := "watermeterSwitch"
	bytes, err := c.get(P{
		"do":     do,
		"device": device,
		"switch": value,
	})
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// WatermeterInstall 安装水表
// device: 水表编号
func (c *SDKClient) WatermeterInstall(device string) (*Result, error) {
	do := "watermeterInstall"
	bytes, err := c.get(P{
		"do":     do,
		"device": device,
	})
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// WatermeterDetail 水表详情
// device: 水表编号
func (c *SDKClient) WatermeterDetail(device string) (*WatermeterDetail, error) {
	do := "watermeterDetail"
	bytes, err := c.get(P{
		"do":     do,
		"device": device,
	})
	if err != nil {
		return nil, err
	}
	var result WatermeterDetail
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// GasmeterSwitch 气表开关
// device: 气表编号
// value: 开关状态("ON"或者“OFF”)
func (c *SDKClient) GasmeterSwitch(device, value string) (*Result, error) {
	do := "gasmeterSwitch"
	bytes, err := c.get(P{
		"do":     do,
		"device": device,
		"switch": value,
	})
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// GasmeterInstall 气表水表
// device: 气表编号
func (c *SDKClient) GasmeterInstall(device string) (*Result, error) {
	do := "gasmeterInstall"
	bytes, err := c.get(P{
		"do":     do,
		"device": device,
	})
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// GasmeterDetail气表详情
// device: 气表编号
func (c *SDKClient) GasmeterDetail(device string) (*GasmeterDetail, error) {
	do := "gasmeterDetail"
	bytes, err := c.get(P{
		"do":     do,
		"device": device,
	})
	if err != nil {
		return nil, err
	}
	var result GasmeterDetail
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

package zubao

import (
	"fmt"
	json "github.com/go-json-experiment/json"
	"log/slog"
	"net/http"
	"strings"
)
import resty "github.com/go-resty/resty/v2"

type Response[T any] struct {
	Result  string `json:"result"`
	Message string `json:"msg"`
	Data    T      `json:"data"`
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
	http      *resty.Client
	account   string
	authToken string
	host      string
}

func NewSDKClient(account, authToken, host string) *SDKClient {
	client := resty.New().SetDebug(true)
	client.SetBaseURL(host)
	sdkClient := SDKClient{
		http:      client,
		account:   account,
		authToken: authToken,
		host:      host,
	}
	sdkClient.init()
	return &sdkClient
}
func (c *SDKClient) init() {
	if c.account == "" {
		slog.Error("account不能为空\n")
	}
	if c.authToken == "" {
		slog.Error("authToken不能为空\n")
	}
	c.http.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		request.FormData.Add("authToken", c.authToken)
		request.FormData.Add("account", c.account)
		return nil
	})
}
func NewSDKClientWithHttpClient(account, authToken, host string, httpClient *http.Client) *SDKClient {

	client := resty.NewWithClient(httpClient)
	client.SetBaseURL(host)

	sdkClient := SDKClient{
		http:      client,
		account:   account,
		authToken: authToken,
		host:      host,
	}
	sdkClient.init()
	return &sdkClient
}

type P map[string]string

func (c *SDKClient) post(params P) ([]byte, error) {
	resp, err := c.http.R().SetFormData(params).
		Post(c.host)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if strings.Contains(string(body), "Access Denied") {
		return nil, fmt.Errorf("没有访问权限")
	}
	return body, err
}

// AmmeterSwitch 电表详情
// device: 电表编号
// value: 开关状态("ON"或者“OFF”)
func (c *SDKClient) AmmeterSwitch(device, value string) (*Result, error) {
	do := "ammeterSwitch"
	bytes, err := c.post(P{
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
func (c *SDKClient) AmmeterDetail(device string) (*Response[AmmeterDetail], error) {
	do := "ammeterDetail"
	bytes, err := c.post(P{
		"do":     do,
		"device": device,
	})
	if err != nil {
		return nil, err
	}
	var result Response[AmmeterDetail]
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// AmmeterInstall 安装电表
// device: 电表编号
func (c *SDKClient) AmmeterInstall(device string) (*Result, error) {
	do := "ammeterInstall"
	bytes, err := c.post(P{
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
	bytes, err := c.post(P{
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
	bytes, err := c.post(P{
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
func (c *SDKClient) WatermeterDetail(device string) (*Response[WatermeterDetail], error) {
	do := "watermeterDetail"
	bytes, err := c.post(P{
		"do":     do,
		"device": device,
	})
	if err != nil {
		return nil, err
	}
	var result Response[WatermeterDetail]
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

// GasmeterSwitch 气表开关
// device: 气表编号
// value: 开关状态("ON"或者“OFF”)
func (c *SDKClient) GasmeterSwitch(device, value string) (*Result, error) {
	do := "gasmeterSwitch"
	bytes, err := c.post(P{
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
	bytes, err := c.post(P{
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
func (c *SDKClient) GasmeterDetail(device string) (*Response[GasmeterDetail], error) {
	do := "gasmeterDetail"
	bytes, err := c.post(P{
		"do":     do,
		"device": device,
	})
	if err != nil {
		return nil, err
	}
	var result Response[GasmeterDetail]
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

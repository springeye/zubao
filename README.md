## 用法
```go
package main

import "github.com/springeye/zubao"

func main() {
	client := zubao.NewSDKClient("1511111111111", "adfjklajfklafjklasdf", "https://www.hzzszf.com/interface/")
	
	//使用自己创建的http.Client
	//client := zubao.NewSDKClientWithHttpClient("1511111111111", "adfjklajfklafjklasdf", "https://www.hzzszf.com/interface/")
	//电表详情
	detail, err := client.AmmeterDetail("deviceidxxxxxxxx")
	if err != nil {
		return
	}
	println(detail)

}

```
## 用法
```go

//创建sdk客户端
client := NewClient("1511111111111", "adfjklajfklafjklasdf", "https://www.hzzszf.com/interface/")
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
```
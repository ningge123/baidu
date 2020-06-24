## 获取代码

```sh

go get -u github.com/medivhzhan/weapp/v2

```

## `目录`

> [百度开放平台](https://ai.baidu.com/ai-doc/)。

✅：代表已经通过线上测试
⚠️：代表还没有或者未完成

- [接口调用凭证](#接口调用凭证)
  - [getAccessToken](#getAccessToken) ✅
- [图像识别](#数据分析)
  - [组合接口API](#访问留存) ⚠️
  - [通用物体和场景识别高级版](#访问留存) ✅
  - [图像单主体检测](#访问留存) ⚠️
  - [动物识别](#访问留存) ✅
  - [植物识别](#访问留存) ✅
  - [logo识别](#访问留存) ✅
  - [果蔬识别](#访问留存) ✅
  - [自定义菜品识别](#访问留存) ⚠️
  - [菜品识别](#访问留存) ✅
  - [红酒识别](#访问留存) ✅
  - [货币识别](#访问留存) ⚠️
  - [地标识别](#访问留存) ⚠️
  - [门脸识别](#访问留存) ⚠️
  - [图像多主体检测](#访问留存) ⚠️
---

## 接口调用凭证

### getAccessToken

> 调用次数有限制 请注意缓存

[官方文档](https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu)

```go

import "github.com/medivhzhan/weapp/v2"

res, err := weapp.GetAccessToken("appid", "secret")
if err != nil {
    // 处理一般错误信息
    return
}

if err := res.GetResponseError(); err !=nil {
    // 处理微信返回错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```
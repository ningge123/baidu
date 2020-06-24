## 获取代码

```sh

go get -u github.com/ningge123/baidu

```

## `目录`

> [百度开放平台](https://ai.baidu.com/ai-doc/)。

✅：代表已经通过线上测试
⚠️：代表还没有或者未完成

- [接口调用凭证](#接口调用凭证)
  - [getAccessToken](#getAccessToken) ✅
- [图像识别]()
  - [组合接口API](#访问留存) ⚠️
  - [通用物体和场景识别高级版](#advancedCeneralIdentify) ✅
  - [图像单主体检测](#访问留存) ⚠️
  - [动物识别](#animalIdentify) ✅
  - [植物识别](#plantIdentify) ✅
  - [logo识别](#访问留存)  ⚠️
  - [果蔬识别](#ingredientIdentify) ✅
  - [自定义菜品识别](#访问留存) ⚠️
  - [菜品识别](#dishIdentify) ✅
  - [红酒识别](#redwineIdentify) ✅
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

import "github.com/ningge123/baidu"

res, err := baidu.GetAccessToken("appid", "secret")
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```

## 通用物体和场景识别高级版

### advancedCeneralIdentify

[官方文档](https://ai.baidu.com/ai-doc/IMAGERECOGNITION/Xk3bcxe21)

```go

import "github.com/ningge123/baidu"

res, err := baidu.AdvancedCeneralIdentify(token, image, baikeNum string)
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```

## 动物识别

### animalIdentify

[官方文档](https://ai.baidu.com/ai-doc/IMAGERECOGNITION/Zk3bcxdfr)

```go

import "github.com/ningge123/baidu"

res, err := baidu.AnimalIdentify(token, image, topNum, baikeNum string)
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```

## 植物识别

### plantIdentify

[官方文档](https://ai.baidu.com/ai-doc/IMAGERECOGNITION/Mk3bcxe9i)

```go

import "github.com/ningge123/baidu"

res, err := baidu.PlantIdentify(token, image, baikeNum string)
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```

## 果蔬识别

### ingredientIdentify

[官方文档](https://ai.baidu.com/ai-doc/IMAGERECOGNITION/wk3bcxevq)

```go

import "github.com/ningge123/baidu"

res, err := baidu.IngredientIdentify(token, image, topNum string)
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```

## 菜品识别

### dishIdentify

[官方文档](https://ai.baidu.com/ai-doc/IMAGERECOGNITION/tk3bcxbb0)

```go

import "github.com/ningge123/baidu"

res, err := baidu.DishIdentify(token, image, filterThreshold, topNum, baikeNum string)
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```

## 红酒识别

### redwineIdentify

[官方文档](https://ai.baidu.com/ai-doc/IMAGERECOGNITION/Tk3bcxctf)

```go

import "github.com/ningge123/baidu"

res, err := baidu.RedwineIdentify(token, image string)
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```
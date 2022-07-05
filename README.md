领星 ERP 开放平台 SDK
====================

## 文档地址

https://openapidoc.lingxing.com

## 安装

```go
go get github.com/hiscaler/lingxing
```

## 使用

```
var lingXingClient *LingXing
b, err := os.ReadFile("./config/config_test.json")
if err != nil {
    panic(fmt.Sprintf("Read config error: %s", err.Error()))
}
var c config.Config
err = jsoniter.Unmarshal(b, &c)
if err != nil {
    panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
}
lingXingClient = NewLingXing(c)
m.Run()
```

## 服务

### ~~授权~~

- ~~Auth(appId, appSecret string) (ar AuthResponse, err error)~~
- ~~Refresh(appId, refreshToken string) (ar AuthResponse, err error)~~

### 基础数据

- 亚马逊店铺信息

```go
lingXingClient.Services.BasicData.Sellers()
```

- ERP 账号列表

```go
lingXingClient.Services.BasicData.Accounts()
```

- 汇率

```go
lingXingClient.Services.BasicData.Rates()
```

### 销售

- 亚马逊订单列表

```go
lingXingClient.Services.Sale.Order.All()
```

- 亚马逊订单详情

```go
lingXingClient.Services.Sale.Order.One()
```

- 亚马逊自发货订单（FBM）列表

```go
lingXingClient.Services.Sale.FBM.Order.All()
```

- 查询 Listing

```go
lingXingClient.Services.Sale.Listing.All()
```

- 批量添加/编辑 Listing 配对

```go
lingXingClient.Services.Sale.Listing.Pair()
```

- 查询售后评价

```go
lingXingClient.Services.Sale.Review.All()
```

### 产品

- 本地产品列表

```go
lingXingClient.Services.Product.All()
```

- 本地产品详情

```go
lingXingClient.Services.Product.One()
```

- 本地产品品牌列表

```go
lingXingClient.Services.Product.Brand.All()
```

- 新增/更新品牌

```go
lingXingClient.Services.Product.Brand.Upsert()
```

- 产品分类列表

```go
lingXingClient.Services.Product.Category.All()
```

- 新增/更新分类

```go
lingXingClient.Services.Product.Category.Upsert()
```

### 客服

- 邮件列表

```go
lingXingClient.Services.CustomerService.Emails()
```

- 邮件详情

```go
lingXingClient.Services.CustomerService.Email()
```
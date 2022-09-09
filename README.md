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
// 加载配置文件
b, err := os.ReadFile("./config/config.json")
if err != nil {
    panic(fmt.Sprintf("Read config error: %s", err.Error()))
}
var c config.Config
err = jsoniter.Unmarshal(b, &c)
if err != nil {
    panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
}
lingXingClient = NewLingXing(c)
```

获取到 lingXingClient 后，则可以根据业务需求调用对应的服务获取数据，比如我要调取仓库数据，您可以使用以下代码：

```go
params := WarehousesQueryParams{
    Type: 1,
}
params.Limit = 20
for {
    items, nextOffset, isLastPage, err := lingXingClient.Services.Warehouse.All(params)
    if err != nil {
        break
    }
    for _, item := range items {
        // Read item
        _ = item
    }
    if isLastPage {
        break
    }
    params.Offset = nextOffset
}
```

### 注意

所有的列表方法都会返回四个值，分别是 `items`, `nextOffset`, `isLastPage`, `err`，它们所表示的含义为：

- items: 接口返回的数据
- nextOffset 下一次调取的位置
- isLastPage 是否为最后一页
- err 包含的错误，如果没有错误，则为 nil

**任何情况下，您都应该首先判断 err 是否为 nil，然后进行下一步的业务逻辑处理。**

如果是单个数据的请求，比如获取亚马逊订单详情：

```go
item, err := lingXingClient.Services.Sale.Order.One(1)
```

则会返回 `item`, `err` 两个值，第一个表示返回的数据，第二个则是错误信息，如果没有错误的话返回的是 nil，和列表数据一样，在处理 data 数据前，您需要先判断 err 是否为 nil，然后再进行下一步的处理。

## 服务

### 授权

- 获取 Token

```go
lingXingClient.Services.Authorization.GetToken()
```

- 刷新 Token

```go
lingXingClient.Services.Authorization.RefreshToken(refreshToken)
```

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
lingXingClient.Services.BasicData.Rates(RatesQueryParams{})
```

### 销售

- 亚马逊订单列表

```go
lingXingClient.Services.Sale.Order.All(AmazonOrdersQueryParams{})
```

- 亚马逊订单详情

```go
lingXingClient.Services.Sale.Order.One(orderId)
```

- 亚马逊自发货订单（FBM）列表

```go
lingXingClient.Services.Sale.FBM.Order.All(AmazonFBMOrdersQueryParams{})
```

- 亚马逊自发货订单（FBM）详情

```go
lingXingClient.Services.Sale.FBM.Order.One(number)
```

- 查询 Listing

```go
lingXingClient.Services.Sale.Listing.All(ListingsQueryParams{})
```

- 批量添加/编辑 Listing 配对

```go
lingXingClient.Services.Sale.Listing.Pair(ListingPairRequest{})
```

- 查询售后评价

```go
lingXingClient.Services.Sale.Review.All(ReviewsQueryParams{})
```

### FBA

- 发货单列表

```go
lingXingClient.Services.FBA.Shipment.All(FBAShipmentsQueryParams{})
```

- 发货单详情

```go
lingXingClient.Services.FBA.Shipment.One(shipmentSN)
```

- 查询 FBA 发货计划

```go
lingXingClient.Services.FBA.Shipment.Plans(FBAShipmentPlansQueryParams{})
```

- 查询 FBA 长期仓储费

```go
lingXingClient.Services.FBA.StorageFee.LongTerm(FBALongTermStorageFeesQueryParams{})
```

- 查询 FBA 月仓储费

```go
lingXingClient.Services.FBA.StorageFee.Month(FBAMonthStorageFeesQueryParams{})
```

### 产品

- 本地产品列表

```go
lingXingClient.Services.Product.All(ProductsQueryParams{})
```

- 本地产品详情

```go
lingXingClient.Services.Product.One(id)
```

- 本地产品品牌列表

```go
lingXingClient.Services.Product.Brand.All(BrandsQueryParams{})
```

- 新增/更新品牌

```go
lingXingClient.Services.Product.Brand.Upsert(UpsertBrandRequest{})
```

- 产品分类列表

```go
lingXingClient.Services.Product.Category.All(CategoriesQueryParams{})
```

- 新增/更新分类

```go
lingXingClient.Services.Product.Category.Upsert(UpsertCategoryRequest{})
```

- 产品辅料列表

```go
lingXingClient.Services.Product.AuxMaterial.All(ProductAuxMaterialsQueryParams)
```

- 添加/编辑辅料

```go
lingXingClient.Services.Product.AuxMaterial.Upsert(UpsertProductAuxMaterialRequest)
```

- 查询捆绑产品关系列表

```go
lingXingClient.Services.Product.Bundle.All(BundledProductsQueryParams{})
```

### 客服

- 邮件列表

```go
lingXingClient.Services.CustomerService.Email.All(CustomerServiceEmailsQueryParams{})
```

- 邮件详情

```go
lingXingClient.Services.CustomerService.Email.One(webMailUUID)
```

- Review 列表

```go
lingXingClient.Services.CustomerService.Review.All(CustomerServicesQueryParams{})
```

### 广告

- 广告组列表

```go
lingXingClient.Services.Ad.Groups(AdGroupsQueryParams{})
```

- 用户搜索词列表

```go
lingXingClient.Services.Ad.QueryWords(AdQueryWordsQueryParams{})
```

- 商品定位列表

```go
lingXingClient.Services.Ad.ProductTargets(AdProductTargetsQueryParams{})
```

### 采购单

- 采购计划列表

```go
lingXingClient.Services.Purchase.Plans(PurchasePlansQueryParams{})
```

- 采购单列表

```go
lingXingClient.Services.Purchase.Orders(PurchaseOrdersQueryParams{})
```

### 统计

- 产品表现列表

```go
lingXingClient.Services.Statistic.Products(ProductStatisticQueryParams)
```

### 仓库

- 本地仓库列表

```go
lingXingClient.Services.Warehouse.All(WarehousesQueryParams)
```

- 获取入库单列表

```go
lingXingClient.Services.Warehouse.InboundOrders(InboundsQueryParams{})
```

- 获取出库单列表

```go
lingXingClient.Services.Warehouse.OutboundOrders(OutboundsQueryParams{})
```

## 贡献

如果您在使用中遇到问题，或者有更好的建议或意见，您可以

1. [报告问题](https://github.com/hiscaler/lingxing/issues/new)
2. Fork 它并修改或实现需求，提交 Pull Request

谢谢！
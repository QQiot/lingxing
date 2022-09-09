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
lingXingClient.Services.Statistic.Products()
```

### 仓库

- 本地仓库列表

```go
lingXingClient.Services.Warehouse.All()
```

- 获取入库单列表

```go
lingXingClient.Services.Warehouse.InboundOrders(InboundsQueryParms{})
```

- 获取出库单列表

```go
lingXingClient.Services.Warehouse.OutboundOrders(OutboundsQueryParms{})
```

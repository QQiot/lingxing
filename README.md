领星 ERP 开放平台 SDK
====================

## 文档地址

https://openapidoc.lingxing.com

## 支持的方法

### 授权

- Auth(appId, appSecret string) (ar AuthResponse, err error)
- Refresh(appId, refreshToken string) (ar AuthResponse, err error)

### 基础数据

- Sellers() (items []Seller, err error)                                                     // 亚马逊店铺信息
- Accounts() (items []Account, err error)                                                   // ERP账号列表
- Rates(params RatesQueryParams) (items []Rate, nextOffset int, isLastPage bool, err error) // 汇率

### 销售

- AmazonOrders(params AmazonOrdersQueryParams) (items []AmazonOrder, nextOffset int, isLastPage bool, err error)          // 亚马逊订单列表
- AmazonOrder(params AmazonOrderQueryParams) (detail AmazonOrderDetail, err error)                                        // 亚马逊订单详情
- AmazonFBMOrders(params AmazonFBMOrdersQueryParams) (items []AmazonFBMOrder, nextOffset int, isLastPage bool, err error) // 亚马逊自发货订单（FBM）列表

### 产品

-- Products(params ProductsQueryParams) (items []Product, nextOffset int, isLastPage bool, err error) // 本地产品列表

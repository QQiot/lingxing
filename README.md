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

- Products(params ProductsQueryParams) (items []Product, nextOffset int, isLastPage bool, err error)      // 本地产品列表
- Product(id int) (item ProductDetail, err error)                                                         // 本地产品详情
- Brands(params BrandsQueryParams) (items []Brand, nextOffset int, isLastPage bool, err error)            // 本地产品品牌列表
- UpsertBrand(req UpsertBrandRequest) (items []Brand, err error)                                          // 新增/更新品牌
- Categories(params CategoriesQueryParams) (items []Category, nextOffset int, isLastPage bool, err error) // 产品分类列表
- UpsertCategory(req UpsertCategoryRequest) (items []Category, err error)                                 // 新增/更新分类

### 客服

- Emails(params EmailsQueryParams) (items []Email, nextOffset int, isLastPage bool, err error) // 邮件列表
- Email(webMailUUID string) (item EmailDetail, err error)                                      // 邮件详情
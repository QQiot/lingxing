领星 ERP 开放平台 SDK
====================

## 支持的方法

### 授权

- Auth(appId, appSecret string) (ar AuthResponse, err error)
- Refresh(appId, refreshToken string) (ar AuthResponse, err error)


### 基础数据

- Sellers() (items []Seller, err error)                                     // 亚马逊店铺信息
- Accounts() (items []Account, err error)                                   // ERP账号列表
- Rates(params RatesQueryParams) (items []Rate, isLastPage bool, err error) // 汇率

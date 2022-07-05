package lingxing

// 亚马逊订单
type amazonOrderService struct {
	Order orderService
}

type orderService service

// 自发货订单
type saleService struct {
	FBM     saleFBMService
	Order   orderService
	Listing listingService
}

type saleFBMService struct {
	Order fbmOrderService
}

type fbmOrderService service

// Listing
type listingService service

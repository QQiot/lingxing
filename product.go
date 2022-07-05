package lingxing

type prodService service

type productServiceProduct service
type productServiceBrand service
type productServiceCategory service

type productService struct {
	productServiceProduct
	Brand    productServiceBrand
	Category productServiceCategory
}

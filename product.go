package lingxing

type productService struct {
	productProductService
	Brand       productBrandService
	Category    productCategoryService
	AuxMaterial productAuxMaterialService
}

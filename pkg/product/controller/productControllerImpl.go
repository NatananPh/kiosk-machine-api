package controller

import "github.com/NatananPh/kiosk-machine-api/pkg/product/service"

type productControllerImpl struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productControllerImpl{
		productService: productService,
	}
}
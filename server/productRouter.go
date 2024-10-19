package server

import (
	"github.com/NatananPh/kiosk-machine-api/pkg/product/controller"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/repository"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/service"
)

func (s *echoServer) registerProductRoutes() {
	productRepository := repository.NewProductRepository(s.db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router := s.app.Group("/v1/products")
	router.GET("", productController.GetProducts)
	router.GET("/:id", productController.GetProductByID)
	router.PUT("/:id", productController.UpdateProduct)
	router.DELETE("/:id", productController.DeleteProduct)
}
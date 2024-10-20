package server

import (
	"github.com/NatananPh/kiosk-machine-api/pkg/middleware"
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

	adminMiddleware := middleware.RoleBasedMiddleware()
	router.POST("", productController.CreateProduct, adminMiddleware)
	router.PUT("/:id", productController.UpdateProduct, adminMiddleware)
	router.DELETE("/:id", productController.DeleteProduct, adminMiddleware)
}
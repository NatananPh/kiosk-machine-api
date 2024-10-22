package controller

import "github.com/labstack/echo/v4"

type ProductController interface {
	CreateProduct(ctx echo.Context) error
	GetProducts(ctx echo.Context) error
	GetProductByID(ctx echo.Context) error
	UpdateProduct(ctx echo.Context) error
	DeleteProduct(ctx echo.Context) error
	PurchaseProduct(ctx echo.Context) error
}
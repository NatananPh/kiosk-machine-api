package controller

import "github.com/labstack/echo/v4"

type ProductController interface {
	GetProducts(ctx echo.Context) error
	GetProductByID(ctx echo.Context) error
	UpdateProduct(ctx echo.Context) error
	DeleteProduct(ctx echo.Context) error
}
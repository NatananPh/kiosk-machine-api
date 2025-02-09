package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/NatananPh/kiosk-machine-api/pkg/custom"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/exception"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type productControllerImpl struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productControllerImpl{
		productService: productService,
	}
}

func (controller *productControllerImpl) CreateProduct(ctx echo.Context) error {
	var product model.ProductCreateRequest
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
	}

	createdProduct, err := controller.productService.CreateProduct(&product)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, json.RawMessage(`{"error": "Internal server error"}`))
	}
	return ctx.JSON(http.StatusCreated, createdProduct)
}

func (controller *productControllerImpl) GetProducts(ctx echo.Context) error {
	productFilter := new(model.ProductFilter)
	if err := ctx.Bind(productFilter); err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
	}

	validating := validator.New()
	if err := validating.Struct(productFilter); err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
	}

	products, err := controller.productService.GetProducts(productFilter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, json.RawMessage(`{"error": "Internal server error"}`))
	}
	return ctx.JSON(http.StatusOK, products)
}

func (controller *productControllerImpl) GetProductByID(ctx echo.Context) error {
	id, err := custom.GetParamInt(ctx, "id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
	}

	product, err := controller.productService.GetProductByID(id)
	if err != nil {
		if errors.Is(err, &exception.ProductNotFound{}) {
			return ctx.JSON(http.StatusNotFound, json.RawMessage(`{"error": "Product not found"}`))
		}
		return ctx.JSON(http.StatusInternalServerError, json.RawMessage(`{"error": "Internal server error"}`))
	}
	return ctx.JSON(http.StatusOK, product)
}

func (controller *productControllerImpl) UpdateProduct(ctx echo.Context) error {
	id, err := custom.GetParamInt(ctx, "id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
	}

	var product *model.Product
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
	}

	_, err = controller.productService.UpdateProduct(id, product)
	if err != nil {
		if errors.Is(err, &exception.ProductNotFound{}) {
			return ctx.JSON(http.StatusNotFound, json.RawMessage(`{"error": "Product not found"}`))
		}
		return ctx.JSON(http.StatusInternalServerError, json.RawMessage(`{"error": "Internal server error"}`))
	}
	return ctx.JSON(http.StatusOK, json.RawMessage(`{"message": "Product updated"}`))
}

func (controller *productControllerImpl) DeleteProduct(ctx echo.Context) error {
	id, err := custom.GetParamInt(ctx, "id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
	}

	err = controller.productService.DeleteProduct(id)
	if err != nil {
		if errors.Is(err, &exception.ProductNotFound{}) {
			return ctx.JSON(http.StatusNotFound, json.RawMessage(`{"error": "Product not found"}`))
		}
		return ctx.JSON(http.StatusInternalServerError, json.RawMessage(`{"error": "Internal server error"}`))
	}
	return ctx.JSON(http.StatusOK, json.RawMessage(`{"message": "Product deleted"}`))
}

func (controller *productControllerImpl) PurchaseProduct(ctx echo.Context) error {
	id, err := custom.GetParamInt(ctx, "id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
	}

	requestBody := new(model.ProductPurchaseRequest)
	if err := ctx.Bind(requestBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
	}

	product, err := controller.productService.PurchaseProduct(id, uint(requestBody.PaymentAmount))
	if err != nil {
		if errors.Is(err, &exception.ProductNotFound{}) {
			return ctx.JSON(http.StatusNotFound, json.RawMessage(`{"error": "Product not found"}`))
		}
		if errors.Is(err, &exception.InsufficientMoney{}) {
			return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Insufficient money"}`))
		}
		if errors.Is(err, &exception.ProductOutOfStock{}) {
			return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Product out of stock"}`))
		}
		return ctx.JSON(http.StatusInternalServerError, json.RawMessage(`{"error": "Internal server error"}`))
	
	}
	return ctx.JSON(http.StatusOK, product)
}
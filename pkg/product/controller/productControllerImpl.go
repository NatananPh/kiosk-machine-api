package controller

import (
	"encoding/json"
	"net/http"

	"github.com/NatananPh/kiosk-machine-api/pkg/custom"
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
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	createdProduct, err := controller.productService.CreateProduct(&product)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, createdProduct)
}

func (controller *productControllerImpl) GetProducts(ctx echo.Context) error {
	productFilter := new(model.ProductFilter)
	if err := ctx.Bind(productFilter); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	validating := validator.New()
	if err := validating.Struct(productFilter); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	products, err := controller.productService.GetProducts(productFilter)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, products)
}

func (controller *productControllerImpl) GetProductByID(ctx echo.Context) error {
	id, err := custom.GetParamInt(ctx, "id")
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	product, err := controller.productService.GetProductByID(id)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, product)
}

func (controller *productControllerImpl) UpdateProduct(ctx echo.Context) error {
	id, err := custom.GetParamInt(ctx, "id")
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	var product *model.Product
	if err := ctx.Bind(&product); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	_, err = controller.productService.UpdateProduct(id, product)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, json.RawMessage(`{"message": "Product updated"}`))
}

func (controller *productControllerImpl) DeleteProduct(ctx echo.Context) error {
	id, err := custom.GetParamInt(ctx, "id")
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	err = controller.productService.DeleteProduct(id)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, json.RawMessage(`{"message": "Product deleted"}`))
}

func (controller *productControllerImpl) PurchaseProduct(ctx echo.Context) error {
	id, err := custom.GetParamInt(ctx, "id")
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	requestBody := new(model.ProductPurchaseRequest)
	if err := ctx.Bind(requestBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request body"}`))
	}

	product, err := controller.productService.PurchaseProduct(id, uint(requestBody.PaymentAmount))
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, product)
}
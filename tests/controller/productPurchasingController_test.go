package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NatananPh/kiosk-machine-api/pkg/product/controller"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPurchasingProductSuccess(t *testing.T) {
	e := echo.New()

	productPurchaseRequest := &model.ProductPurchaseRequest{
		PaymentAmount: 100,
	}

	productPurchaseResponse := &model.ProductPurchaseResponse{
		ProductID: 1,
		Change: map[string]int{
			"50": 1,
		},
	}

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("PurchaseProduct", 1, uint(productPurchaseRequest.PaymentAmount)).Return(productPurchaseResponse, nil)
	productServiceMock.On("GetProductByID", 1).Return(&model.Product{
		ID:       1,
		Name:     "Test Product",
		Price:    50,
		Amount:   10,
		Category: "Category 1",
	}, nil)

	productController := controller.NewProductController(productServiceMock)

	body, _ := json.Marshal(productPurchaseRequest)
	req := httptest.NewRequest(http.MethodPost, "/v1/products/1/purchase", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id/purchase")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.PurchaseProduct(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"product_id":1,"change":{"50":1}}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "PurchaseProduct", 1, uint(productPurchaseRequest.PaymentAmount))
}

func TestPurchasingProductInvalidRequest(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodPost, "/v1/products/1/purchase", bytes.NewReader([]byte(`{}`)))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, productController.PurchaseProduct(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "Invalid request"}`, rec.Body.String())
	}

	productServiceMock.AssertNotCalled(t, "PurchaseProduct")
}

func TestPurchasingProductInternalServerError(t *testing.T) {
	e := echo.New()

	productPurchaseRequest := &model.ProductPurchaseRequest{
		PaymentAmount: 100,
	}

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("PurchaseProduct", 1, uint(productPurchaseRequest.PaymentAmount)).Return(nil, assert.AnError)

	productController := controller.NewProductController(productServiceMock)

	body, _ := json.Marshal(productPurchaseRequest)
	req := httptest.NewRequest(http.MethodPost, "/v1/products/1/purchase", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id/purchase")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.PurchaseProduct(ctx)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"error": "Internal server error"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "PurchaseProduct", 1, uint(productPurchaseRequest.PaymentAmount))
}
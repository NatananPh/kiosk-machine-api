package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/NatananPh/kiosk-machine-api/pkg/product/controller"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/exception"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/service"
)
func TestGettingProductByIDSuccess(t *testing.T) {
	e := echo.New()

	productMock := &model.Product{
		ID:    1,
		Name:  "Test Product",
		Price: 100,
		Amount: 10,
		Category: "Category 1",
	}

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("GetProductByID", 1).Return(productMock, nil)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.GetProductByID(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"id":1,"name":"Test Product","price":100,"amount":10,"category":"Category 1"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "GetProductByID", 1)
}

func TestGettingProductByIDInvalidRequest(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodGet, "/v1/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, productController.GetProductByID(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "Invalid request"}`, rec.Body.String())
	}

	productServiceMock.AssertNotCalled(t, "GetProductByID")
}

func TestGettingProductByIDProductNotFound(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("GetProductByID", 1).Return(nil, &exception.ProductNotFound{})

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodGet, "/v1/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.GetProductByID(ctx)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.JSONEq(t, `{"error": "Product not found"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "GetProductByID", 1)
}

func TestGettingProductByIDInternalServerError(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("GetProductByID", 1).Return(nil, assert.AnError)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodGet, "/v1/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.GetProductByID(ctx)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"error": "Internal server error"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "GetProductByID", 1)
}

func TestGettingProductsSuccess(t *testing.T) {
	e := echo.New()

	productsMock := []*model.Product{
		{
			ID:    1,
			Name:  "Test Product 1",
			Price: 100,
			Amount: 10,
			Category: "Category 1",
		},
		{
			ID:    2,
			Name:  "Test Product 2",
			Price: 200,
			Amount: 20,
			Category: "Category 2",
		},
	}

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("GetProducts", &model.ProductFilter{}).Return(productsMock, nil)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodGet, "/v1/products", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, productController.GetProducts(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `[{"id":1,"name":"Test Product 1","price":100,"amount":10,"category":"Category 1"},{"id":2,"name":"Test Product 2","price":200,"amount":20,"category":"Category 2"}]`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "GetProducts", &model.ProductFilter{})
}

func TestGettingProductsInternalServerError(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("GetProducts", &model.ProductFilter{}).Return(nil, assert.AnError)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodGet, "/v1/products", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, productController.GetProducts(ctx)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"error": "Internal server error"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "GetProducts", &model.ProductFilter{})
}

func TestCreatingProductSuccess(t *testing.T) {
	e := echo.New()

	productCreateRequest := &model.ProductCreateRequest{
		Name:   "Test Product",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productMock := &model.Product{
		ID:    1,
		Name:  "Test Product",
		Price: 100,
		Amount: 10,
		Category: "Category 1",
	}

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("CreateProduct", productCreateRequest).Return(productMock, nil)

	productController := controller.NewProductController(productServiceMock)

	body, _ := json.Marshal(productCreateRequest)
	req := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products")

	if assert.NoError(t, productController.CreateProduct(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, `{"id":1,"name":"Test Product","price":100,"amount":10,"category":"Category 1"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "CreateProduct", productCreateRequest)
}

func TestCreatingProductInvalidRequest(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader([]byte(`{}`)))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, productController.CreateProduct(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "Invalid request"}`, rec.Body.String())
	}

	productServiceMock.AssertNotCalled(t, "CreateProduct")
}

func TestCreatingProductInternalServerError(t *testing.T) {
	e := echo.New()

	productCreateRequest := &model.ProductCreateRequest{
		Name:   "Test Product",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("CreateProduct", productCreateRequest).Return(nil, assert.AnError)

	productController := controller.NewProductController(productServiceMock)

	body, _ := json.Marshal(productCreateRequest)
	req := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products")

	if assert.NoError(t, productController.CreateProduct(ctx)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"error": "Internal server error"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "CreateProduct", productCreateRequest)
}

// func (controller *productControllerImpl) UpdateProduct(ctx echo.Context) error {
// 	id, err := custom.GetParamInt(ctx, "id")
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
// 	}

// 	var product *model.Product
// 	if err := ctx.Bind(&product); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, json.RawMessage(`{"error": "Invalid request"}`))
// 	}

// 	_, err = controller.productService.UpdateProduct(id, product)
// 	if err != nil {
// 		if errors.Is(err, &exception.ProductNotFound{}) {
// 			return ctx.JSON(http.StatusNotFound, json.RawMessage(`{"error": "Product not found"}`))
// 		}
// 		return ctx.JSON(http.StatusInternalServerError, json.RawMessage(`{"error": "Internal server error"}`))
// 	}
// 	return ctx.JSON(http.StatusOK, json.RawMessage(`{"message": "Product updated"}`))
// }
func TestUpdatingProductSuccess(t *testing.T) {
	e := echo.New()

	productUpdateRequest := &model.Product{
		ID:     1,
		Name:   "Test Product",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("UpdateProduct", 1, productUpdateRequest).Return(productUpdateRequest, nil)

	productController := controller.NewProductController(productServiceMock)

	body, _ := json.Marshal(productUpdateRequest)
	req := httptest.NewRequest(http.MethodPut, "/v1/products/1", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.UpdateProduct(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"message": "Product updated"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "UpdateProduct", 1, productUpdateRequest)
}

func TestUpdatingProductInvalidRequest(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodPut, "/v1/products/1", bytes.NewReader([]byte(`{}`)))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, productController.UpdateProduct(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "Invalid request"}`, rec.Body.String())
	}

	productServiceMock.AssertNotCalled(t, "UpdateProduct")
}

func TestUpdatingProductProductNotFound(t *testing.T) {
	e := echo.New()

	productUpdateRequest := &model.Product{
		ID:     1,
		Name:   "Test Product",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("UpdateProduct", 1, productUpdateRequest).Return(nil, &exception.ProductNotFound{})

	productController := controller.NewProductController(productServiceMock)

	body, _ := json.Marshal(productUpdateRequest)
	req := httptest.NewRequest(http.MethodPut, "/v1/products/1", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.UpdateProduct(ctx)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.JSONEq(t, `{"error": "Product not found"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "UpdateProduct", 1, productUpdateRequest)
}

func TestUpdatingProductInternalServerError(t *testing.T) {
	e := echo.New()

	productUpdateRequest := &model.Product{
		ID:     1,
		Name:   "Test Product",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("UpdateProduct", 1, productUpdateRequest).Return(nil, assert.AnError)

	productController := controller.NewProductController(productServiceMock)

	body, _ := json.Marshal(productUpdateRequest)
	req := httptest.NewRequest(http.MethodPut, "/v1/products/1", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.UpdateProduct(ctx)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"error": "Internal server error"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "UpdateProduct", 1, productUpdateRequest)
}

func TestDeletingProductSuccess(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("DeleteProduct", 1).Return(nil)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodDelete, "/v1/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.DeleteProduct(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"message": "Product deleted"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "DeleteProduct", 1)
}

func TestDeletingProductInvalidRequest(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodDelete, "/v1/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, productController.DeleteProduct(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "Invalid request"}`, rec.Body.String())
	}

	productServiceMock.AssertNotCalled(t, "DeleteProduct")
}

func TestDeletingProductProductNotFound(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("DeleteProduct", 1).Return(&exception.ProductNotFound{})

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodDelete, "/v1/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.DeleteProduct(ctx)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.JSONEq(t, `{"error": "Product not found"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "DeleteProduct", 1)
}

func TestDeletingProductInternalServerError(t *testing.T) {
	e := echo.New()

	productServiceMock := new(service.ProductServiceMock)

	productServiceMock.On("DeleteProduct", 1).Return(assert.AnError)

	productController := controller.NewProductController(productServiceMock)

	req := httptest.NewRequest(http.MethodDelete, "/v1/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/v1/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, productController.DeleteProduct(ctx)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"error": "Internal server error"}`, rec.Body.String())
	}

	productServiceMock.AssertCalled(t, "DeleteProduct", 1)
}



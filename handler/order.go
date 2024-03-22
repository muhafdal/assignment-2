package handler

import (
	"h8-assignment-2/dto"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/service/order_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	OrderService order_service.Service
}

func NewOrderHandler(orderService order_service.Service) orderHandler {
	return orderHandler{
		OrderService: orderService,
	}
}

// CreateOrder godoc
// @Tags orders
// @Description Create Order Data
// @ID create-new-order
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewOrderRequestDto true "request body json"
// @Success 201 {object} dto.NewOrderResponseDto
// @Router /orders [post]
func (oh *orderHandler) CreateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequestDto

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := oh.OrderService.CreateOrderWithItems(newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// @Tags orders
// @Description Get Order with Item Data
// @ID get-orders-with-items
// @Produce json
// @Success 200 {object} GetOrdersResponse
// @Router /orders [get]
func (oh *orderHandler) GetOrders(ctx *gin.Context) {
	response, err := oh.OrderService.GetOrders()

	if err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(response.BaseResponse.StatusCode, response)
}

// @Tags orders
// @Description Update Order Data By Id
// @ID update-order
// @Accept json
// @Produce json
// @Param orderId path int true "order's id"
// @Param RequestBody body dto.NewOrderRequestDto true "request body json"
// @Success 200 {object} dto.NewOrderResponseDto
// @Router /orders/{orderId} [put]
func (oh *orderHandler) UpdateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequestDto

	orderId, errParam := strconv.Atoi(ctx.Param("orderId"))

	if errParam != nil {
		errParsingParam := errs.NewBadRequest("orderId has to be a valid number value")
		ctx.AbortWithStatusJSON(errParsingParam.Status(), errParsingParam)
		return
	}

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := oh.OrderService.UpdateOrder(orderId, newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)

		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (oh *orderHandler) DeleteOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "orderId has to be a valid number value"})
		return
	}

	err = oh.OrderService.DeleteOrder(orderId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

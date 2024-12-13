package handler

import (
	"Marketplace/internal/domain/order"
	"Marketplace/internal/service/interfaces"
	"Marketplace/internal/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	orderService interfaces.OrderService
}

func NewOrderHandler(service interfaces.OrderService) *OrderHandler {
	return &OrderHandler{orderService: service}
}

// Create godoc
// @Summary Create a new order
// @Description Creates a new order for a company and influencer
// @Tags order
// @Accept json
// @Produce json
// @Param order body order.CreateRequest true "Order details"
// @Success 201 {object} response.Response "Order created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 500 {object} response.Response "Failed to create order"
// @Router /order [post]
func (h *OrderHandler) Create(ctx *gin.Context) {
	var req order.CreateRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid input", nil, err.Error()))
		return
	}

	res, err := h.orderService.Create(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "failed to create order", nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "order created successfully", res, nil))
}

// GetByID godoc
// @Summary Get order by ID
// @Description Retrieve an order by its ID
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response "Order retrieved successfully"
// @Failure 404 {object} response.Response "Order not found"
// @Failure 500 {object} response.Response "Failed to retrieve order"
// @Router /order/{id} [get]
func (h *OrderHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.orderService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ClientResponse(http.StatusNotFound, "order not found", nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "order retrieved successfully", res, nil))
}

// UpdateStatus godoc
// @Summary Update order status
// @Description Update the status of an existing order
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param status body order.UpdateStatusRequest true "New status"
// @Success 200 {object} response.Response "Status updated successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 500 {object} response.Response "Failed to update status"
// @Router /order/{id}/status [put]
func (h *OrderHandler) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var req order.UpdateStatusRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid input", nil, err.Error()))
		return
	}

	err := h.orderService.UpdateStatus(ctx.Request.Context(), id, req.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "failed to update status", nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "status updated successfully", nil, nil))
}

// GetByCompanyID godoc
// @Summary Get orders by company ID
// @Description Retrieve all orders for a specific company
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} response.Response "Orders retrieved successfully"
// @Failure 500 {object} response.Response "Failed to retrieve orders"
// @Router /order/company/{id} [get]
func (h *OrderHandler) GetByCompanyID(ctx *gin.Context) {
	companyID := ctx.Param("id")
	res, err := h.orderService.GetByCompanyID(ctx.Request.Context(), companyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "failed to retrieve orders", nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "orders retrieved successfully", res, nil))
}

// GetByInfluencerID godoc
// @Summary Get orders by influencer ID
// @Description Retrieve all orders for a specific influencer
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "Influencer ID"
// @Success 200 {object} response.Response "Orders retrieved successfully"
// @Failure 500 {object} response.Response "Failed to retrieve orders"
// @Router /order/influencer/{id} [get]
func (h *OrderHandler) GetByInfluencerID(ctx *gin.Context) {
	influencerID := ctx.Param("id")
	res, err := h.orderService.GetByInfluencerID(ctx.Request.Context(), influencerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "failed to retrieve orders", nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "orders retrieved successfully", res, nil))
}

// Delete godoc
// @Summary Delete an order by ID
// @Description Delete a specific order by its ID
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response "Order deleted successfully"
// @Failure 500 {object} response.Response "Failed to delete order"
// @Router /order/{id} [delete]
func (h *OrderHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.orderService.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "failed to delete order", nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "order deleted successfully", nil, nil))
}

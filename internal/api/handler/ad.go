package handler

import (
	"Marketplace/internal/domain/ad"
	"Marketplace/internal/service/interfaces"
	"Marketplace/internal/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdHandler struct {
	adService interfaces.AdService
}

func NewAdHandler(service interfaces.AdService) *AdHandler {
	return &AdHandler{
		adService: service,
	}
}

// CreateAd godoc
// @Summary Create a new ad
// @Description Create a new ad with the input payload
// @Tags ad
// @Accept json
// @Produce json
// @Param ad body ad.CreateRequest true "Ad Creation Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /ad/create [post]
func (h *AdHandler) Create(ctx *gin.Context) {
	req := ad.CreateRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "invalid input", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	// Валидация данных
	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	// Создание объявления
	adID, err := h.adService.Create(ctx.Request.Context(), req)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to create ad", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "ad created successfully", gin.H{"ad_id": adID}, nil)
	ctx.JSON(http.StatusCreated, successRes)
}

// GetAdByID godoc
// @Summary Get ad by ID
// @Description Get ad details by its ID
// @Tags ad
// @Accept json
// @Produce json
// @Param id path string true "Ad ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /ad/{id} [get]
func (h *AdHandler) GetByID(ctx *gin.Context) {
	adID := ctx.Param("id")

	// Получение объявления по ID
	adData, err := h.adService.GetByID(ctx.Request.Context(), adID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusNotFound, "ad not found", nil, err.Error())
		ctx.JSON(http.StatusNotFound, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "ad found", adData, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// GetAllAds godoc
// @Summary Get all ads
// @Description Get a list of all ads
// @Tags ad
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /ad/all [get]
func (h *AdHandler) GetAll(ctx *gin.Context) {
	// Получение всех объявлений
	ads, err := h.adService.GetAll(ctx.Request.Context())
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve ads", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "ads retrieved successfully", ads, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// UpdateAdStatus godoc
// @Summary Update ad status
// @Description Update the status of an ad (e.g., open/closed)
// @Tags ad
// @Accept json
// @Produce json
// @Param id path string true "Ad ID"
// @Param status body string true "Ad Status"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /ad/{id}/status [put]
func (h *AdHandler) UpdateStatus(ctx *gin.Context) {
	adID := ctx.Param("id")
	var req struct {
		Status string `json:"status"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "invalid status input", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	// Обновление статуса объявления
	err := h.adService.UpdateStatus(ctx.Request.Context(), adID, req.Status)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to update ad status", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "ad status updated successfully", nil, nil)
	ctx.JSON(http.StatusOK, successRes)
}

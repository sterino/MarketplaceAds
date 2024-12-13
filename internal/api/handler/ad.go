package handler

import (
	"Marketplace/internal/domain/ad"
	"Marketplace/internal/service/interfaces"
	"Marketplace/internal/utils/response"
	"github.com/gin-gonic/gin"
	"log"
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
// @Security BearerAuth
// @Router /ad/create [post]
func (h *AdHandler) Create(ctx *gin.Context) {
	req := ad.CreateRequest{}
	user, exists := ctx.Get("user")
	if !exists || user == nil {
		errRes := response.ClientResponse(http.StatusUnauthorized, "User not authorized", nil, nil)
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}
	if err := ctx.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "invalid input", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
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
// @Security BearerAuth
// @Router /ad/{id} [get]
func (h *AdHandler) GetByID(ctx *gin.Context) {
	adID := ctx.Param("id")

	user, exists := ctx.Get("user")
	if !exists || user == nil {
		errRes := response.ClientResponse(http.StatusUnauthorized, "User not authorized", nil, nil)
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}

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
// @Security BearerAuth
// @Router /ad/all [get]
func (h *AdHandler) GetAll(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists || user == nil {
		errRes := response.ClientResponse(http.StatusUnauthorized, "User not authorized", nil, nil)
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}
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
// @Security BearerAuth
// @Router /ad/{id}/status [put]
func (h *AdHandler) UpdateStatus(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists || user == nil {
		log.Printf("User not found in context: %v", user)
		errRes := response.ClientResponse(http.StatusUnauthorized, "User not authorized", nil, nil)
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}
	adID := ctx.Param("id")
	var req struct {
		Status string `json:"status"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "invalid status input", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := h.adService.UpdateStatus(ctx.Request.Context(), adID, req.Status)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to update ad status", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "ad status updated successfully", nil, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// GetByCompanyID godoc
// @Summary Get ads by company ID
// @Description Get all ads for a specific company
// @Tags ad
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Security BearerAuth
// @Router /ad/company/{id} [get]
func (h *AdHandler) GetByCompanyID(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists || user == nil {
		errRes := response.ClientResponse(http.StatusUnauthorized, "User not authorized", nil, nil)
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}
	companyID := ctx.Param("id")

	ads, err := h.adService.GetByCompanyID(ctx.Request.Context(), companyID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve ads", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "ads retrieved successfully", ads, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// Delete godoc
// @Summary Delete an ad
// @Description Delete an ad by its ID
// @Tags ad
// @Accept json
// @Produce json
// @Param id path string true "Ad ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Security BearerAuth
// @Router /ad/delete/{id} [delete]
func (h *AdHandler) Delete(ctx *gin.Context) {
	adID := ctx.Param("id")

	user, _ := ctx.Get("user")
	if user == nil {
		response.ClientResponse(http.StatusUnauthorized, "User not authorized", nil, nil)
		return
	}

	err := h.adService.Delete(ctx.Request.Context(), adID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to delete ad", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "ad deleted successfully", nil, nil)
	ctx.JSON(http.StatusOK, successRes)
}

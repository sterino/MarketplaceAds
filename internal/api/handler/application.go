package handler

import (
	"Marketplace/internal/domain/application"
	"Marketplace/internal/service/interfaces"
	"Marketplace/internal/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApplicationHandler struct {
	applicationService interfaces.ApplicationService
}

func NewApplicationHandler(service interfaces.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{
		applicationService: service,
	}
}

// CreateApplication godoc
// @Summary Create an application
// @Description Create a new application with the provided data
// @Tags application
// @Accept json
// @Produce json
// @Param application body application.CreateRequest true "Application Create Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /application/create [post]
func (h *ApplicationHandler) Create(ctx *gin.Context) {
	req := application.CreateRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid request data", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	id, err := h.applicationService.Create(ctx.Request.Context(), req)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Failed to create application", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Application created successfully", gin.H{"id": id}, nil)
	ctx.JSON(http.StatusCreated, successRes)
}

// GetApplicationByAdID godoc
// @Summary Get applications by Ad ID
// @Description Retrieve applications based on the Ad ID
// @Tags application
// @Accept json
// @Produce json
// @Param id path string true "Ad ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /application/ad/{id} [get]
func (h *ApplicationHandler) GetByAdID(ctx *gin.Context) {
	adID := ctx.Param("id")
	if adID == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "Ad ID is required", nil, nil)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	applications, err := h.applicationService.GetByAdID(ctx.Request.Context(), adID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusNotFound, "Applications not found", nil, err.Error())
		ctx.JSON(http.StatusNotFound, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Applications retrieved successfully", applications, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// GetApplicationsByInfluencer godoc
// @Summary Get applications by Influencer ID
// @Description Retrieve applications based on the Influencer ID
// @Tags application
// @Accept json
// @Produce json
// @Param id path string true "Influencer ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /application/influencer/{id} [get]
func (h *ApplicationHandler) GetByInfluencer(ctx *gin.Context) {
	influencerID := ctx.Param("id")
	if influencerID == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "Influencer ID is required", nil, nil)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	applications, err := h.applicationService.GetByInfluencerID(ctx.Request.Context(), influencerID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusNotFound, "Applications not found", nil, err.Error())
		ctx.JSON(http.StatusNotFound, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Applications retrieved successfully", applications, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// UpdateApplicationStatus godoc
// @Summary Update the status of an application
// @Description Update the status of an application by its ID
// @Tags application
// @Accept json
// @Produce json
// @Param id path string true "Application ID"
// @Param status body string true "New status"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /application/{id}/status [put]
func (h *ApplicationHandler) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "Application ID is required", nil, nil)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	var status string
	if err := ctx.BindJSON(&status); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid status data", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := h.applicationService.UpdateStatus(ctx.Request.Context(), id, status)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Failed to update status", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Application status updated successfully", nil, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// DeleteApplication godoc
// @Summary Delete an application by ID
// @Description Delete the application based on its ID
// @Tags application
// @Accept json
// @Produce json
// @Param id path string true "Application ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /application/{id}/delete [delete]
func (h *ApplicationHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "Application ID is required", nil, nil)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := h.applicationService.Delete(ctx.Request.Context(), id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Failed to delete application", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Application deleted successfully", nil, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// GetApplicationByID godoc
// @Summary Get an application by its ID
// @Description Retrieve an application based on its unique ID
// @Tags application
// @Accept json
// @Produce json
// @Param id path string true "Application ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /application/{id} [get]
func (h *ApplicationHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "Application ID is required", nil, nil)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	application, err := h.applicationService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusNotFound, "Application not found", nil, err.Error())
		ctx.JSON(http.StatusNotFound, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Application retrieved successfully", application, nil)
	ctx.JSON(http.StatusOK, successRes)
}

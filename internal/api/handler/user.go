package handler

import (
	"Marketplace/internal/service/interfaces"
	"Marketplace/internal/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

// GetAccountType godoc
// @Summary Get user account type
// @Description Retrieve account type (company/influencer) based on user ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response "Account type retrieved successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/account_type/{id} [get]
func (h *UserHandler) GetAccountType(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "user ID is required", nil, nil)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	accountType, err := h.userService.GetAccountTypeByID(ctx.Request.Context(), userID)
	if err != nil {
		if err.Error() == "user not found" {
			errRes := response.ClientResponse(http.StatusNotFound, "user not found", nil, err.Error())
			ctx.JSON(http.StatusNotFound, errRes)
		} else {
			errRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve account type", nil, err.Error())
			ctx.JSON(http.StatusInternalServerError, errRes)
		}
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "account type retrieved successfully", gin.H{"account_type": accountType}, nil)
	ctx.JSON(http.StatusOK, successRes)
}

package handler

import (
	"Marketplace/internal/domain/company"
	"Marketplace/internal/service/interfaces"
	"Marketplace/internal/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CompanyHandler struct {
	companyService interfaces.CompanyService
}

func NewCompanyHandler(service interfaces.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		companyService: service,
	}
}

// LoginCompany godoc
// @Summary Login for companies
// @Description Login for companies with the input payload
// @Tags company
// @Accept json
// @Produce json
// @Param project body company.LoginRequest true "Login Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /company/login [post]
func (h *CompanyHandler) Login(ctx *gin.Context) {
	req := company.LoginRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	token, expiresAt, err := h.companyService.Login(ctx.Request.Context(), req)
	if err != nil {
		errRes := response.ClientResponse(http.StatusForbidden, "authorization failed", nil, err.Error())
		ctx.JSON(http.StatusForbidden, errRes)
	}

	successRes := response.ClientResponse(http.StatusCreated, "authorized", gin.H{"access token": token, "expires_at": expiresAt}, nil)
	ctx.JSON(http.StatusCreated, successRes)
}

// RegisterCompany godoc
// @Summary Register for companies
// @Description Register for companies with the input payload
// @Tags company
// @Accept json
// @Produce json
// @Param project body company.RegisterRequest true "Login Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /company/register [post]
func (h *CompanyHandler) Register(ctx *gin.Context) {
	req := company.RegisterRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := h.companyService.Register(ctx.Request.Context(), req)
	if err != nil {
		errRes := response.ClientResponse(http.StatusForbidden, "registration failed", nil, err.Error())
		ctx.JSON(http.StatusForbidden, errRes)
	}

	successRes := response.ClientResponse(http.StatusCreated, "registered", res, nil)
	ctx.JSON(http.StatusCreated, successRes)
}

//func (h *CompanyHandler) GetByID(ctx *gin.Context){
//	var body
//}
//
//func (h *CompanyHandler) GetByEmail(ctx *gin.Context){
//	var body
//}

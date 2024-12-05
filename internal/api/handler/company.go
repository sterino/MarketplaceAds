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
// @Param company body company.LoginRequest true "Login Request"
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
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "authorized", gin.H{"access_token": token, "expires_at": expiresAt}, nil)
	ctx.JSON(http.StatusCreated, successRes)
}

// RegisterCompany godoc
// @Summary Register for companies
// @Description Register for companies with the input payload
// @Tags company
// @Accept json
// @Produce json
// @Param company body company.RegisterRequest true "Login Request"
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

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := h.companyService.Register(ctx.Request.Context(), req)
	if err != nil {
		errRes := response.ClientResponse(http.StatusForbidden, "registration failed", nil, err.Error())
		ctx.JSON(http.StatusForbidden, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "registered", res, nil)
	ctx.JSON(http.StatusCreated, successRes)
}

// VerifyEmail godoc
// @Summary Verify email with code for company
// @Description Verify company email with the code sent via email
// @Tags company
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param code body string true "Verification Code"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /company/verify [post]
func (h *CompanyHandler) VerifyEmail(ctx *gin.Context) {
	var req struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "invalid input", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := h.companyService.VerifyEmail(ctx.Request.Context(), req.Email, req.Code)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "verification failed", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "email verified successfully", nil, nil)
	ctx.JSON(http.StatusOK, successRes)
}

//func (h *CompanyHandler) GetByID(ctx *gin.Context){
//	var body
//}
//
//func (h *CompanyHandler) GetByEmail(ctx *gin.Context){
//	var body
//}

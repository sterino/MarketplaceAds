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

// SendCode godoc
// @Summary Send verification code to email for company
// @Description Send verification code to the email address for company registration
// @Tags company
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Success 200 {object} response.Response "Verification code sent successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /company/send_code [post]
func (h *CompanyHandler) SendCode(ctx *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	// Привязка данных из тела запроса
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// Возвращаем ошибку при неправильном вводе
		errRes := response.ClientResponse(http.StatusBadRequest, "invalid input", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	// Отправка кода верификации
	err := h.companyService.SendCode(ctx.Request.Context(), req.Email)
	if err != nil {
		// Ошибка при отправке кода
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to send verification code", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	// Успешная отправка кода
	successRes := response.ClientResponse(http.StatusOK, "verification code sent successfully", nil, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// GetCompanyByID godoc
// @Summary Get company by ID
// @Description Get company details using company ID
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /company/{id} [get]
func (h *CompanyHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "company ID is required", nil, nil)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := h.companyService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusNotFound, "company not found", nil, err.Error())
		ctx.JSON(http.StatusNotFound, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "company retrieved successfully", res, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// GetCompanyByEmail godoc
// @Summary Get company by Email
// @Description Get company details using company Email
// @Tags company
// @Accept json
// @Produce json
// @Param email path string true "Company Email"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /company/email/{email} [get]
func (h *CompanyHandler) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "company email is required", nil, nil)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := h.companyService.GetByEmail(ctx.Request.Context(), email)
	if err != nil {
		errRes := response.ClientResponse(http.StatusNotFound, "company not found", nil, err.Error())
		ctx.JSON(http.StatusNotFound, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "company retrieved successfully", res, nil)
	ctx.JSON(http.StatusOK, successRes)
}

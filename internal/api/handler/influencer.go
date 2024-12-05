package handler

import (
	"Marketplace/internal/domain/influencer"
	"Marketplace/internal/service/interfaces"
	"Marketplace/internal/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfluencerHandler struct {
	influencerService interfaces.InfluencerService
}

func NewInfluencerHandler(service interfaces.InfluencerService) *InfluencerHandler {
	return &InfluencerHandler{
		influencerService: service,
	}
}

// LoginInfluencer godoc
// @Summary Login for influencers
// @Description Login for influencers with the input payload
// @Tags influencer
// @Accept json
// @Produce json
// @Param influencer body influencer.LoginRequest true "Login Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /influencer/login [post]
func (h *InfluencerHandler) Login(ctx *gin.Context) {
	req := influencer.LoginRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	token, expiresAt, err := h.influencerService.Login(ctx.Request.Context(), req)
	if err != nil {
		errRes := response.ClientResponse(http.StatusForbidden, "authorization failed", nil, err.Error())
		ctx.JSON(http.StatusForbidden, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "authorized", gin.H{"access_token": token, "expires_at": expiresAt}, nil)
	ctx.JSON(http.StatusCreated, successRes)
}

// RegisterInfluencer godoc
// @Summary Register for influencers
// @Description Register for influencers with the input payload
// @Tags influencer
// @Accept json
// @Produce json
// @Param influencer body influencer.RegisterRequest true "Register Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /influencer/register [post]
func (h *InfluencerHandler) Register(ctx *gin.Context) {
	req := influencer.RegisterRequest{}
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

	res, err := h.influencerService.Register(ctx.Request.Context(), req)
	if err != nil {
		errRes := response.ClientResponse(http.StatusForbidden, "registration failed", nil, err.Error())
		ctx.JSON(http.StatusForbidden, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "registered", gin.H{"id": res}, nil)
	ctx.JSON(http.StatusCreated, successRes)
}

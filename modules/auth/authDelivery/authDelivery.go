package authDelivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tsileo/defender"
	"net/http"
	"tinderleaked/middleware"
	authModel "tinderleaked/models/auth"
	responseModel "tinderleaked/models/response"
	"tinderleaked/modules/auth"
)

type authDelivery struct {
	d           *defender.Defender
	authUsecase auth.UsecaseAuth
}

func NewAuthHTTPHandler(r *gin.Engine, def *defender.Defender, authUsecase auth.UsecaseAuth) {
	handler := authDelivery{authUsecase: authUsecase, d: def}

	auth := r.Group("/auth")
	auth.POST("/register", middleware.BasicAuth, handler.AuthRegister)
	auth.POST("/login", middleware.BasicAuth, handler.AuthLogin)
}

func (handler *authDelivery) AuthRegister(c *gin.Context) {
	moduleName := "AuthRegister"
	var req authModel.RegisterRequest

	errBind := c.BindJSON(&req)
	if errBind != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errBind.Error()).Error())
		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: responseModel.FailedHttpStatus,
			Data:            nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	errSubmit := handler.authUsecase.SubmitRegisterAccount(&req)
	if errSubmit != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errSubmit.Error()).Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": "Failed",
			},
		)
		return
	}

	response := responseModel.JSONResponse{
		ResponseCode:    responseModel.SuccessHttpCode,
		ResponseMessage: responseModel.SuccessHttpStatus,
		Data:            &[]string{},
	}

	c.JSON(http.StatusOK, response)
	return
}

func (handler *authDelivery) AuthLogin(c *gin.Context) {
	moduleName := "AuthLogin"
	var req authModel.LoginRequest

	errBind := c.BindJSON(&req)
	if errBind != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errBind.Error()).Error())
		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: responseModel.FailedHttpStatus,
			Data:            nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data, errSubmit := handler.authUsecase.Login(&req)
	if errSubmit != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errSubmit.Error()).Error())

		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: responseModel.FailedHttpStatus,
			Data:            nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responseModel.JSONResponse{
		ResponseCode:    responseModel.SuccessHttpCode,
		ResponseMessage: responseModel.SuccessHttpStatus,
		Data:            data,
	}

	c.JSON(http.StatusOK, response)
	return
}

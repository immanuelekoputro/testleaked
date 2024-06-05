package usersDelivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"tinderleaked/middleware"
	responseModel "tinderleaked/models/response"
	usersModel "tinderleaked/models/users"
	"tinderleaked/modules/users"
	"tinderleaked/tools/jwt"
)

type usersDelivery struct {
	usersUsecase users.UsecaseUsers
}

func NewUsersHTTPHandler(r *gin.Engine, usersUsecase users.UsecaseUsers) {
	handler := usersDelivery{usersUsecase: usersUsecase}

	users := r.Group("/users")
	users.Use(middleware.JWTAuth)
	{
		users.GET("/viewable-user", handler.GetViewableUser)
		users.PUT("/action", handler.UpdateVisitor)
	}
}

func (handler *usersDelivery) UpdateVisitor(c *gin.Context) {
	moduleName := "UpdateVisitor"

	var req usersModel.SubmitActionUser

	errBind := c.BindJSON(&req)
	if errBind != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errBind.Error()).Error())
		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: errBind.Error(),
			Data:            nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	getUserIDFromToken, errGetUserIDFromToken := jwt.GetUserIDFromToken(c)
	if errGetUserIDFromToken != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errGetUserIDFromToken.Error()).Error())
		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: responseModel.FailedHttpStatus,
			Data:            nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	log.Debug().Msg(getUserIDFromToken)
	userIDConverted, errUserIDConverted := strconv.Atoi(getUserIDFromToken)
	if errUserIDConverted != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errUserIDConverted.Error()).Error())
		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: responseModel.FailedHttpStatus,
			Data:            nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	errSubmitAction := handler.usersUsecase.SubmitActionUser(userIDConverted, req)
	if errSubmitAction != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errSubmitAction.Error()).Error())
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
		Data:            &[]string{},
	}

	c.JSON(http.StatusOK, response)
	return
}

func (handler *usersDelivery) GetViewableUser(c *gin.Context) {
	moduleName := "AssignPackageToCustomer"

	getUserIDFromToken, errGetUserIDFromToken := jwt.GetUserIDFromToken(c)
	if errGetUserIDFromToken != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errGetUserIDFromToken.Error()).Error())
		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: responseModel.FailedHttpStatus,
			Data:            nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	log.Debug().Msg(getUserIDFromToken)
	userIDConverted, errUserIDConverted := strconv.Atoi(getUserIDFromToken)
	if errUserIDConverted != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errUserIDConverted.Error()).Error())
		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: responseModel.FailedHttpStatus,
			Data:            nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userViewable, errUserViewable := handler.usersUsecase.ListAnotherUsers(userIDConverted)
	if errUserViewable != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errUserViewable.Error()).Error())
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
		Data:            userViewable,
	}

	c.JSON(http.StatusOK, response)
	return
}

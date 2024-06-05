package packagesDelivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"tinderleaked/middleware"
	packagesModel "tinderleaked/models/packages"
	responseModel "tinderleaked/models/response"
	"tinderleaked/modules/packages"
	"tinderleaked/tools/jwt"
)

type packagesDelivery struct {
	packagesUsecase packages.UsecasePackages
}

func NewPackagesHTTPHandler(r *gin.Engine, packagesUsecase packages.UsecasePackages) {
	handler := packagesDelivery{packagesUsecase: packagesUsecase}

	packages := r.Group("/packages")
	packages.Use(middleware.JWTAuth)
	{
		packages.GET("/all", handler.GetAllPackage)
		packages.GET("/my", handler.GetUserLoginPackage)
		packages.POST("/buy", handler.AssignPackageToCustomer)
	}
}

func (handler *packagesDelivery) GetAllPackage(c *gin.Context) {
	moduleName := "GetAllPackage"

	allPackages, errAllPackages := handler.packagesUsecase.GetPackages()
	if errAllPackages != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errAllPackages.Error()).Error())
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
		Data:            allPackages,
	}

	c.JSON(http.StatusOK, response)
	return
}

func (handler *packagesDelivery) GetUserLoginPackage(c *gin.Context) {
	moduleName := "GetUserLoginPackage"
	getUserIDFromToken, errGetUserIDFromToken := jwt.GetUserIDFromToken(c)
	if errGetUserIDFromToken != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errGetUserIDFromToken.Error()).Error())

		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: errGetUserIDFromToken.Error(),
			Data:            &[]string{},
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
			ResponseMessage: errUserIDConverted.Error(),
			Data:            &[]string{},
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	myPackage, errMyPackage := handler.packagesUsecase.GetLoggedUserPackage(userIDConverted)
	if errMyPackage != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errMyPackage.Error()).Error())
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
		Data:            myPackage,
	}

	c.JSON(http.StatusOK, response)
	return
}

func (handler *packagesDelivery) AssignPackageToCustomer(c *gin.Context) {
	moduleName := "GetAllPackage"

	var req packagesModel.AssignUserPackage

	errBind := c.BindJSON(&req)
	if errBind != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errBind.Error()).Error())
		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: errBind.Error(),
			Data:            &[]string{},
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	getUserIDFromToken, errGetUserIDFromToken := jwt.GetUserIDFromToken(c)
	if errGetUserIDFromToken != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errGetUserIDFromToken.Error()).Error())

		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: errGetUserIDFromToken.Error(),
			Data:            &[]string{},
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
			ResponseMessage: errUserIDConverted.Error(),
			Data:            &[]string{},
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	errAllPackages := handler.packagesUsecase.AssignPackageByLoggedUser(userIDConverted, int(req.PackageID))
	if errAllPackages != nil {
		log.Error().Msg(errors.New(moduleName + " | " + errAllPackages.Error()).Error())

		response := responseModel.JSONResponse{
			ResponseCode:    responseModel.FailedHttpCode,
			ResponseMessage: errAllPackages.Error(),
			Data:            &[]string{},
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

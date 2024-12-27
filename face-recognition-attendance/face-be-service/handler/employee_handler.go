package handler

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"face-be-service/common/constants"
	"face-be-service/common/utils"
	"face-be-service/handler/models"
)

func (h *Handler) Insert(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "Insert"))
	request := models.ImageInsertRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "ShouldBind", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err = request.Validate(); err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Validate", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	imgFile, err := request.Image.Open()
	defer func(file multipart.File) {
		utils.CloseFile(file)
	}(imgFile)

	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "request.Image.Open", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.employeeService.Insert(ctx, request.ToImageInsertInput(imgFile, request.Image.Filename,
		request.Image.Size))
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Insert", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) Search(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "Search"))
	request := models.ImageSearchRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "ShouldBind", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err = request.Validate(); err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Validate", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	imgFile, err := request.Image.Open()
	defer func(file multipart.File) {
		utils.CloseFile(file)
	}(imgFile)

	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "request.Image.Open", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.employeeService.Search(ctx, request.ToImageSearchInput(imgFile, request.Image.Filename,
		request.Image.Size))
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Search", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

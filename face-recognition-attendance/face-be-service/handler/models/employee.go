package models

import (
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"face-be-service/internal/domains/employee/models"
)

type ImageSearchRequest struct {
	Image *multipart.FileHeader `form:"image"`
}

type ImageInsertRequest struct {
	Image        *multipart.FileHeader `form:"image"`
	EmployeeID   string                `form:"employee_id"`
	EmployeeName string                `form:"employee_name"`
}

func (r *ImageSearchRequest) ToImageSearchInput(imgFile multipart.File, imgName string,
	imgSize int64) *models.ImageSearchInput {
	if r == nil {
		return nil
	}

	return &models.ImageSearchInput{
		ImageFile: imgFile,
		ImageName: imgName,
		ImageSize: imgSize,
	}
}

func (r *ImageInsertRequest) ToImageInsertInput(imgFile multipart.File, imgName string,
	imgSize int64) *models.ImageInsertInput {
	if r == nil {
		return nil
	}

	return &models.ImageInsertInput{
		ImageFile:    imgFile,
		ImageName:    imgName,
		ImageSize:    imgSize,
		EmployeeID:   r.EmployeeID,
		EmployeeName: r.EmployeeName,
	}
}

func (r *ImageSearchRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Image, validation.Required),
	)
}

func (r *ImageInsertRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Image, validation.Required),
		validation.Field(&r.EmployeeID, validation.Required),
		validation.Field(&r.EmployeeName, validation.Required),
	)
}

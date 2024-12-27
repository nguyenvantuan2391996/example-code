package models

import (
	"mime/multipart"
)

type ImageSearchInput struct {
	ImageFile   multipart.File
	ImageName   string
	ImageSize   int64
	NumberOfIDs int
}

type ImageInsertInput struct {
	ImageFile    multipart.File
	ImageName    string
	EmployeeName string
	EmployeeID   string
	ImageSize    int64
}

type ImageListInput struct {
	IDs []string
}

type ProductOutput struct {
	ImagePath  string `json:"image_path"`
	EmployeeID int64  `json:"employee_id"`
}

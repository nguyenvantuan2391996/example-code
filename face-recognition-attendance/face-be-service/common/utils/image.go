package utils

import (
	"mime/multipart"

	"github.com/sirupsen/logrus"

	"face-be-service/common/constants"
)

func CloseFile(file multipart.File) {
	if file == nil {
		return
	}

	errClose := file.Close()
	if errClose != nil {
		logrus.Errorf(constants.FormatTaskErr, "file.Close", errClose)
	}
}

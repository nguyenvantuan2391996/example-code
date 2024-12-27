package utils

import (
	"io"
	"mime/multipart"

	"github.com/sirupsen/logrus"
)

func CreateFormFile(w *multipart.Writer, rd io.Reader, fieldName, fileName string) (*multipart.Writer, error) {
	formWriter, err := w.CreateFormFile(fieldName, fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(formWriter, rd)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func CloseResponse(body io.ReadCloser) {
	if body == nil {
		return
	}

	errClose := body.Close()
	if errClose != nil {
		logrus.Error(errClose.Error())
	}
}

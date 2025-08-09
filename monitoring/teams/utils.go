package teams

import (
	"io"

	"github.com/sirupsen/logrus"
)

func CloseResponse(body io.ReadCloser) {
	if body == nil {
		return
	}

	errClose := body.Close()
	if errClose != nil {
		logrus.Error(errClose.Error())
	}
}

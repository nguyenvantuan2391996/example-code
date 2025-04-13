package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/sirupsen/logrus"
)

func CloseFile(file multipart.File) {
	if file == nil {
		return
	}

	err := file.Close()
	if err != nil {
		logrus.Error(err)
	}
}

func ConvertImageBase64ToFile(imgBase64 string) (io.Reader, error) {
	parts := strings.SplitN(imgBase64, ",", 2)
	if len(parts) == 2 {
		imgBase64 = parts[1]
	}

	// Decode the base64 data
	decodedData, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 data: %v", err)
	}

	return bytes.NewReader(decodedData), nil
}

func CalcBase64LengthInByte(imgBase64 string) int64 {
	parts := strings.SplitN(imgBase64, ",", 2)
	if len(parts) == 2 {
		imgBase64 = parts[1]
	}

	l := len(imgBase64)

	// count how many trailing '=' there are (if any)
	eq := 0
	if l >= 2 {
		if imgBase64[l-1] == '=' {
			eq++
		}
		if imgBase64[l-2] == '=' {
			eq++
		}

		l -= eq
	}

	// basically:
	// eq == 0 :	bits-wasted = 0
	// eq == 1 :	bits-wasted = 2
	// eq == 2 :	bits-wasted = 4

	// so orig length ==  (l*6 - eq*2) / 8

	return int64((l*3 - eq) / 4)
}

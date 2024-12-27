package third_party

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"face-be-service/common/constants"
	"face-be-service/common/utils"
)

const (
	ExtractionPath = "api/v1/extract/"
)

var faceExtractionClient *FaceExtractionClient

type FaceExtractionClient struct {
	client   *http.Client
	endpoint string
	apiKey   string
}

func NewFaceExtractionClient() {
	if faceExtractionClient == nil {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client := &http.Client{
			Timeout:   60 * time.Second,
			Transport: tr,
		}
		faceExtractionClient = &FaceExtractionClient{
			client:   client,
			endpoint: viper.GetString(constants.FaceExtractionEndpoint),
			apiKey:   viper.GetString(constants.FaceExtractionAPIKey),
		}
	}
	logrus.Info(fmt.Sprintf("completed initializing client at %s", faceExtractionClient.endpoint))
}

func GetInstance() *FaceExtractionClient {
	return faceExtractionClient
}

func (avc *FaceExtractionClient) ExtractImage(req *ImageExtractionRequest) (*ImageExtractionResponse, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	w, err := utils.CreateFormFile(w, req.Image, "file", req.FileName)
	if err != nil {
		return nil, err
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	extractionReq, err := http.NewRequest(http.MethodPost, strings.Join([]string{avc.endpoint, ExtractionPath},
		"/"), &buf)
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "http.NewRequest", err)
		return nil, err
	}

	extractionReq.Header.Set(constants.ContentTypeHeader, w.FormDataContentType())
	extractionReq.Header.Set(constants.XAPIKeyHeader, avc.apiKey)

	resp, err := avc.client.Do(extractionReq)
	if resp != nil {
		defer func(body io.ReadCloser) {
			utils.CloseResponse(body)
		}(resp.Body)
	}
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "avc.client.Do", err)
		return nil, err
	}

	bBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("image extraction request failed | status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("image extraction request failed")
	}

	var respBody ImageExtractionResponse
	err = json.Unmarshal(bBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

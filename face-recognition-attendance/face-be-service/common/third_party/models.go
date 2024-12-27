package third_party

import "io"

type ImageExtractionRequest struct {
	Image    io.Reader
	FileName string
}

type ImageExtractionResponse struct {
	Vector []float32 `json:"vector"`
}

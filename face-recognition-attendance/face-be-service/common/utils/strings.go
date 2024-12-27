package utils

import "encoding/json"

func ConvertArrayFloat32(arr []float32) (string, error) {
	marshal, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

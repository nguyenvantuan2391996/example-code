import "encoding/base64"

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	return string(data), err
}

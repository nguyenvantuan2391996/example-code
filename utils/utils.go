package utils

import (
	"fmt"
	"regexp"
	"json"
)

func MakeLockKey(category, identifier string) string {
	return fmt.Sprintf("%s:%s", category, identifier)
}

func SpanName(action, entity, scope string) string {
	return fmt.Sprintf("%s_%s_%s", action, entity, scope)
}

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$`
	
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

func SafeMarshalJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal struct to json: %w", err)
	}
	return jsonData, nil
}

func UnmarshalJSONBytes(data []byte, target interface{}) error {
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal json to struct: %w", err)
	}
	return nil
}
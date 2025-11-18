package utils

import (
	"fmt"
	"regexp"
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
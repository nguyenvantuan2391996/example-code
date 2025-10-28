package utils

import (
	"fmt"
)

func MakeLockKey(category, identifier string) string {
	return fmt.Sprintf("%s:%s", category, identifier)
}

func SpanName(action, entity, scope string) string {
	return fmt.Sprintf("%s_%s_%s", action, entity, scope)
}

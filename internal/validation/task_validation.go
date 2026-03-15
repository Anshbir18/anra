package validation

import (
	"errors"
	"strings"
)

func ValidateTitle(title string) error {
	title = strings.TrimSpace(title)

	if title == "" {
		return errors.New("title is required")
	}

	if len(title) > 200 {
		return errors.New("title must be less than 200 characters")
	}

	return nil
}

func ValidateStatus(status string) bool {
	switch status {
	case "todo", "in_progress", "done":
		return true
	default:
		return false
	}
}
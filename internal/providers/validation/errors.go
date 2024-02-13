package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "The field ist required",
		"email":    "The field must have a valid email address",
		"min":      "Should be more than limit",
		"max":      "Should be less than limit",
	}
}

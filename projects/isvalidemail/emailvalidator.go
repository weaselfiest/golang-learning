package emailvalidator

import "strings"

func IsValidEmail(email string) bool {
	if email == "" || strings.Contains(email, " ") {
		return false
	}

	local, domain, ok := strings.Cut(email, "@")
	if !ok || local == "" || domain == "" || strings.Contains(domain, "@") {
		return false
	}

	dotIndex := strings.Index(domain, ".")
	dotLastIndex := strings.LastIndex(domain, ".")

	return dotIndex > 0 && dotLastIndex < len(domain)-1
}

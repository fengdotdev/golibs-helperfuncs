package validators

import "strings"

func EmailValidator(email string) (bool, error) {
	// Check if the email contains "@" and "."
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return false, nil
	}

	// Split the email into local part and domain part
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false, nil
	}

	localPart := parts[0]
	domainPart := parts[1]

	// Check if the local part is empty
	if len(localPart) == 0 {
		return false, nil
	}

	// Check if the domain part contains "."
	if !strings.Contains(domainPart, ".") {
		return false, nil
	}

	return true, nil
}

package validators

func PhoneNumberValidator(phoneNumber string) (bool, error) {
	// Check if the phone number is empty
	if len(phoneNumber) == 0 {
		return false, nil
	}

	// Check if the phone number contains only digits and optional "+" at the beginning
	for i, c := range phoneNumber {
		if i == 0 && c == '+' {
			continue
		}
		if c < '0' || c > '9' {
			return false, nil
		}
	}

	return true, nil
}

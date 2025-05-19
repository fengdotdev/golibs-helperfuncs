package secret

func AssertAESGCM_RequirementsOrErr(key, iv, plaintext, AdditionalData []byte) error {

	err := AssertKeyOrErr(key)
	if err != nil {
		return err
	}
	err = AssertIVOrErr(iv)
	if err != nil {
		return err
	}

	err = AssertPlaintextOrErr(plaintext)
	if err != nil {
		return err
	}

	err = AssertAdditionalDataOrErr(AdditionalData)
	if err != nil {
		return err
	}

	return nil
}

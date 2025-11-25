package generator

func MutatePassword(password string) string {
	var substitutionMap = map[string]string{
		"a": "@",
		"i": "1",
		"s": "$",
		"o": "0",
		"A": "4",
		"I": "!",
		"S": "5",
		"O": "0",
	}

	byteRepresentedPassword := []byte(password)
	for i := range byteRepresentedPassword {
		if substitution, exists := substitutionMap[string(byteRepresentedPassword[i])]; exists {
			byteRepresentedPassword[i] = substitution[0]
		}
	}

	if len(byteRepresentedPassword) > 3 {
		positionToSwap := len(byteRepresentedPassword) / 2
		temp := byteRepresentedPassword[positionToSwap]
		byteRepresentedPassword[positionToSwap] = byteRepresentedPassword[positionToSwap-1]
		byteRepresentedPassword[positionToSwap-1] = temp
	}

	randomSalt := "28"
	byteRepresentedPassword = append(byteRepresentedPassword, []byte(randomSalt)...)

	flipStep := 2
	for i := range byteRepresentedPassword {
		if i%flipStep == 0 {
			byteRepresentedPassword[i] = byteRepresentedPassword[i] ^ 0x20
		}
	}

	return string(byteRepresentedPassword)
}

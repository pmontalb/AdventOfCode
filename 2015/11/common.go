package main

func IsValid(password string) bool {
	if len(password) < 3 || len(password) > 8 {
		return false
	}
	containsIncreasingStraight := false
	containsBlacklistedLetters := false

	nPairs := 0
	for i := 0; i < len(password); i++ {
		if i < len(password)-1 {
			if i < len(password)-3 {
				containsIncreasingStraight = containsIncreasingStraight || (password[i+1] == password[i]+1 && password[i+2] == password[i]+2)
			}

			if password[i] == password[i+1] {
				if i < len(password)-2 {
					if password[i+2] != password[i] {
						nPairs++
					}
				} else {
					nPairs++
				}
			}
		}
		containsBlacklistedLetters = containsBlacklistedLetters || (password[i] == 'i' || password[i] == 'o' || password[i] == 'l')
	}

	return containsIncreasingStraight && !containsBlacklistedLetters && nPairs >= 2
}

func NextChar(str string) string {
	byteStr := []byte(str)
	carry := false

	i := len(byteStr) - 1
	if byteStr[i] >= 'z' {
		byteStr[i] = 'a'
		carry = true
	} else {
		byteStr[i]++
	}

	if carry {
		for i >= 0 {
			i--
			if byteStr[i] >= 'z' {
				byteStr[i] = 'a'
			} else {
				byteStr[i]++
				break
			}
		}
	}
	return string(byteStr)
}

func FindNextPassword(password string) string {
	password = NextChar(password)

	// todo: better logic here?
	for !IsValid(password) {
		password = NextChar(password)
	}
	return password
}

package utils

import (
	"regexp"
	"strings"
	"xuoxod/adminhelper/internal/constants"
)

func IsLettersOnly(arg string) bool {
	re := regexp.MustCompile(constants.PatternAlphaChars)
	return re.MatchString(strings.TrimSpace(arg))
}

func IsLettersAndSpace(arg string) bool {
	re := regexp.MustCompile(constants.PatternAlphaAndSpaceChars)
	return re.MatchString(strings.TrimSpace(arg))
}

func IsNumbersOnly(arg string) bool {
	re := regexp.MustCompile(constants.PatternNumbers)
	return re.MatchString(strings.TrimSpace(arg))
}

func IsDecimalsAndOrNumbersOnly(arg string) bool {
	re := regexp.MustCompile(constants.PatternNumbersAndOrDecimals)
	return re.MatchString(strings.TrimSpace(arg))
}

func IsAlphaNumeric(arg string) bool {
	re := regexp.MustCompile(constants.PatternAlphaNumeric)
	return re.MatchString(strings.TrimSpace(arg))
}

func IsNonAlphaNumeric(arg string) bool {
	re := regexp.MustCompile(constants.PatternNonAlphaNumeric)
	return re.MatchString(strings.TrimSpace(arg))
}

func IsValidEmail(arg string) bool {
	re := regexp.MustCompile(constants.PatternEmail)
	return re.MatchString(strings.TrimSpace(arg))
}

func IsValidPhoneNumber(arg string) bool {
	re := regexp.MustCompile(constants.PatternPhoneNumber)
	return re.MatchString(strings.TrimSpace(arg))
}

func IsValidAge(arg string) bool {
	re := regexp.MustCompile(constants.PatternAdultAge)
	return re.MatchString(strings.TrimSpace(arg))
}

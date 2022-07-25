package helper

import (
	"regexp"
	"strings"
)

var (
	// -- NOTE: Below regex pattern can be found here: https://www.golangprograms.com/regular-expression-to-validate-phone-number.html
	regexPhone = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	// -- NOTE: Below regex pattern can be found here: https://stackoverflow.com/questions/2113908/what-regular-expression-will-match-valid-international-phone-numbers
	// -- not support country_code+number without '+' prefix
	regexAreaCode  = regexp.MustCompile(`^\+((?:9[679]|8[035789]|6[789]|5[90]|42|3[578]|2[1-689])|9[0-58]|8[1246]|6[0-6]|5[1-8]|4[013-9]|3[0-469]|2[70]|7|1)`)
	regexNonNumber = regexp.MustCompile(`[^0-9]`)

	// not alphanum chars
	regexNonAlphaNum      = regexp.MustCompile("[^a-zA-Z0-9]+")
	regexNonAlphaNumSpace = regexp.MustCompile("[^a-zA-Z0-9 ]+")
)

func StrLimit(str *string, limit int) *string {
	if nil == str {
		return nil
	}
	if limit < 0 {
		empty := ""
		return &empty
	}
	newStr := *str
	if len(newStr) > limit {
		newStr = newStr[0:limit]
		return &newStr
	}
	return &newStr
}

// StrToLower returns `str` with all Unicode letters mapped to their
// lower case.
//
// https://golang.org/pkg/strings/#ToLower
func StrToLower(str string) string {
	return strings.ToLower(str)
}

func StrReplace(str string, before string, after string) string {
	return strings.ReplaceAll(str, before, after)
}

func IsStringPhoneNumber(contact string) bool {
	return regexPhone.MatchString(contact)
}

func NormalizePhoneNumber(phone string) string {

	// -- remove area code
	compiledAreaCode := string(regexAreaCode.ReplaceAll([]byte(phone), []byte("")))

	// -- remove non number characters (dots, dash, space, a-z, A-Z, and special characters)
	compiledNonNumber := string(regexNonNumber.ReplaceAll([]byte(compiledAreaCode), []byte("")))

	// -- trim 0 as the first character of phone number
	trimmedPhone := strings.TrimLeft(compiledNonNumber, "0")

	return trimmedPhone
}

func NormalizeEmail(email string) string {
	replacedSpace := strings.ReplaceAll(email, " ", "")

	return strings.ToLower(replacedSpace)
}

func IsEmptyTrimString(str *string) bool {
	return str == nil || strings.TrimSpace(*str) == ""
}

// TrimAndLowerString will return whitespace-trimmed and lowered-case version of the given str
func TrimAndLowerString(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

// RemoveNonAlphaNumChar will return string with removed non alphanumeric chars or empty string if nil was given
func RemoveNonAlphaNumChar(str string) string {
	return string(regexNonAlphaNum.ReplaceAll([]byte(str), []byte("")))
}

func RemoveNonAlphaNumSpaceChar(str string) string {
	return string(regexNonAlphaNumSpace.ReplaceAll([]byte(str), []byte("")))
}

// EmptyStringWhenNil will return empty string when arg is nil, or arg when not nil
func EmptyStringWhenNil(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

// NormalizeSimple will remove all non alphanumeric chars, lower, and trim given string
// will return empty string if str is nil
func NormalizeSimple(str *string) string {
	return TrimAndLowerString(RemoveNonAlphaNumChar(EmptyStringWhenNil(str)))
}

// keep space intact
func NormalizeSimpleKeepSpace(str *string) string {
	return TrimAndLowerString(RemoveNonAlphaNumSpaceChar(EmptyStringWhenNil(str)))
}

func IsStringInArray(str string, arr []string) bool {

	var isInArray = true
	if str == "" {
		return !isInArray
	}

	for _, item := range arr {
		if item == str {
			return isInArray
		}
	}
	return !isInArray
}

// IsStringContainsPattern will check based on regex
func IsStringContainsPattern(str string, pattern string) bool {
	result, _ := regexp.MatchString(pattern, str)
	return result
}

func MaskPhone(str string) string {
	strLen := len(str)
	if strLen <= 4 {
		return str
	}
	return str[0:strLen-4] + "****"
}

func MaskEmail(str string) string {
	atIdx := strings.Index(str, "@")
	if atIdx == -1 {
		return "****@****.***"
	}
	return "****" + str[atIdx:]
}

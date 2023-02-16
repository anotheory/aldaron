package ReadData

import (
	"aldaron/Constant"
	"regexp"
	"strings"
)

func ValidateQueryPlaceholder(queryString string) {
	if !strings.Contains(queryString, Constant.PLACEHOLDER_COLUMNS) {
		panic(Constant.ERROR_MESSAGE.NotFoundColumn)
	} else if !strings.Contains(queryString, Constant.PLACEHOLDER_VALUES) {
		panic(Constant.ERROR_MESSAGE.NotFoundValue)
	}
}

func ValidateValueQueryPlaceholder(valueQueryString string) {
	if matched, _ := regexp.MatchString(Constant.REGEX_VALUES, valueQueryString); !matched {
		panic(Constant.ERROR_MESSAGE.ZeroValuePlaceholder)
	}
}

package logic

import (
	"aldaron/constant"
	"regexp"
	"strings"
)

type ReadDataLogic struct{}

func (logic ReadDataLogic) ValidateQueryPlaceholder(queryString string) {
	if !strings.Contains(queryString, constant.PLACEHOLDER_COLUMNS) {
		panic(constant.ERROR_MESSAGE.NotFoundColumn)
	} else if !strings.Contains(queryString, constant.PLACEHOLDER_VALUES) {
		panic(constant.ERROR_MESSAGE.NotFoundValue)
	}
}

func (logic ReadDataLogic) ValidateValueQueryPlaceholder(valueQueryString string) {
	if matched, _ := regexp.MatchString(constant.REGEX_VALUES, valueQueryString); !matched {
		panic(constant.ERROR_MESSAGE.ZeroValuePlaceholder)
	}
}

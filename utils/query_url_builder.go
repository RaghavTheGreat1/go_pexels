package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

func BuildQueryParams(params interface{}) (urlValues url.Values) {

	urlValues = url.Values{}

	value := reflect.ValueOf(params)
	valueType := value.Type()
	// stores number of fields
	numFields := value.NumField()

	for i := 0; i < numFields; i++ {
		if !value.Field(i).IsZero() {
			urlValues.Set(toSnakeCase(valueType.Field(i).Name), fmt.Sprint(value.Field(i)))
		}

	}
	return urlValues
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

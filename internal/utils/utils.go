package utils

import (
	"regexp"
	"sync"

	v "gopkg.in/go-playground/validator.v9"
)

var (
	correctSlug *regexp.Regexp
	once        sync.Once
)

// ValidateModels ..
func ValidateModels(mod interface{}) error {
	err := v.New().Struct(mod)

	return err
}

// PanicErr ...
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

package inputvar

import (
	"errors"
	"regexp"
	"strings"
)

type Input struct {
	Message    string
	Default    string
	Validation regexp.Regexp
	Value      string
}

func (i *Input) SetValue(value string) error {
	trimmedValue := strings.Trim(value, " ")

	if !i.Validation.Match([]byte(trimmedValue)) {
		return errors.New("invalid input value")
	}

	i.Value = trimmedValue

	return nil
}

func (i *Input) IsEmpty() bool {
	return i.Value == ""
}

type Options struct {
	Message    string
	Default    string
	Validation regexp.Regexp
}

func Create(inputCreation Options) (*Input, error) {
	var input = &Input{
		Message:    inputCreation.Message,
		Default:    inputCreation.Default,
		Validation: inputCreation.Validation,
	}

	var err error = nil
	if input.Default != "" {
		err = input.SetValue(inputCreation.Default)
	}

	return input, err
}

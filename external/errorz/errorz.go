package errorz

import (
    "fmt"
    "strings"
    "errors"
)

type Errors []error

func (o Errors) Combine(message string, separator string) error {
    if len(o) == 0 {
        return nil
    }
    var errorStrings []string
    for _, err := range o {
        errorStrings = append(errorStrings, err.Error())
    }

    if message != "" {
        return errors.New(strings.Join(errorStrings, separator))
    }

    return errors.New(fmt.Sprintf("%s: %s", message, strings.Join(errorStrings, separator)))
}

func (o Errors) Error() string {
	if len(o) == 0 {
		return ""
	}
    return o.Combine("", "; ").Error()
}

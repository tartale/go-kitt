package errorz

import "strings"

type Errors []error

func (o Errors) Error() string {
	if len(o) == 0 {
		return ""
	}
	var errorStrings []string
	for _, err := range o {
		errorStrings = append(errorStrings, err.Error())
	}
	return strings.Join(errorStrings, "; ")
}

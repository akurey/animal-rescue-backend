package errors

import (
	"fmt"
)

type ErrStruct struct {
	Url     string
	Code    int
	Message string
}

func (eaci ErrStruct) Error() string {
	return fmt.Sprintf(
		"Invalid resource - url: '%s', response code: %d, response: '%s'",
		eaci.Url,
		eaci.Code,
		eaci.Message,
	)
}

func CodeRunner(url string, succeed bool) (string, error) {
	if succeed {
		return "Success!", nil
	}
	return "", ErrStruct{url, 400, "authorization code expired"}
}

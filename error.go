package cfclient

//go:generate go run gen_error.go

import "fmt"

type CloudFoundryErrors struct {
	Errors []CloudFoundryError `json:"errors"`
}

func (cfErrs CloudFoundryErrors) Error() string {
	err := ""

	for _, cfErr := range cfErrs.Errors {
		err += fmt.Sprintf("%s\n", cfErr)
	}

	return err
}

type CloudFoundryError struct {
	Code        int    `json:"code"`
	ErrorCode   string `json:"error_code"`
	Description string `json:"description"`
	RawBody     []byte `json:"-"`
}

func (cfErr CloudFoundryError) Error() string {
	if cfErr.Code > 0 {
		return fmt.Sprintf("cfclient: error (%d): %s", cfErr.Code, cfErr.ErrorCode)
	}
	return fmt.Sprintf("cfclient: null error; raw error body was: %s", string(cfErr.RawBody))
}

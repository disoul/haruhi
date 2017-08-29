package err

import (
	"fmt"
)

// HaruhiErrorCode error code for HaruhiError
type HaruhiErrorCode uint16

const (
	UNKNOWN_ERROR HaruhiErrorCode = 0001

	JSON_DECODE_ERROR HaruhiErrorCode = 0002
	UNEXPECT_REGISTER HaruhiErrorCode = 0003
)

// HaruhiError error contaienr for haruhi
type HaruhiError struct {
	Error     error
	ErrorMsg  string
	ErrorCode HaruhiErrorCode
}

func (e HaruhiError) String() string {
	return fmt.Sprintf(
		"HaruhiError: %s\nErrorCode: %v\nError: %s",
		e.ErrorMsg,
		e.ErrorCode,
		e.Error,
	)
}

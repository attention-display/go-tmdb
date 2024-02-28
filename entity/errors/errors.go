package errors

import (
	"fmt"
)

func InvalidParamsErr(msg string) error {
	return fmt.Errorf("InvalidParamsErr: %s", msg)
}

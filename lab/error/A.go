package error_t

import (
	"github.com/pkg/errors"
)

func AFunc() error {
	if err := BFunc(); err != nil {
		return errors.Wrap(err, "AFunc Call BFunc Fail")
	}
	return nil
}

func BFunc() error {
	if err := CFunc(); err != nil {
		return errors.Wrap(err, "BFuncCallCFuncFail")
	}
	return nil
}

func CFunc() error {
	return errors.New("Nil Return ")
}

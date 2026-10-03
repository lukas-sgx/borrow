package utils

import "errors"

func ObjectNil(param any) error {
	if param == nil {
		return errors.New("object: nil")
	}

	return nil
}

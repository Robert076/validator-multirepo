package validator

import "errors"

func IsNameValid(name string) (bool, error) {
	if name == "" {
		return false, errors.New("empty name is not allowed")
	}
	return true, nil
}

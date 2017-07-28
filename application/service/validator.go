package Service

import (
	"regexp"
	"errors"
)

func AuthorizationHeaderValidator(h string) (token string, err error)  {
	if h == "" {
		return "", errors.New("empty string")
	}

	reg, err := regexp.MatchString("Bearer [A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*$", h)
	if !reg || err != nil {
		return "", errors.New("not a valid phonenumber")
	}

	return h[7:], nil
}

func PhoneNumberValidator(n string) (matched bool, err error) {
	if n == "" {
		return false, errors.New("empty string")
	}

	reg, err := regexp.MatchString("07[0-9]{8}$", n)
	if !reg || err != nil {
		return false, errors.New("not a valid phonenumber")
	}

	return true, nil
}

func NumberURIValidator(n string) (matched bool, err error) {
	reg, err := regexp.MatchString("^[0-9]{1,2}$", n)
	if !reg || err != nil {
		return false, errors.New("not a valid number")
	}

	return true, nil
}

func RanValidator(n string) (matched bool, err error) {
	if n == "" {
		return false, errors.New("empty string")
	}

	reg, err := regexp.MatchString("^[0-9]{6}$", n)
	if !reg || err != nil {
		return false, errors.New("not a valid number")
	}

	return true, nil
}

func RoleValidator(n string) (matched bool, err error) {
	if n == "" {
		return false, errors.New("empty string")
	}

	reg, err := regexp.MatchString("customer|administrator", n)
	if !reg || err != nil {
		return false, errors.New("not a valid role")
	}

	return true, nil
}

func NameValidator(n string) (matched bool, err error) {
	if n == "" {
		return false, errors.New("empty string")
	}

	reg, err := regexp.MatchString("^[\u00C0-\u017Fa-zA-Z]{2,64}$", n)
	if !reg || err != nil {
		return false, errors.New("not a valid name")
	}

	return true, nil
}

func LanguageValidator(n string) (matched bool, err error) {
	if n == "" {
		return false, errors.New("empty string")
	}

	reg, err := regexp.MatchString("de|en", n)
	if !reg || err != nil {
		return false, errors.New("not a valid language")
	}

	return true, nil
}

func MongoIdValidator(n string) (matched bool, err error) {
	if n == "" {
		return false, errors.New("empty string")
	}

	reg, err := regexp.MatchString("^[a-f0-9]{24}$", n)
	if !reg || err != nil {
		return false, errors.New("not a valid name")
	}

	return true, nil
}
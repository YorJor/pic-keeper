package fieldvalidate

import (
	"errors"

	"regexp"

	"github.com/Roongkun/software-eng-ii/internal/model"
)

func LoginUser(cred model.LoginCredentials) []error {
	fieldErrs := []error{}

	// TODO: regex for email
	if cred.Email == "" {
		fieldErrs = append(fieldErrs, errors.New(
			"email must be specified",
		))
	}
	if cred.Password == "" {
		fieldErrs = append(fieldErrs, errors.New(
			"password must be specified",
		))
	}

	return fieldErrs
}

func RegCustomer(newUser model.UserInput) []error {
	fieldErrs := []error{}

	// empty name
	if newUser.Name == "" {
		fieldErrs = append(fieldErrs, errors.New(
			"Name must be specified",
		))
	}

	// empty email
	if newUser.Email == "" {
		fieldErrs = append(fieldErrs, errors.New(
			"email must be specified",
		))
	} else {
		// email format
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		if emailRegex.MatchString(newUser.Email) == false {
			fieldErrs = append(fieldErrs, errors.New(
				"invalid email address",
			))
		}
	}

	// empty password
	if newUser.Password != nil || *newUser.Password == "" {
		fieldErrs = append(fieldErrs, errors.New(
			"password must be specified",
		))
	}

	return fieldErrs
}

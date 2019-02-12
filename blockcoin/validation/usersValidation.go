package validation

import (
	"errors"
	"api/utils"
	"api/models"
)

var (
	ErrEmptyFields = errors.New("Um ou mais campos vazios.")
	ErrInvalidEmail = errors.New("Email inválido.")
)

func VerifyPostUser(user models.User) (models.User, error) {
	if utils.IsEmpty(user.Nickname) || utils.IsEmpty(user.Email) || utils.IsEmpty(user.Password) {
		return user, ErrEmptyFields
	}
	if !utils.IsEmail(user.Email) {
		return user, ErrInvalidEmail
	}
	return user, nil
}
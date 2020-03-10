package requestValidate

import (
	"goerhubApi/model"
	"gopkg.in/go-playground/validator.v9"
)

func UserNameValidate(field validator.FieldLevel) bool {
	userModel := model.UserModel{}
	username := field.Field().String()
	if userModel.CheckUsernameExist(username) {
		return false
	}
	return true
}
func UserEmailValidate(field validator.FieldLevel) bool {
	userModel := model.UserModel{}
	email := field.Field().String()
	if userModel.CheckEmailExist(email) {
		return false
	}
	return true
}

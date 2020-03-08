package validators

import (
	"goerhubApi/model"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

func UserNameValidate(field validator.FieldLevel) bool {
	userModel := model.UserModel{}
	username := field.Field().String()
	log.Println(username)
	if userModel.CheckUsernameExist(username) {
		return false
	}
	return true
}

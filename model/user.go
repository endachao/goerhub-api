package model

import (
	"goerhubApi/dao"
	"time"
)

type Users struct {
	ID            int       `gorm:"primary_key;column:id;type:int(11);not null"`
	Nickname      string    `gorm:"column:nickname;type:varchar(20);not null"`
	Username      string    `gorm:"column:username;type:varchar(20);not null"`
	Password      string    `gorm:"column:password;type:varchar(255);not null"`
	Email         string    `gorm:"column:email;type:varchar(50);not null"`
	EmailVerified int8      `gorm:"column:email_verified;type:tinyint(1);not null"`
	Phone         string    `gorm:"column:phone;type:varchar(11);not null"`
	GoldNumber    int       `gorm:"column:gold_number;type:int(11);not null"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp"`
}

type UserModel struct {
}

func (UserModel) GetUserInfoByUserEmail(email string) (user Users, ok bool) {
	if dao.DB().Where("email = ?", email).Take(&user).RecordNotFound() {
		return
	}
	return user, true
}

func (UserModel) CheckUsernameExist(username string) bool {
	var user Users
	if dao.DB().Where("username = ?", username).Take(&user).RecordNotFound() {
		return false
	}
	return true
}

func (UserModel) CheckEmailExist(email string) bool {
	var user Users
	if dao.DB().Where("email = ?", email).Take(&user).RecordNotFound() {
		return false
	}
	return true
}

func (UserModel) CreateUser(users Users) error {
	if err := dao.DB().Create(&users).Error; err != nil {
		return err
	}
	return nil
}

package user

import (
	"golang.org/x/crypto/bcrypt"

	"giligili/pkg/database"
)


// GetUser 用ID获取用户
func GetUser(ID interface{}) (userModel User, err error) {
	database.DB.Where("id = ?", ID).First(&userModel)
	return 
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

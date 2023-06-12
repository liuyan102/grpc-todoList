package repository

import (
	"errors"
	"user/internal/service"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User struct {
	UserId         uint   `gorm:"primaryKey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

const (
	PasswordCost = 12
)

// GetUserInfo 获取用户信息
func (user *User) GetUserInfo(req *service.UserRequest) error {
	if err := DB.Where("user_name=?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return errors.New("userName Not Exist")
	}
	return nil
}

// CreateUser 获取用户信息
func (*User) CreateUser(req *service.UserRequest) error {
	var count int64
	DB.Where("user_name=?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("userName Exist")
	}
	user := &User{
		UserName: req.UserName,
		NickName: req.NickName,
	}

	_ = user.SetPassword(req.Password)
	err := DB.Create(&user).Error
	if err != nil {
		return errors.New("user Create Error")
	}
	return nil
}

// SetPassword 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// BuildUser 序列化User
func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		UserID:   uint32(item.UserId),
		UserName: item.UserName,
		NickName: item.NickName,
	}
	return &userModel
}

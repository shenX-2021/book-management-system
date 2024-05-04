package service

import (
	"book-management-system-server/dto"
	"book-management-system-server/global"
	"book-management-system-server/model"
	"crypto"
	"encoding/hex"
	"errors"

	"gorm.io/gorm"
)

type UserService struct{}

// 注册账号
func (s UserService) Register(dto *dto.UserRegisterDto) error {
	err := global.DB.Where("account=?", dto.Account).First(&model.User{}).Error
	if err == nil {
		return errors.New("已存在该账号")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hash := crypto.SHA256.New()
	hash.Write([]byte(dto.Pwd))

	pwdHash := hex.EncodeToString(hash.Sum(nil))

	return global.DB.Create(&model.User{Name: dto.Name, Account: dto.Account, Pwd: pwdHash}).Error
}

// 获取用户信息
func (s UserService) GetUser(dto *dto.UserLoginDto) (*model.User, error) {
	user := &model.User{}

	hash := crypto.SHA256.New()
	hash.Write([]byte(dto.Pwd))
	pwdHash := hex.EncodeToString(hash.Sum(nil))

	result := global.DB.Where("account=? AND pwd=?", dto.Account, pwdHash).First(user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("账号或密码错误")
	} else if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

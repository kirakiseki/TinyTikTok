package dao

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/model/dto"
	"TinyTikTok/utils"
	"gorm.io/gorm"
)

type UserPO struct {
	gorm.Model
	dto.UserDTO
}

func (u UserPO) TableName() string {
	return "users"
}

// SearchUserByUserName 通过username查询数据库中User信息 返回UserDTO
func SearchUserByUserName(username string) dto.UserDTO {
	user := dto.UserDTO{}
	//查询用户信息
	err := setup.Mdb.Model(&UserPO{}).Where("user_name = ?", username).Find(&user).Error
	if err != nil {
		setup.Logger("common").Err(err).Send()
	}
	//将查询到的密码解密
	password, err := utils.Base64Decode(user.Password)
	if err != nil {
		setup.Logger("common").Err(err).Send()
	}
	user.Password = string(password)
	//返回数据
	return user
}

// SaveUser 插入一条user 返回值表示是否插入成功
func SaveUser(userDTO *dto.UserDTO) bool {
	//加密
	userDTO.Password = utils.Base64Encode([]byte(userDTO.Password))
	//插入对象
	err := setup.Mdb.Create(&UserPO{
		UserDTO: *userDTO,
	}).Error
	if err != nil {
		setup.Logger("common").Err(err).Send()
		return false
	}
	return true
}

// SearchUserByUserId  根据userId查询用户信息
func SearchUserByUserId(userId int64) dto.UserDTO {
	userDTO := dto.UserDTO{}
	err := setup.Mdb.Model(&UserPO{}).Where("user_id = ?", userId).Find(&userDTO).Error
	if err != nil {
		setup.Logger("common").Err(err).Send()
	}
	password, err := utils.Base64Decode(userDTO.Password)
	if err != nil {
		setup.Logger("common").Err(err).Send()
	}
	userDTO.Password = string(password)
	return userDTO
}

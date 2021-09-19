package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	// data.Password = ScryptPwd(data.Password)
	err := db.Create(&data).Error

	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)

	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}

	return errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	offset := (pageNum - 1) * pageSize

	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}

	err := db.Limit(pageSize).Offset(offset).Find(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	return users
}

// EditUser 编辑用户
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	var user User
	maps["username"] = data.Username
	maps["role"] = data.Role

	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// ScryptPwd 密码加密
func ScryptPwd(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 5, 55, 22, 222, 11}

	HashPwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	fpw := base64.StdEncoding.EncodeToString(HashPwd)
	return fpw
}

// BeforeSave gorm钩子函数
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPwd(u.Password)

	return
}

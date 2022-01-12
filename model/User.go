package model

import (
	"GinProject/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=18" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

//查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

//新建用户
func CreateUser(data *User) int {
	data.Password = ScryptPwd(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//2022年1月7日
//钩子函数，创建用户前执行密码加密
//func (u *User) BeforeSave(_ *gorm.DB) (err error) {
//	u.Password = ScryptPwd(u.Password)
//	return nil
//}

//查询用户列表，分页显示
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := db.Model(&users).Count(&total).Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

//编辑用户（密码除外）
func EditUser(id int, data *User) int {
	//根据 `struct` 更新属性，只会更新非零值的字段
	//根据 `map` 更新属性
	var user User
	maps := map[string]interface{}{
		"username": data.Username,
		"role":     data.Role,
	}
	//maps["username"] = data.Username
	//maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func ScryptPwd(password string) string {
	//func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error)
	KeyLen := 10
	slat := make([]byte, 8)
	slat = []byte{12, 21, 32, 23, 43, 34, 54, 45}
	newPwd, err := scrypt.Key([]byte(password), slat, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(newPwd)
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//登陆验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ErrorUserNotExist
	}
	if ScryptPwd(password) != user.Password {
		return errmsg.ErrorPasswordWrong
	}
	if user.Role != 0 {
		return errmsg.ErrorUserNoRight
	}
	return errmsg.SUCCESS
}

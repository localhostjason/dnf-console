package model

import (
	"github.com/localhostjason/webserver/db"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	// 基本信息
	Id            int64      `json:"id" gorm:"primaryKey"`
	Username      string     `json:"username" gorm:"type:string;size:64;unique;not null"`
	Password      string     `json:"-" gorm:"column:_password;type:string;size:128"`
	LastLoginTime *time.Time `json:"last_login_time"`
	Time          time.Time  `json:"time"`                           // 创建时间
	JwtKey        uuid.UUID  `json:"-" gorm:"type:string;size:128;"` // 为每个用户存一个唯一的jwt key (通用唯一识别码)

	Role  string `json:"role"`
	Email string `json:"email" gorm:"type:string;size:64"`
	Desc  string `json:"desc" gorm:"type:string;size:256"`

	IsSuperAdmin bool `json:"is_super_admin"`
}

func (u *User) SetPassword(password string) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return
	}
	u.Password = string(b)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"username":        u.Username,
		"role":            u.Role,
		"email":           u.Email,
		"desc":            u.Desc,
		"last_login_time": u.LastLoginTime,
	}
}

func init() {
	db.RegTables(&User{})
}

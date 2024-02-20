package domian

import "time"

// 用户相关信息
type User struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	EmailValidated bool      `json:"emailValidated"`
	Phone          string    `json:"phone"`
	PhoneValidated bool      `json:"phoneValidated"`
	Status         bool      `json:"status"`
	CreateTime     time.Time `json:"createTime"`
}

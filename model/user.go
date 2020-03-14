package model

import (
	"fmt"

	"github.com/demos/api_server/pkg/auth"
	"github.com/demos/api_server/pkg/constant"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm: "column:not null" binding: "required" validate: "min=1,max=32"`
	Password string `json:"password" gorm: "column:not null" binding: "required" validate: "min=5,max=32"`
}

// TableName for user table
func (userModel *UserModel) TableName() string {
	return "tb_users"
}

// Create new user
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func (u *UserModel) Update() error {
	return DB.Self.Save(&u).Error
}

func GetUser(user string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", user).First(&u)
	return u, d.Error
}

func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constant.DefaultLimit
	}
	users := make([]*UserModel, 0)
	where := fmt.Sprintf("username like '%%%s%%'", username)
	var count uint64
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return err
}

package models

import (
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/prometheus/common/log"
	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `orm:"unique;size(120)" valid:"Email; MaxSize(100);Required" alias:"邮箱" json:"email"`
	PassWord string `orm:"size(100)" valid:"MaxSize(100);Required" alias:"密码" json:"-"`
	Salt     string `orm:"size(50)" json:"-"`
	Avatar   string `xorm:"size(2048)" json:"avatar"`
}

const userTableName = "user"

func (u *User) TableName() string {
	return userTableName
}

func (u *User) Valid() error {
	valid := validation.Validation{}
	b, err := valid.Valid(u)
	if err != nil {
		log.Info("valid user error ", err)
		return err
	}
	if !b {
		for _, err := range valid.Errors {
			return errors.New(fmt.Sprintf("%s %s ", err.Field, err.Message))
		}
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	o := orm.NewOrm()
	if err := o.Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) GetByEmail(fields ...string) error {
	o := orm.NewOrm()
	if err := o.QueryTable(userTableName).Filter("email", u.Email).One(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Save() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) EncodePassword() {
	newPasswd := pbkdf2.Key([]byte(u.PassWord), []byte(u.Salt), 10000, 50, sha256.New)
	u.PassWord = fmt.Sprintf("%x", newPasswd)
}

func (u *User) ValidatePassword(passwd string) bool {
	newUser := &User{PassWord: passwd, Salt: u.Salt}
	newUser.EncodePassword()
	return subtle.ConstantTimeCompare([]byte(u.PassWord), []byte(newUser.PassWord)) == 1
}

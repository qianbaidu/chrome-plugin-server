package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/prometheus/common/log"
)

type Collection struct {
	Id       int64  `json:"id"`
	Url      string `orm:"size(2048)" valid:"MaxSize(2048);Required" json:"url"`
	Title    string `orm:"size(100)" valid:"MaxSize(100);Required" json:"title"`
	UserId   int64  `valid:"Required" json:"user_id"`
	Category int    `orm:"size(2)" json:"category"`
	Status   int    `xorm:"size(2)" json:"status"`
}

const collectionTableName = "collection"

func (c *Collection) TableName() string {
	return collectionTableName
}

func (c *Collection) Valid() error {
	valid := validation.Validation{}
	b, err := valid.Valid(c)
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

func (c *Collection) Save() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *Collection) Edit() error {
	log.Info(c.Title, c.UserId)
	if _, err := orm.NewOrm().Update(c, "url", "title", "category"); err != nil {
		log.Error("update error ", err)
		return err
	}
	return nil
}

func (c *Collection) GetByUserId() ([]Collection, error) {
	o := orm.NewOrm()
	list := make([]Collection, 0)
	if _, err := o.QueryTable(collectionTableName).Filter("user_id", c.UserId).All(&list); err != nil {
		return list, err
	}
	return list, nil
}

func (c *Collection) GetById() error {
	o := orm.NewOrm()
	if err := o.QueryTable(collectionTableName).Filter("id", c.Id).One(c); err != nil {
		return err
	}
	return nil
}

func (c *Collection) GetUserUrl() error {
	o := orm.NewOrm()
	if err := o.QueryTable(collectionTableName).Filter("id", c.Id).Filter("user_id", c.UserId).One(c); err != nil {
		return err
	}
	return nil
}

func (c *Collection) DeleteById() error {
	o := orm.NewOrm()
	if _, err := o.Delete(c); err != nil {
		return err
	}
	return nil
}

func (c *Collection) DeleteUserUrl() error {
	o := orm.NewOrm()
	if _, err := o.Delete(c); err != nil {
		log.Error("delete user url error ",err)
		return err
	}
	return nil
}

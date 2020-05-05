package controllers

import (
	"github.com/prometheus/common/log"
	"github.com/qianbaidu/chrome-server/models"
	"github.com/qianbaidu/chrome-server/util"
	"strings"
)

type CollectionController struct {
	BaseAuthController
}

// @router /collection/save-url [post]
func (c *CollectionController) SaveUrl() {
	m := &models.Collection{}
	m.Url = strings.TrimSpace(c.GetString("url"))
	if !strings.HasPrefix(m.Url,"http") {
		c.Data["json"] = util.JsonMsg("参数有误", util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}
	m.Title = util.HtmlDecode(strings.TrimSpace(c.GetString("title")))
	m.Category, _ = c.GetInt("category")
	m.UserId = c.UserId
	if err := m.Valid(); err != nil {
		c.Data["json"] = util.JsonMsg(err, util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}

	if err := m.Save(); err == nil {
		c.Data["json"] = util.JsonMsg("", util.SUCCESS, "")
	} else {
		c.Data["json"] = util.JsonMsg(err, util.FAILED, "")
	}
	c.ServeJSON()
	return
}

// @router /collection/list [get]
func (c *CollectionController) List() {
	m := &models.Collection{}
	m.UserId = c.UserId
	if list, err := m.GetByUserId(); err != nil {
		c.Data["json"] = util.JsonMsg("", util.FAILED, list)
	} else {
		res := make(map[string][]models.Collection, 0)
		for _, v := range list {
			if cate := util.GetCategoryName(v.Category); len(cate) > 0 {
				if cateList, ok := res[cate]; ok {
					cateList = append(cateList, v)
					res[cate] = cateList
				} else {
					cateList := make([]models.Collection, 0)
					cateList = append(cateList, v)
					res[cate] = cateList
				}
			}
		}
		c.Data["json"] = util.JsonMsg("", util.SUCCESS, res)
	}
	c.ServeJSON()
	return
}

// @router /collection/delete-url [post]
func (c *CollectionController) DeleteUrl() {
	m := &models.Collection{}
	id, err := c.GetInt64("id")
	if err != nil || id < 1{
		c.Data["json"] = util.JsonMsg("参数有误", util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}
	m.Id = id
	m.UserId = c.UserId
	if err := m.DeleteUserUrl(); err != nil {
		c.Data["json"] = util.JsonMsg("参数有误", util.FAILED, "")
	} else {
		c.Data["json"] = util.JsonMsg("", util.SUCCESS, "")
	}
	c.ServeJSON()
	return
}

// @router /collection/edit-url/:id [post]
func (c *CollectionController) EditUrl() {
	m := &models.Collection{}
	id, err := c.GetInt64("id")
	if err != nil {
		c.Data["json"] = util.JsonMsg("参数有误", util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}
	url := strings.TrimSpace(c.GetString("url"))
	if !strings.HasPrefix(url,"http") {
		c.Data["json"] = util.JsonMsg("参数有误", util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}
	title := util.HtmlDecode(strings.TrimSpace(c.GetString("title")))
	category, _ := c.GetInt("category")
	m.Id = id
	m.UserId = c.UserId
	if err = m.GetUserUrl(); err == nil {
		m.Url = url
		m.Category = category
		m.Title = title
		log.Info("m.Title",m.Title)
		if err := m.Edit(); err == nil {
			c.Data["json"] = util.JsonMsg("", util.SUCCESS, "")
		} else {
			c.Data["json"] = util.JsonMsg("", util.FAILED, "")
		}
	} else {
		c.Data["json"] = util.JsonMsg("", util.PARAM_ERROR, m)
	}

	c.ServeJSON()
	return
}

// @router /collection/url/:id [get]
func (c *CollectionController) Url() {
	m := &models.Collection{}
	id, err := c.GetInt64("id")
	if err != nil {
		c.Data["json"] = util.JsonMsg("参数有误", util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}
	m.Id = id
	m.UserId = c.UserId
	if err = m.GetUserUrl(); err != nil {
		c.Data["json"] = util.JsonMsg("参数有误", util.FAILED, "")
	} else {
		c.Data["json"] = util.JsonMsg("", util.SUCCESS, m)
	}

	c.ServeJSON()
	return
}

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"github.com/qianbaidu/chrome-server/models"
	"github.com/qianbaidu/chrome-server/util"
	"strings"
)

type AuthController struct {
	beego.Controller
}

// @router /login [post]
func (c *AuthController) Login() {
	m := &models.User{}
	m.Email = strings.TrimSpace(c.GetString("email"))
	passWord := strings.TrimSpace(c.GetString("password"))
	m.PassWord = passWord
	if err := m.Valid(); err != nil {
		c.Data["json"] = util.JsonMsg(err, util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}

	if err := m.GetByEmail(); err != nil {
		c.Data["json"] = util.JsonMsg("用户名或密码错误", util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}

	if b := m.ValidatePassword(passWord); b == true {
		token := util.GenerateToken(m.Id, m.Email)
		type Resp struct {
			*models.User
			Token string `json:"token"`
		}
		m.PassWord = ""
		m.Salt = ""
		m.Email = strings.Split(m.Email, "@")[0]
		data := &Resp{
			m,
			token,
		}
		c.Ctx.SetCookie("Authorization", token, util.JWT_EXPIRE_SECONDS)
		c.SetSession("UserId", m.Id)
		c.Data["json"] = util.JsonMsg("", util.SUCCESS, data)
	} else {
		c.Data["json"] = util.JsonMsg("用户名或密码错误", util.PARAM_ERROR, "")
	}

	c.ServeJSON()
	return
}

// @router /logout [post]
func (c *AuthController) Logout() {
	c.Ctx.SetCookie("Authorization", "")
	c.Ctx.SetCookie("UserId", "")
	c.Data["json"] = util.JsonMsg("", util.SUCCESS, "")

	c.ServeJSON()
	return
}

// @router /register [post]
func (c *AuthController) Register() {
	m := &models.User{}
	var err error
	m.Email = strings.TrimSpace(c.GetString("email"))
	m.PassWord = strings.TrimSpace(c.GetString("password"))
	if err := m.Valid(); err != nil {
		c.Data["json"] = util.JsonMsg(err, util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}

	//查询邮箱是否存在
	if err := m.GetByEmail(); err == nil {
		c.Data["json"] = util.JsonMsg("邮箱已存在", util.PARAM_ERROR, "")
		c.ServeJSON()
		return
	}

	//生成salt 和密码
	m.Salt, err = util.GetUserSalt()
	if err != nil {
		log.Error("GetUserSalt error ", err)
	}
	m.EncodePassword()

	if err := m.Save(); err != nil {
		c.Data["json"] = util.JsonMsg(err, util.FAILED, "")
	} else {
		c.Data["json"] = util.JsonMsg("", util.SUCCESS, "")
	}

	c.ServeJSON()
	return
}

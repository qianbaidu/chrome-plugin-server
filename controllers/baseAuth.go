package controllers

import (
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"github.com/qianbaidu/chrome-server/util"
	"strconv"
)

type BaseAuthController struct {
	beego.Controller
	UserId   int64
	UserName string
}

func (c *BaseAuthController) Prepare() {
	if  err := c.checkToken(); err != nil {
		c.Data["json"] = util.JsonMsg(err, util.AUTH_ERROR, "")
		c.ServeJSON()
		return
	}
}

func (c *BaseAuthController) checkToken() error {
	token := c.Ctx.Input.Cookie("Authorization")
	if user, err := util.ValidateToken(token); err != nil {
		log.Info("ValidateToken error ", err)
		return err
	} else {
		userId, err := strconv.ParseInt(user.Id, 10, 64)
		if err != nil {
			return nil
		}
		c.UserId = userId
		c.UserName = user.Name
		return nil
	}
}

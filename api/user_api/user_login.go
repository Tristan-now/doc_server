package user_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/res"
	"gvd_server/utils/ip"
	"gvd_server/utils/jwts"
	"gvd_server/utils/pwd"
	"time"
)

type UserLoginRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

// UserLoginView 用户登录
// @Tags 用户管理
// @Summary 用户登录
// @Description 用户登录
// @Param data body UserLoginRequest true "参数"
// @Router /api/login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserLoginView(c *gin.Context) {
	var cr UserLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidError(err, &cr, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在", cr.UserName)
		log_stash.NewFailLogin("用户名不存在", cr.UserName, cr.Password, c)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	if !pwd.CheckPwd(user.Password, cr.Password) {
		global.Log.Warn("用户密码错误", cr.UserName, cr.Password)
		log_stash.NewFailLogin("用户密码错误", cr.UserName, cr.Password, c)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	token, err := jwts.GenToken(jwts.JwyPayLoad{
		UserName: user.UserName,
		NickName: user.NickName,
		RoleID:   user.RoleID,
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("生成token失败", c)
		return
	}
	c.Request.Header.Set("token", token)
	log_stash.NewSuccessLogin(c)
	global.DB.Model(&user).Update("lastLogin", time.Now())

	_ip := c.ClientIP()
	addr := ip.GetAddr(_ip)
	ua := c.Request.Header.Get("User-Agent")

	go func() {
		// 加一条登录记录
		err = global.DB.Create(&models.LoginModel{
			UserID:   user.ID,
			IP:       _ip,
			NickName: user.NickName,
			UA:       ua,
			Token:    token,
			Addr:     addr,
		}).Error

		if err != nil {
			global.Log.Error(err)
		}
	}()
	res.OKWithData(token, c)
	return

}

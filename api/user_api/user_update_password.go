package user_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"gvd_server/utils/pwd"
)

type UserUpdatePasswordRequest struct {
	OldPwd   string `json:"oldPwd" binding:"required" label:"之前的密码"`
	Password string `json:"password" binding:"required" label:"新密码"`
}

// UserUpdatePasswordView 用户修改密码
// @Tags 用户管理
// @Summary 用户修改密码
// @Description 用户修改密码
// @Param token header string true "token"
// @Param data body UserUpdatePasswordRequest true "参数"
// @Router /api/users_password [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	var cr UserUpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)
	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMsg("原密码错误", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	global.DB.Model(user).Update("password", hashPwd)

	res.OKWithMsg("用户密码修改成功", c)
	return

}

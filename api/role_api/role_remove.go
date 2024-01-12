package role_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
)

func (RoleApi) RoleRemoveView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var role models.RoleModel
	err = global.DB.Preload("DocsList").Preload("UserList").Take(&role, cr.ID).Error
	if err != nil {
		res.FailWithMsg("不存在的角色", c)
		return
	}

	if role.IsSystem {
		res.FailWithMsg("系统角色，不可删除", c)
		return
	}
	global.Log.Infof("删除角色 %s，关联用户 %d 个，删除关联文档 %d 个", role.Title, len(role.UserList), len(role.DocsList))
	if len(role.UserList) > 0 {
		// 统一修改用户的角色id为2
		global.DB.Model(&role.UserList).Update("roleID", "2")
	}
	if len(role.DocsList) > 0 {
		global.DB.Model(&role).Association("DocsList").Delete(role.DocsList)
	}

	global.DB.Delete(&role)
	res.OKWithMsg("删除角色成功", c)
}

package role_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/models"
)

type RoleListResponse struct {
	models.RoleModel
	DocCount  int `json:"DocCount"`  //角色拥有的文档数
	UserCount int `json:"UserCount"` //角色下的用户数
}

func (RoleApi) RoleListView(c *gin.Context) {
	var cr models.Pagination
}

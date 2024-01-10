package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvd_server/global"
	"gvd_server/models"
)

type JwyPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
	UserName string `json:"userName"`
}

type CustomClaims struct {
	JwyPayLoad
	jwt.StandardClaims
}

func (c CustomClaims) GetUser() (user *models.UserModel, err error) {
	user = new(models.UserModel)
	err = global.DB.Take(user, c.UserID).Error
	return
}

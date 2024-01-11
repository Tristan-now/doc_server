package flags

import (
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
)

func DB() {
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").
		AutoMigrate(
			&models.UserModel{},
			&models.RoleModel{},
			&models.DocModel{},
			&models.UserCollDocModel{}, // 如果有自定义连接表，一定要放另外两张表的后面
			&models.RoleDocModel{},
			&models.ImageModel{},
			&models.UserPwdDocModel{},
			&models.LoginModel{},
			&models.DocDataModel{},
			//&log_stash.LogModel{},
		)

	if err != nil {
		logrus.Fatalf("数据库迁移失败 err:%s", err.Error())
	}
	logrus.Infof("数据库迁移成功！")

}

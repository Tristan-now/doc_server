package models

type UserPwdDocModel struct {
	Model
	UserID uint `gorm:"column:user_id" json:"userID"`
	DocID  uint `gorm:"column:doc_id" json:"docID"`
}

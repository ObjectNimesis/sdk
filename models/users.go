package models

type User struct {
	ID uint32 `gorm:"primaryKey;auto_increment" json:"id"`

	Username    string `gorm:"unique;not null" json:"user_name"`
	DisplayName string `gorm:"column:display_name;not null" json:"display_name"`

	Password string `gorm:"not null" json:"password"`
	Gender   string `gorm:"type:varchar(255);default:null" json:"gender,omitempty"`

	Pronouns string `gorm:"type:varchar(255);default:null" json:"pronouns,omitempty"`

	Avatar string  `gorm:"type:varchar(255);default:null" json:"avatar,omitempty"`
	Emails []Email `gorm:"foreignKey:UserID" json:"emails"`
}

func (User) TableName() string {
	return "users"
}

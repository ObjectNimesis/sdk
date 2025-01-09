package models

type EmailAddressType string

const (
	primaryEmail     EmailAddressType = "primary"
	recoveryEmail    EmailAddressType = "recovery"
	contactEmail     EmailAddressType = "contact"
	alternativeEmail EmailAddressType = "alternative"
)

// EmailTypes provides accessible constants for EmailAddressType
var EmailTypes = struct {
	Primary     EmailAddressType
	Recovery    EmailAddressType
	Contact     EmailAddressType
	Alternative EmailAddressType
}{
	Primary:     primaryEmail,
	Recovery:    recoveryEmail,
	Contact:     contactEmail,
	Alternative: alternativeEmail,
}

type Email struct {
	ID int `gorm:"primaryKey;auto_increment" json:"id"`

	Address  string           `gorm:"unique;not null" json:"address"`
	Type     EmailAddressType `gorm:"type:varchar(255);not null" json:"type"`
	Verified bool             `gorm:"default:false" json:"verified"`

	UserID uint32 `gorm:"not null" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
}

func (Email) TableName() string {
	return "emails"
}

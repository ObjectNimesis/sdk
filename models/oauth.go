package models

import (
	pq "github.com/lib/pq"
)

type OAuthClient struct {
	ClientID     string `gorm:"primaryKey;column:client_id;unique" json:"client_id"`
	ClientSecret string `gorm:"column:client_secret;unique" json:"client_secret"`

	Name        string `gorm:"column:name" json:"client_name"`
	Description string `gorm:"column:description" json:"description"`
	Homepage    string `gorm:"column:homepage" json:"homepage"`

	CallbackUrls pq.StringArray `gorm:"column:callback_urls;type:string[]" json:"callback_urls"`
	Tags         pq.StringArray `gorm:"column:tags;type:string[]" json:"tags"`

	OwnerID uint `gorm:"column:owner_id" json:"owner_id"`
}

func (OAuthClient) TableName() string {
	return "oauth_clients"
}

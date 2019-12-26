package database

import("github.com/jinzhu/gorm")

type UserPost struct{
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT"; primary_key;"`
	UserUID string
	Text string
	UserName string
	PhotoURL string
}
package database

import("github.com/jinzhu/gorm")

type UserPost struct{
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT"; gorm:primary_key;`
	UserUID string
	UserName string
	PhotoURL string
	Text string
}
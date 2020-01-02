package database

type UserPost struct{
	ID int `gorm:"AUTO_INCREMENT" gorm:"primary_key" json:"ID"`
	UserUID string `json:"userID"`
	UserName string `json:"userName"`
	PhotoURL string `json:"photoURL"`
	Text string	`json:"text"`
}
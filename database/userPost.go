package database
import ("time")

type UserPost struct{
	ID int `gorm:"AUTO_INCREMENT" gorm:"primary_key" json:"ID"`
	UserName string `json:"userName"`
	PhotoURL string `json:"photoURL"`
	Text string	`json:"text"`
	CreatedAt time.Time `json:"time"`
}
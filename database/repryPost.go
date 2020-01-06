package database
import ("time")

type RepryPost struct {
	ID int `gorm:"AUTO_INCREMENT" gorm:"primary_key" json:"ID"`
	RepryID int `json:"RepryID"`
	UserUID string `json:"userID"`
	UserName string `json:"userName"`
	PhotoURL string `json:"photoURL"`
	Text string	`json:"text"`
	CreatedAt time.Time `json:"time"`
}
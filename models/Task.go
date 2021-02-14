package models

//Task represent a scheduled task
type Task struct {
	ID          uint `json:"id" gorm:"private_key"`
	Hour        int  `json:"hour"`
	Minute      int  `json:"minute"`
	Second      int  `json:"second"`
	Day         int  `json:"day"`
	Periodicity int  `json:"peridicity"`
	ServerID    uint `json:"server_id"`
	IsEnable    bool `json:"isenable"`
}

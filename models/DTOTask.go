package models

//DTOTask is struct for create task
type DTOTask struct {
	Hour        int  `json:"hour"`
	Minute      int  `json:"minute"`
	Second      int  `json:"second"`
	Day         int  `json:"day"`
	Periodicity int  `json:"peridicity"`
	ServerID    uint `json:"serverid"`
}

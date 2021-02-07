package models

//Server is a struct that represent a server
type Server struct {
	ID         uint   `json:"id" gorm:"private_key"`
	Name       string `json:"name"`
	Addr       string `json:"addr"`
	SSHPort    int    `json:"sshport"`
	User       string `json:"user"`
	PrivateKey string `json:"privatekey"`
	Tasks      []Task `json:"tasks"`
}

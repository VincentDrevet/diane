package models

//DTOServer is struct for mainly POST request (creation of Server struct)
type DTOServer struct {
	Name       string `json:"name"`
	Addr       string `json:"addr"`
	SSHPort    int    `json:"sshport"`
	User       string `json:"user"`
	PrivateKey string `json:"privatekey"`
	Tasks      []Task `json:"tasks"`
}

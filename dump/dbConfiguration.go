package dump

type DBConfiguration struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DB       string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
}

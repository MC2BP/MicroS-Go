package configlib

type Config struct {
	ID          int                `json:"Id"`
	Environment string             `json:"Environment"`
	Webserver   Webserver          `json:"Webserver"`
	Services    map[string]Service `json:"Services"`
}

type Webserver struct {
	Host string   `json:"Host"`
	Port int      `json:"Port"`
	ReadTimeout int `json:"ReadTimeout"`
	WriteTimeout int `json:"WriteTimeout"`
	Cors []string `json:"Cors"`
}

type Service struct {
	ID   int    `json:"Id"`
	Host string `json:"Host"`
}

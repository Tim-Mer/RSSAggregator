package config


type Config struct {
	DbURL string `json:"db_url"`
}

func Read() Config {
	return Config{}
}
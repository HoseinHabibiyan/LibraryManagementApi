package config

type Config struct {
	DbName   string
	Host     string
	Port     string
	Username string
	Password string
}

func GetConfig() *Config {
	cfg := Config{
		DbName:   "library_management",
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
	}

	return &cfg
}

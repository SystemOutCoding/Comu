package configTemplate

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: "sampleuser",
			Password: "samplepassword",
			Name:     "sampledb",
		},
	}
}

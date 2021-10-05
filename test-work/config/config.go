package config

// MyConfig is needed to work correctly with the go-envconfig library
type MyConfig struct {
	Ps PsConfig
}

// PsConfig includes Environment variables for accessing the postgres database
type PsConfig struct {
	Dbname   string `env:"POSTGRES_DBNAME"`
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

package config

type MyConfig struct {
	Ps PsConfig
}

type PsConfig struct {
	Dbname   string `env:"POSTGRES_DBNAME"`
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

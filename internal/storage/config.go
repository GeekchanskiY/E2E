package storage

type Config struct {
	Name     string `config:"DB_NAME"`
	User     string `config:"DB_USER"`
	Password string `config:"DB_PASSWORD"`
	Host     string `config:"DB_HOST"`
	Port     string `config:"DB_PORT"`
}

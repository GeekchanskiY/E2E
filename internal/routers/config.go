package routers

type Config struct {
	Host   string `config:"APP_HOST"`
	Port   int    `config:"APP_PORT"`
	Secret string `config:"APP_SECRET"`
}

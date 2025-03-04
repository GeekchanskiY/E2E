package controllers

type Config struct {
	Secret string `config:"APP_SECRET"`
}

func (cfg *Config) GetSecret() string {
	return cfg.Secret
}

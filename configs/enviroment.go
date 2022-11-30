package config

type configuration struct {
	port string `env:"PORT" envDefault:":3333"`
}

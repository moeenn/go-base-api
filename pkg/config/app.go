package config

type App struct {
	Port string
}

var AppConfig = App{
	Port: ":5000",
}

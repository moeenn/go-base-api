package config

type App struct {
	Port        string
	RequiredEnv []string
}

var AppConfig = App{
	Port: ":5000",
	RequiredEnv: []string{
		"APP_SECRET",
	},
}

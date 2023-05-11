package config

type Env struct {
	Secret       string
	FrontendHost string
}

// TODO: read from env
var EnvConfig = Env{
	Secret:       "abc12345",
	FrontendHost: "http://localhost:3000",
}

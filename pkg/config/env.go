package config

type Env struct {
	Secret string `json:"secret"`
}

var EnvConfig = Env{
	Secret: "abc12345",
}

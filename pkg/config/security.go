package config

type Security struct {
	LoginTokenExpiryMinutes uint
}

var SecurityConfig = Security{
	LoginTokenExpiryMinutes: 15,
}

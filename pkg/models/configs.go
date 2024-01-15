package models

type AppConfigs struct {
	Port     string
	LogLevel string
	Cache    *Redis
	Balances []Balance
}

type Balance struct {
	Url    string
	Weight int
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

package models

type AppConfigs struct {
	Port     string
	LogLevel string
	Cache    *Redis
	Ratio    *PercentageDivision
	API1     string
	API2     string
}

type Redis struct {
	Host string
	Port string
}

type PercentageDivision struct {
	Api1Percent float64
	Api2Percent float64
}

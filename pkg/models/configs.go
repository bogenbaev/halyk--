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
	Total       int
	Api1Count   int
	Api2Count   int
	Api1Percent float64
	Api2Percent float64
}

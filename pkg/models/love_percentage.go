package models

type ExternalLovePercentage struct {
	FName      string `json:"f_name"`
	SName      string `json:"s_name"`
	Percentage string `json:"percentage"`
	Result     string `json:"result"`
}

type LovePercentage struct {
	SName string
	FName string
}

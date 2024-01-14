package models

type ExternalDateFact struct {
	Text   string `json:"text"`
	Year   int    `json:"year"`
	Number int    `json:"number"`
	Found  bool   `json:"found"`
	Type   string `json:"type"`
}

type DateFact struct {
	Fragment bool
	Json     bool
}

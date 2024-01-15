package models

type Response struct {
	Body   any    `json:"Body"`
	Status string `json:"Method"`
}

type Request struct {
	Body   any
	IP     string
	Method string
	Url    string
}

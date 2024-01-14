package models

type Response struct {
	Body   any
	Status string
}

type Request struct {
	Body   any
	IP     string
	Method string
	Url    string
}

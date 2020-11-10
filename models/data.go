package models

type MessageData struct {
	Surname   string
	Givenname string
	Birthd    string
}
type Message struct {
	Type string
	Data MessageData
}

type ResultMessage struct {
	Type string
	Info string
}

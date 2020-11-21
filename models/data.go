package models

// MessageData : evtl. nicht mehr benötigt nach umbau auf "InMessage"
type MessageData struct {
	Id        string
	Surname   string
	Givenname string
	Birthd    string
}

type Message struct {
	Type string
	Data MessageData
}

type InMessage struct {
	Type string
	Data map[string]interface{} `json:"data"`
}

type ReqNewVisitor struct {
	Surname   string
	Givenname string
	Birthd    string
}

type ReqUpdVisitor struct {
	ID        string
	Surname   string
	Givenname string
	Birthd    string
}

type ResultMessage struct {
	Type string
	Info string
}

type Visitor struct {
	ID        int
	Surname   string
	Givenname string
	Birthd    string
	Chd       string
}

// DBInfo : Database connection info
var DBInfo string

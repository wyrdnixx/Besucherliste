package models

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

/*   -> müsste mit interface gehen
type Message struct {
	Type string
	Data interface{}
}
*/
type ResultMessage struct {
	Type string
	Info string
}

type Visitor struct {
	Id        int
	Surname   string
	Givenname string
	Birthd    string
	chd       string
}

// DBInfo : Database connection info
var DBInfo string

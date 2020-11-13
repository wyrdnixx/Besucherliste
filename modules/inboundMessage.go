package modules

type transmitter struct {
	Client  Client
	Message []byte
}

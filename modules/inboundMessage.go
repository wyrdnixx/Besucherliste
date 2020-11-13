package modules

type InboundMessage struct {
	Client  Client
	Message []byte
}

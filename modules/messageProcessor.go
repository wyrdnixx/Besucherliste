package modules

import (
	"encoding/json"
	"fmt"

	"github.com/wyrdnixx/Besucherliste/models"
)

func ConsumeMessage(_msg []byte) models.ResultMessage {

	var m models.Message
	var result models.ResultMessage

	err := json.Unmarshal(_msg, &m)
	if err != nil {
		fmt.Printf("Error decoding: %s\n", err.Error())

		result.Type = "Error"
		result.Info = err.Error()
		return result
	} else {
		fmt.Printf("got message type: %s\n", m.Type)

		switch m.Type {
		case "NewVisitor":
			fmt.Printf("Request new Visitor created...\n")
			testfunc()
		default:
			fmt.Printf("Unknown message type...\n")

		}

		result.Type = "Success"
		result.Info = "Message Consumend"
		return result
	}

}

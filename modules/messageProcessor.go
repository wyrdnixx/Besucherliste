package modules

import (
	"encoding/json"
	"fmt"

	"github.com/wyrdnixx/Besucherliste/models"
)

// ConsumeMessage : getting inbound message and processing
func ConsumeMessage(_inbound transmitter) models.ResultMessage {

	var _msg = _inbound.Message
	var m models.Message
	var result models.ResultMessage

	//_inbound.Client.hub.broadcast <- []byte("BC-Message:")

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
			// 			res, err := testfunc(m.Data)
			res, err := InsertVisitor(m.Data)

			if err != nil {
				result.Type = "Error"
				result.Info = err.Error()
			} else {
				result.Type = "Success"
				result.Info = "Visitor created result: " + res

			}

			//ToDo: currently stops processing
			//bc(_inbound)

			return result
		default:
			fmt.Printf("Unknown message type...\n")
			result.Type = "Error"
			result.Info = "Unknown message type: " + m.Type
			return result
		}

	}

}

func bc(_inbound transmitter) {
	fmt.Printf("in bc func\n")
	_inbound.Client.hub.broadcast <- []byte("Holla BC")
	fmt.Printf("end bc func\n")
}

package modules

import (
	"encoding/json"
	"fmt"

	"github.com/wyrdnixx/Besucherliste/models"
)

// ConsumeMessage : gets inbound Message and selects type of Message
func ConsumeMessage(_inbound transmitter) models.ResultMessage {

	var _msg = _inbound.Message
	var m models.InMessage
	var result models.ResultMessage

	err := json.Unmarshal(_msg, &m)
	if err != nil {
		fmt.Printf("Error decoding Message : %s", err.Error())
		result.Type = "Error"
		result.Info = err.Error()
		return result
	} else {

		fmt.Printf("got message type: %s\n", m.Type)
		fmt.Printf("got message data: %s\n", m.Data)
		result.Type = "Success"
		result.Info = "Message unmarshall successfull: "

		switch m.Type {
		case "NewVisitor":
			fmt.Printf("NewVisitor requested\n")
			result = processRewNewVisitor(m.Data)

		case "UpdateVisitor":
			fmt.Printf("UpdateVisitor requested\n")
			result = processUpdateVisitor(m.Data)
		}

		return result
	}

}

func processRewNewVisitor(_m interface{}) models.ResultMessage {
	var nv models.ReqNewVisitor
	var result models.ResultMessage

	// convert back to []byte
	x, _ := json.Marshal(_m)
	// unmarshal to type struct
	err := json.Unmarshal(x, &nv)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		result.Type = "Error"
		result.Info = err.Error()
	} else {
		fmt.Printf("nv.Surname: %s\n", nv.Surname)
		res, err := InsertVisitor(nv)
		if err != nil {
			result.Type = "Error"
			result.Info = err.Error()
		} else {
			result.Type = "Success"
			result.Info = "Visitor created result: " + res
		}
	}

	return result
}

func processUpdateVisitor(_m interface{}) models.ResultMessage {
	var uv models.ReqUpdVisitor
	var result models.ResultMessage

	// convert back to []byte
	x, _ := json.Marshal(_m)
	// unmarshal to type struct
	err := json.Unmarshal(x, &uv)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		result.Type = "Error"
		result.Info = err.Error()
	} else {
		res, err := UpdateVisitor(uv)
		if err != nil {
			result.Type = "Error"
			result.Info = err.Error()
		} else {
			result.Type = "Success"
			result.Info = "Update Visitor result: " + res
		}
	}

	return result
}

// ConsumeMessageBAK : getting inbound message and processing
func ConsumeMessageBAK(_inbound transmitter) models.ResultMessage {

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
			/* 			res, err := InsertVisitor(m.Data)

			   			if err != nil {
			   				result.Type = "Error"
			   				result.Info = err.Error()
			   			} else {
			   				result.Type = "Success"
			   				result.Info = "Visitor created result: " + res

			   			} */

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

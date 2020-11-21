package modules

import (
	"encoding/json"
	"fmt"
	"strconv"

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
			result = processReqNewVisitor(m.Data, _inbound)

		case "UpdateVisitor":
			fmt.Printf("UpdateVisitor requested\n")
			result = processUpdateVisitor(m.Data)

			// Send Broadcast to all Clients
			var updateInfo models.ResultMessage
			updateInfo.Type = "VisitorUpdate"
			//updateInfo.Info = m.Data
			//go bc(_inbound, updateInfo)

		}

		return result
	}

}

func processReqNewVisitor(_m interface{}, _inboud transmitter) models.ResultMessage {
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
			i := strconv.FormatInt(res, 10)
			result.Info = i

			// Test get visitor by ID
			v, err := GetVisitorById(res)
			if err != nil {
			} else {
				fmt.Printf("Inserted Result: %s", v)
				go bc(_inboud, v)
			}

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

func bc(_inbound transmitter, _data models.Visitor) {
	fmt.Printf("in bc func\n")
	var x, _ = json.Marshal(_data)
	//_inbound.Client.hub.broadcast <- []byte(_data)
	_inbound.Client.hub.broadcast <- x
	fmt.Printf("end bc func\n")
}

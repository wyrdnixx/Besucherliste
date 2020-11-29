package testing

import (
	"testing"

	"github.com/wyrdnixx/Besucherliste/models"
	"github.com/wyrdnixx/Besucherliste/modules"
)

func TestUpdate(t *testing.T) {

	modules.ReadConfig()

	var updVisitorOK models.ReqUpdVisitor
	updVisitorOK.ID = "58"
	updVisitorOK.Surname = "UpdatedHans"
	updVisitorOK.Givenname = "UpdatedEberdinger"
	updVisitorOK.Birthd = "1999-12-01"

	var updVisitorFail models.ReqUpdVisitor
	updVisitorFail.ID = "-1"
	updVisitorFail.Surname = "UpdatedHans"
	updVisitorFail.Givenname = "UpdatedEberdinger"
	updVisitorFail.Birthd = "2001-12-01"

	//okVisitor, _ := json.Marshal(updVisitorOK)
	//failVisitor, _ := json.Marshal(updVisitorFail)
	//fmt.Printf(string(teststring))

	var _, errOk = modules.UpdateVisitor(updVisitorOK)
	var _, errFail = modules.UpdateVisitor(updVisitorFail)

	if errOk != nil {
		t.Errorf("Test-Error: %v", errOk.Error())
	}
	if errFail == nil {
		t.Errorf("Error-Visitor returned other nil - should have been error: \n")
	}

}

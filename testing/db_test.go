package testing

import (
	"testing"

	"github.com/wyrdnixx/Besucherliste/models"
	"github.com/wyrdnixx/Besucherliste/modules"
)

func init() {
	modules.ReadConfig()

}

func TestCreate(t *testing.T) {
	var insVisitorOK models.ReqNewVisitor
	insVisitorOK.Surname = "InsertedHans"
	insVisitorOK.Givenname = "InsertedMueller"
	insVisitorOK.Birthd = "2001-12-01"
	var _, errIns = modules.InsertVisitor(insVisitorOK)

	if errIns != nil {
		t.Errorf("Error on insert new visitor: %v \n", errIns.Error())
	}
}

func TestFailedTest(t *testing.T) {
	//var i = 2
	//t.Errorf("Test-Testerror: %v", i)

}
func TestUpdate(t *testing.T) {

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

	//okVisitor, _ := json.Mar´shal(updVisitorOK)
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

func TestDelete(t *testing.T) {

	var errIns = modules.DeleteVisitor(2)

	if errIns != nil {
		t.Errorf("Error on delete new visitor: %v \n", errIns.Error())
	}
}

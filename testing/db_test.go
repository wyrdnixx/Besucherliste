package testing

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/wyrdnixx/Besucherliste/models"
	"github.com/wyrdnixx/Besucherliste/modules"
)

func TestUpdate(t *testing.T) {

	modules.ReadConfig()

	var updVisitor models.ReqUpdVisitor

	updVisitor.ID = "5"
	updVisitor.Surname = "UpdatedHans"
	updVisitor.Givenname = "UpdatedEberdinger"
	updVisitor.Birthd = "1999-12-01"

	teststring, _ := json.Marshal(updVisitor)

	fmt.Printf(string(teststring))

	var r, err = modules.UpdateVisitor(updVisitor)

	fmt.Printf("res %v\n", r)
	if err != nil {
		t.Errorf("tot != 1")
	}
}

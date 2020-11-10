package modules

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wyrdnixx/Besucherliste/models"
)

func testfunc(m models.MessageData) (string, error) {
	fmt.Printf("Testfunktion")

	dbInfo := AppConfig.DBUser + ":" + AppConfig.DBPassword + "@tcp(" + AppConfig.DBHost + ":" + AppConfig.DBPort + ")/" + AppConfig.DBName

	fmt.Printf("Mysql-DB connection: %s\n", dbInfo)

	db, err := sql.Open("mysql", dbInfo)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var version string
	type newVisitor struct {
		id        int
		Surname   string
		Givenname string
		Birthd    string
	}
	var nv newVisitor

	//err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	var sql = "select * from Visitors;"
	//var sql = "INSERT INTO `mydb`.`Visitors` (`Surname`, `Givenname`, `Bithd`) VALUES ('" + m.Surname + "', '" + m.Givenname + "', '" + m.Birthd + "');"

	//	err2 := db.QueryRow(sql).Scan(&version)
	err2 := db.QueryRow(sql).Scan(&nv)

	fmt.Printf("testfunc got: %s\n", m.Surname)

	if err2 != nil {
		fmt.Printf("Error-Mysql: %s\n", err2)
		return "", err2
		//log.Fatal(err2)
	}

	fmt.Println(nv)
	return version, nil
}

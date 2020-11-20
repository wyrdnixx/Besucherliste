package modules

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wyrdnixx/Besucherliste/models"
)

func InsertVisitor(m models.MessageData) (string, error) {
	fmt.Printf("InsertVisitor")
	fmt.Printf("Database Connection parameters: %s \n", models.DBInfo)

	db, err := sql.Open("mysql", models.DBInfo)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var sql = "INSERT INTO `mydb`.`visitors` (`Surname`, `Givenname`, `Birthd`) VALUES ('" + m.Surname + "', '" + m.Givenname + "', '" + m.Birthd + "');"

	result, err := db.Query(sql)

	if err != nil {
		fmt.Printf("Error-Mysql: %s\n", err)
		return "", err
		//log.Fatal(err2)
	}
	fmt.Printf("Insert sucessfully id: %s \n", result)
	return "Success", nil
}

func testfunc(m models.MessageData) (string, error) {
	fmt.Printf("Testfunktion")
	dbInfo := AppConfig.DBUser + ":" + AppConfig.DBPassword + "@tcp(" + AppConfig.DBHost + ":" + AppConfig.DBPort + ")/" + AppConfig.DBName

	fmt.Printf("Mysql-DB connection: %s\n", dbInfo)

	db, err := sql.Open("mysql", dbInfo)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	//TESTS
	var version string
	/* 	type Visitor struct {
		id        int
		Surname   string
		Givenname string
		Birthd    string
	} */

	//err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	var sql = "select * from visitors;"
	//var sql = "INSERT INTO `mydb`.`Visitors` (`Surname`, `Givenname`, `Bithd`) VALUES ('" + m.Surname + "', '" + m.Givenname + "', '" + m.Birthd + "');"

	//	err2 := db.QueryRow(sql).Scan(&version)
	results, err2 := db.Query(sql)

	fmt.Printf("testfunc got: %s\n", m.Surname)

	if err2 != nil {
		fmt.Printf("Error-Mysql: %s\n", err2)
		return "", err2
		//log.Fatal(err2)
	}

	for results.Next() {
		var v models.Visitor
		err = results.Scan(&v.Id, &v.Surname, &v.Givenname, &v.Birthd)
		if err != nil {
			fmt.Printf("Error on sql select: %s", err.Error())
		} else {
			log.Printf("SQL: %v;%s;%s;%s", v.Id, v.Surname, v.Givenname, v.Birthd)
			log.Printf("SQL: %v", v)
			log.Printf("SQL: %+v", v)

		}
	}

	//fmt.Println(nv)
	return version, nil
}

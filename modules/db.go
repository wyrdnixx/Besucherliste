package modules

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wyrdnixx/Besucherliste/models"
)

func InsertVisitor(m models.ReqNewVisitor) (int64, error) {
	fmt.Printf("InsertVisitor")
	fmt.Printf("Database Connection parameters: %s \n", models.DBInfo)

	db, err := sql.Open("mysql", models.DBInfo)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var sql = "INSERT INTO `mydb`.`visitors` (`Surname`, `Givenname`, `Birthd`, `chd`) VALUES ('" + m.Surname + "', '" + m.Givenname + "', '" + m.Birthd + "', NOW() );"

	result, err := db.Exec(sql)

	if err != nil {
		fmt.Printf("Error-Mysql: %s\n", err)
		return -1, err
		//log.Fatal(err2)
	}
	fmt.Printf("Insert sucessfully id: %s \n", result)
	id, _ := result.LastInsertId()
	fmt.Printf("Insert sucessfully id: %v \n", id)
	return id, nil
}

func UpdateVisitor(m models.ReqUpdVisitor) (string, error) {
	fmt.Printf("UpdateVisitor")
	fmt.Printf("Database Connection parameters: %s \n", models.DBInfo)

	db, err := sql.Open("mysql", models.DBInfo)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var sql = "UPDATE `mydb`.`visitors` SET `Surname` = '" + m.Surname + "', `Givenname` = '" + m.Givenname + "', `Birthd` = '" + m.Birthd + "', `chd`= NOW() WHERE (`id` = '" + m.ID + "');"
	fmt.Printf("SQL: %s \n", sql)

	//result, err := db.Query(sql)
	result, err := db.Exec(sql)

	if err != nil {
		fmt.Printf("Error-Mysql: %s\n", err)
		return "", err
		//log.Fatal(err2)
	}
	fmt.Printf("Udpate Result: %v", result)
	return "Success", nil
}

func GetVisitorById(_i int64) (models.Visitor, error) {

	db, err := sql.Open("mysql", models.DBInfo)
	var v models.Visitor

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	i := strconv.FormatInt(_i, 10)

	var sql = "select * from mydb.visitors where id = " + i + ";"
	fmt.Printf("SQL: %s \n", sql)
	result, err := db.Query(sql)

	if err != nil {
		fmt.Printf("Error-Mysql: %s\n", err)
		return v, err
	}
	for result.Next() {

		err = result.Scan(&v.ID, &v.Surname, &v.Givenname, &v.Birthd, &v.Chd)
		if err != nil {
			fmt.Printf("Error getting visitor: %s \n", err.Error())

		} else {
			fmt.Printf("Found visitor: %s \n", v.Surname)

		}
	}

	return v, nil
}

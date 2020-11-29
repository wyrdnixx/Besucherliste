package modules

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wyrdnixx/Besucherliste/models"
)

func InsertVisitor(m models.ReqNewVisitor) (int, error) {
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
	x, _ := result.LastInsertId()
	id := int(x)
	fmt.Printf("Insert sucessfully id: %v \n", id)
	return id, nil
}

func UpdateVisitor(m models.ReqUpdVisitor) (int, error) {
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
		return -1, err
		//log.Fatal(err2)
	}

	count, _ := result.RowsAffected()
	id, _ := strconv.Atoi(m.ID)

	if count != 1 {

		return -1, fmt.Errorf("Error updating visitor id %v returned no result\n", id)
	} else {
		return id, nil
	}

	/*
		if id, err := strconv.Atoi(m.ID); err == nil {
			return id, nil

		} else {
			fmt.Println(m.ID, "is not an integer.")
			return -1, nil

		}
	*/
}

func GetVisitorById(_i int) (models.Visitor, error) {

	db, err := sql.Open("mysql", models.DBInfo)
	var v models.Visitor

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	i := strconv.Itoa(_i)

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

func GetAllVisitors() (models.AllVisitors, error) {
	// func GetAllVisitors() {
	db, err := sql.Open("mysql", models.DBInfo)
	var allVisitors models.AllVisitors

	fmt.Printf("AllV: %v \n", allVisitors)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var sql = "select * from mydb.visitors ;"
	fmt.Printf("SQL: %s \n", sql)
	result, err := db.Query(sql)

	if err != nil {
		fmt.Printf("Error-Mysql: %s\n", err)
		return allVisitors, err
	}
	for result.Next() {
		var v models.Visitor
		err = result.Scan(&v.ID, &v.Surname, &v.Givenname, &v.Birthd, &v.Chd)
		if err != nil {
			fmt.Printf("Error getting visitor: %s \n", err.Error())

		} else {
			//fmt.Printf("Found visitor: %v:%s \n", v.ID, v.Surname)
			allVisitors.Visitors = append(allVisitors.Visitors, v)
			//visitors.Visitor = append(v)
		}
	}

	//fmt.Printf("AllV: %v \n", allVisitors)

	return allVisitors, nil
}

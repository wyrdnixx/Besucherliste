package modules

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func testfunc() {
	fmt.Printf("Testfunktion")

	dbInfo := AppConfig.DBUser + ":" + AppConfig.DBPassword + "@tcp(" + AppConfig.DBHost + ":" + AppConfig.DBPort + ")/" + AppConfig.DBName

	fmt.Printf("Mysql-DB connection: %s\n", dbInfo)

	db, err := sql.Open("mysql", dbInfo)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		fmt.Printf("Error-Mysql: %s\n", err2)

		log.Fatal(err2)
	}

	fmt.Println(version)

}

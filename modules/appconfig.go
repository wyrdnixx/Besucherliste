package modules

import (
	"fmt"

	"github.com/tkanos/gonfig"
	"github.com/wyrdnixx/Besucherliste/models"
)

// Configuration -  Allgemeine Config
type Configuration struct {
	SrvPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Info       string
}

// AppConfig Application Configurations
var AppConfig = Configuration{}

// ReadConfig Lies die Config aus der Config File aus
func ReadConfig() {

	// Config aus file laden:

	AppConfig.Info = "TestInfoComment"

	err := gonfig.GetConf("./config.development.json", &AppConfig)
	if err != nil {
		fmt.Printf("ERROR: Could not read config: %v\n ", err.Error())
		fmt.Printf("Trying second file location\n")
		err2 := gonfig.GetConf("../config.development.json", &AppConfig)
		if err2 != nil {
			fmt.Printf("ERROR: Config second location also failed: %v\n ", err.Error())
		}
	}

	//fmt.Println("Info: DBHost: ", AppConfig.DBHost)
	//fmt.Println("Info: DBPort: ", AppConfig.DBPort)

	models.DBInfo = AppConfig.DBUser + ":" + AppConfig.DBPassword + "@tcp(" + AppConfig.DBHost + ":" + AppConfig.DBPort + ")/" + AppConfig.DBName
	//DBInfo := AppConfig.DBUser + ":" + AppConfig.DBPassword + "@tcp(" + AppConfig.DBHost + ":" + AppConfig.DBPort + ")/" + AppConfig.DBName
	//fmt.Printf("got DBInfo: %s", models.DBInfo)
}

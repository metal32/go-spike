package main

import (
	"fmt"
	"todoexample.com/Config"
	"todoexample.com/Models"
	"todoexample.com/Routes"
	"github.com/jinzhu/gorm"
	u "todoexample.com/Utils"
)

var err error

func main() {

	//Checking Log system
	u.GeneralLogger.Println("Starting..")

	// Execution od Code
	u.GeneralLogger.Println("Running Process 1")
	u.GeneralLogger.Println("Running Process 2")

	u.ErrorLogger.Println("An Error Occured: Error abcd")

	u.GeneralLogger.Println("Running Process i")
	u.ErrorLogger.Println("An Error Occured: Error efgh")
	u.GeneralLogger.Println("Running Process n")

	// Creating a connection to the database
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()

	// run the migrations: todo struct
	Config.DB.AutoMigrate(&Models.Todo{})

	//setup routes
	r := Routes.SetupRouter()

	// running
	r.Run()
}
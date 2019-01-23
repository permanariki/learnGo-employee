package main

import (
	"os"

	"github.com/golang/employee/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	envFile := os.Getenv("ENV")
	if envFile == "" {
		envFile = ".env"
	}

	//Load env data
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

	// Create Database
	// db, err := gorm.Open("postgres", os.Getenv("DBCONN"))
	// defer db.Close()
	// tableEmployees := new(handler.Employees)

	// err = tableEmployees.CreateTable(db)
	// if err != nil {
	// 	panic(err)
	// }

	// Endpoint
	e.POST("/employees", handler.CreateData)
	e.GET("/employees", handler.GetData)
	e.GET("/getAll", handler.GetAllData)
	e.PUT("/employees", handler.UpdateData)
	e.DELETE("/employees", handler.DeleteData)

	//Start echo on desired port
	e.Start(":" + os.Getenv("PORT"))
}

package handler

import (
	"github.com/labstack/echo"
)

// CreateData - Input data to Employees table
func CreateData(c echo.Context) error {
	inputData := new(Employees)

	if err := c.Bind(inputData); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer db.Close()

	err = inputData.InsertData(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(201, inputData)
}

// GetData - Get data by Name
func GetData(c echo.Context) error {
	// Get input from query param
	name := c.QueryParam("name")

	// Initiate data from table Employees
	data := Employees{
		Name: name,
	}

	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer db.Close()

	// Query data to database using method from Employees
	response, err := data.GetDataByName(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	return c.JSON(200, response)
}

// GetAllData - to get all data
func GetAllData(c echo.Context) error {
	var d Employees

	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer db.Close()

	res, err := d.GetAllData(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	return c.JSON(200, res)
}

// UpdateData - to update data by id
func UpdateData(c echo.Context) error {
	// Get input from query param
	name := c.QueryParam("name")

	// Initiate data from table Employees
	data := Employees{
		Name: name,
	}

	inputData := new(Employees)

	if err := c.Bind(inputData); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer db.Close()

	// Query data to databse using method from Employees
	response, err := data.UpdateData(db, inputData)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	return c.JSON(200, response)
}

// DeleteData -  delete data by name
func DeleteData(c echo.Context) error {
	// Get input from query param
	name := c.QueryParam("name")

	// Initiate data from table Employees
	data := Employees{
		Name: name,
	}

	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer db.Close()

	// Query data to databse using method from Employees
	err = data.DeleteData(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	return c.JSON(200, "Deleted")
}

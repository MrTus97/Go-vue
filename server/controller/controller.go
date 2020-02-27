package controller

import (
	"echo-example-crud/server/database"
	"echo-example-crud/server/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllUser(c echo.Context) error {
	// Get all user in db
	db := database.CreateCon()
	defer db.Close()

	row, err := db.Query("Select * from employee")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "500", nil)
	}
	defer row.Close()

	each := models.Employee{}
	result := []models.Employee{}

	for row.Next() {
		var id string
		var name string

		err2 := row.Scan(&id, &name)

		if err2 != nil {
			return c.Render(http.StatusInternalServerError, "500", nil)
		}

		each.Id = id
		each.Name = name

		result = append(result, each)
	}
	fmt.Println(result)
	return c.JSON(http.StatusOK, result)
}

func AddEmployee(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	emp := new(models.Employee)
	if err := c.Bind(emp); err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	sql := "INSERT INTO employee(Name) VALUES(?)"
	stmt, err := db.Prepare(sql)

	if err != nil {
		return c.Render(http.StatusInternalServerError, "500", nil)
	}

	defer stmt.Close()
	_, err2 := stmt.Exec(emp.Name)

	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "500", nil)
	}
	return c.JSON(http.StatusOK, "OK")

}

func EditEmployee(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	request_id := c.Param("id")
	request_Name := c.FormValue("Name")

	stmt, err := db.Prepare("UPDATE employee SET Name = ? where Id = ?")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "500", nil)
	}

	_, err2 := stmt.Exec(request_Name, request_id)

	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "500", nil)
	}
	return c.JSON(http.StatusOK, "OK")
}

func DeleteEmployee(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()
	requestDelete := c.Param("id")

	stmt, err := db.Prepare("DELETE FROM employee where Id= ?")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "500", nil)
	}

	_, err2 := stmt.Exec(requestDelete)
	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "500", nil)
	}

	return c.JSON(http.StatusOK, "OK")
}

func GetEmployee(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()
	requestDelete := c.Param("id")

	stmt, err := db.Prepare("SELECT * FROM employee where id=?")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "500", nil)
	}

	emp, err2 := stmt.Query(requestDelete)
	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "500", nil)
	}
	each := models.Employee{}
	if emp.Next() {
		var id, name string

		err2 := emp.Scan(&id, &name)

		if err2 != nil {
			return c.Render(http.StatusInternalServerError, "500", nil)
		}

		each.Id = id
		each.Name = name
	}

	return c.JSON(http.StatusOK, each)
}

func TestIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

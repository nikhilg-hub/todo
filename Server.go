package main

import (
	"log"
	"net/http"

	"github.com/nikhilg-hub/todo/ToDoBackend/orm"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/toDos/get", getAll)
	e.POST("/toDos/add", addOne)

	// Start server
	e.Logger.Fatal(e.Start(":80"))

}

func getAll(c echo.Context) error {
	var list []orm.ToDo
	DB := orm.GetDatabase()
	DB.Find(&list)
	return c.JSON(http.StatusOK, list)
}

func addOne(c echo.Context) error {
	toAdd := new(orm.ToDo)
	err := c.Bind(toAdd)
	if err != nil {
		log.Println(err)
	}
	DB := orm.GetDatabase()
	DB.Create(toAdd)
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Successfully Added", "Id": toAdd.Id})
}

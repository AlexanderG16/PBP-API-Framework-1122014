package routes

import (
	"Echo/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	res, err := controllers.GetAllUsers()
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"message": err.Error()}, "	")
	}

	return c.JSONPretty(res.Status, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, " ")
}

func GetUserByID(c echo.Context) error {
	res, err := controllers.GetUserByID(c.Param("id"))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"message": err.Error()}, "	")
	}

	return c.JSONPretty(res.Status, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, " ")
}

func InsertUser(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	gender := c.QueryParam("gender")
	if name == "" || age == "" || age == "0" || gender == "" {
		return c.JSONPretty(http.StatusBadRequest, map[string]string{"message": "Invalid Variables"}, "	")
	}

	res, _ := controllers.InsertUser(name, age, gender)

	return c.JSONPretty(res.Status, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, " ")
}

func UpdateUser(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "0" || id == "" {
		return c.JSONPretty(http.StatusBadRequest, map[string]string{"message": "Invalid ID"}, "	")
	}
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	gender := c.QueryParam("gender")
	if name == "" || age == "" || age == "0" || gender == "" {
		return c.JSONPretty(http.StatusNotImplemented, map[string]string{"message": "Variables Missing"}, "	")
	}

	res, _ := controllers.UpdateUser(id, name, age, gender)

	return c.JSONPretty(res.Status, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, "	")
}

func DeleteUser(c echo.Context) error {
	res, err := controllers.DeleteUser(c.Param("id"))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"message": err.Error()}, "	")
	}

	return c.JSONPretty(res.Status, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, "	")
}

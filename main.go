package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var name []string
var roll []int
var branch []string
var userID []string

func home(c echo.Context) error {
	return c.String(http.StatusOK,
		"Please go the URL /add to add a student entry.\n"+
			"Please go the URL /add to add a student entry.\n"+
			"Please go the URL /find to find a student entry.\n"+
			"Please go the URL /edit to edit a student entry.\n")
}
func addget(c echo.Context) error {
	return c.String(http.StatusOK, "Please enter the form value as Name, Roll No. ,Branch,UserID")
}
func addpost(c echo.Context) error {
	Name := c.FormValue("Name")
	RollN := c.FormValue("Roll No.")
	Branch := c.FormValue("Branch")
	UserID := c.FormValue("UserID")
	Roll, e := strconv.Atoi(RollN)
	if e != nil {
		return c.String(http.StatusBadRequest, "Invalid String")
	}
	name = append(name, Name)
	roll = append(roll, Roll)
	branch = append(branch, Branch)
	userID = append(userID, UserID)
	return nil
}
func deleteget(c echo.Context) error {
	return c.String(http.StatusOK, "Please enter the name whose entry is to be deleted.")
}
func deletepost(c echo.Context) error {
	Name := c.FormValue("Name")
	f := false
	for i, v := range name {
		if v == Name {
			name = append(name[:i], name[i+1:]...)
			roll = append(roll[:i], roll[i+1:]...)
			branch = append(branch[:i], branch[i+1:]...)
			userID = append(userID[:i], userID[i+1:]...)
			f = true
		}
	}
	if f == false {
		return c.String(http.StatusBadRequest, "Student not found.")
	} else {
		return c.String(http.StatusOK, "Student entry deleted.")
	}
}
func main() {
	e := echo.New()
	e.GET("/", home)
	e.GET("/add", addget)
	e.POST("/add", addpost)
	e.GET("/delete", deleteget)
	e.POST("/delete", deletepost)
	//e.GET("/edit", edit)
	//e.GET("/find", find)
}

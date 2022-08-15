package main

import (
	"fmt"
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
		"Please go to the URL /add to add a student entry.\n"+
			"Please go to the URL /add to add a student entry.\n"+
			"Please go to the URL /find to find a student entry.\n"+
			"Please go to the URL /edit to edit a student entry.\n")
}
func addGet(c echo.Context) error {
	return c.String(http.StatusOK,
		"Please enter Name, Roll No. ,Branch,UserID")
}
func addPost(c echo.Context) error {
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
	return c.String(http.StatusOK, "Entry added")
}
func deleteGet(c echo.Context) error {
	return c.String(http.StatusOK,
		"Please enter the name whose entry is to be deleted.")
}
func deletePost(c echo.Context) error {
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
		return c.String(http.StatusBadRequest,
			"Student not found.")
	} else {
		return c.String(http.StatusOK, "Student entry deleted.")
	}
}
func editGet(c echo.Context) error {
	return c.String(http.StatusOK,
		"Please enter the Name,  parameter and  change")
}
func editPost(c echo.Context) error {
	Name := c.FormValue("Name")
	Parameter := c.FormValue("Parameter")
	Change := c.FormValue("Change")
	f := false
	for i, v := range name {
		if v == Name {
			f = true
			if Parameter == "Name" {
				n := append(name[i+1:])
				name = append(name[:i], Change)
				name = append(name, n...)
			} else if Parameter == "Roll No." {
				Roll, e := strconv.Atoi(Change)
				if e != nil {
					return c.String(http.StatusBadRequest,
						"Invalid Roll No.")
				}
				n := append(roll[i+1:])
				roll = append(roll[:i], Roll)
				roll = append(roll, n...)
			} else if Parameter == "Branch" {
				n := append(branch[i+1:])
				branch = append(branch[:i], Change)
				branch = append(branch, n...)
			} else if Parameter == "UserID" {
				n := append(userID[i+1:])
				userID = append(userID[:i], Change)
				userID = append(userID, n...)
			} else {
				return c.String(http.StatusBadRequest,
					"Please enter a valid parameter")
			}
		}
	}
	if f == false {
		return c.String(http.StatusBadRequest,
			"No such student found")
	} else {
		return c.String(http.StatusOK, "Student Entry edited.")
	}
}
func findGet(c echo.Context) error {
	return c.String(http.StatusOK,
		"Enter the parameter and its Value to be searched")
}
func findPost(c echo.Context) error {
	Parameter := c.FormValue("Parameter")
	Value := c.FormValue("Value")
	f := false
	if Parameter == "Name" {
		for i, v := range name {
			if Value == v {
				f = true
				c.String(http.StatusOK, fmt.Sprintf(
					"Name : %v\tRoll No. : %v\tBranch : %v\tUserID : %v",
					name[i], roll[i], branch[i], userID[i]))
			}
		}
		if f == false {
			return c.String(http.StatusBadRequest,
				"No such student found")
		} else {
			return c.String(http.StatusOK, "Done")
		}
	} else if Parameter == "Roll No." {
		Roll, e := strconv.Atoi(Value)
		if e != nil {
			return c.String(http.StatusBadRequest,
				"Please enter a valid roll no.")
		}
		for i, v := range roll {
			if Roll == v {
				f = true
				c.String(http.StatusOK, fmt.Sprintf(
					"Name : %v\tRoll No. : %v\tBranch : %v\tUserID : %v",
					name[i], roll[i], branch[i], userID[i]))
			}
		}
		if f == false {
			return c.String(http.StatusBadRequest,
				"No such student found")
		} else {
			return c.String(http.StatusOK, "Done")
		}
	} else if Parameter == "Branch" {
		for i, v := range branch {
			if Value == v {
				f = true
				c.String(http.StatusOK, fmt.Sprintf(
					"Name : %v\tRoll No. : %v\tBranch : %v\tUserID : %v",
					name[i], roll[i], branch[i], userID[i]))
			}
		}
		if f == false {
			return c.String(http.StatusBadRequest,
				"No such student found")
		} else {
			return c.String(http.StatusOK, "Done")
		}
	} else if Parameter == "UserID" {
		for i, v := range userID {
			if Value == v {
				f = true
				c.String(http.StatusOK, fmt.Sprintf(
					"Name : %v\tRoll No. : %v\tBranch : %v\tUserID : %v",
					name[i], roll[i], branch[i], userID[i]))
			}
		}
		if f == false {
			return c.String(http.StatusBadRequest,
				"No such student found")
		} else {
			return c.String(http.StatusOK, "Done")
		}
	} else {
		return c.String(http.StatusBadRequest,
			"Please enter a valid parameter")
	}
}
func main() {
	e := echo.New()
	e.GET("/", home)
	e.GET("/add", addGet)
	e.POST("/add", addPost)
	e.GET("/delete", deleteGet)
	e.POST("/delete", deletePost)
	e.GET("/edit", editGet)
	e.POST("/edit", editPost)
	e.GET("/find", findGet)
	e.POST("/find", findPost)
	e.Logger.Fatal(e.Start(":4000"))
}

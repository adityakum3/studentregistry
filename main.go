package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strconv"
	"time"
)

var StudentData *mongo.Collection
var ctx context.Context
var AddResult *mongo.InsertOneResult
var Result *mongo.UpdateResult
var Delete *mongo.DeleteResult

//var name []string
//var roll []int
//var branch []string
//var userID []string
const (
	IncorrectError = "Please enter a correct roll no."
	Done           = "Done"
	NotFound       = "Student not found"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK,
		"Please go to the URL /add to add a student entry.\n"+
			"Please go to the URL /add to add a student entry.\n"+
			"Please go to the URL /find to find a student entry.\n"+
			"Please go to the URL /edit to edit a student entry.\n")
}
func addGet(c echo.Context) error {
	//c.HTML(http.StatusOK, "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta http-equiv=\"X-UA-Compatible\" content = \"IE=edge\"/>\n    <meta name = \"viewpoint\" content = \"width=device-width, initial scale = 10\">\n    <title>AddPage</title>\n</head>\n<body>\n<div class = \"container\">\n    <section style = \"padding-top: 15px;\">\n        <form action=\"./add\" method = \"POST\">\n            <div class=\"table-responsive\">\n                <table class = \"table table-bordered table-striped\">\n                    <tbody>\n                    <tr>\n                        <td>\n                            <b>Name</b>\n                        </td>\n                        <td>\n                            <input type = \"submit\" class=\"form control\" name = \"Name\" value = \"\">\n                        </td>\n                    </tr>\n                    <tr>\n                        <td>\n                            <b>Roll No.</b>\n                        </td>\n                        <td>\n                            <input type = \"submit\" class=\"form control\" name = \"Roll No.\" value = \"\">\n                        </td>\n                    </tr>\n                    <tr>\n                        <td>\n                            <b>Branch</b>\n                        </td>\n                        <td>\n                            <input type = \"submit\" class=\"form control\" name = \"Branch\" value = \"\">\n                        </td>\n                    </tr>\n                    <tr>\n                        <td>\n                            <b>UserID</b>\n                        </td>\n                        <td>\n                            <input type = \"submit\" class=\"form control\" name = \"UserID\" value = \"\">\n                        </td>\n                    </tr>\n                    </tbody>\n                </table>\n            </div>\n        </form>\n    </section>\n</div>\n</body>\n</html>")
	return c.String(http.StatusOK,
		"Please enter Name, Roll No. ,Branch,UserID")
}
func addPost(c echo.Context) error {
	Name := c.FormValue("Name")
	RollN := c.FormValue("Roll No.")
	Branch := c.FormValue("Branch")
	UserID := c.FormValue("UserID")
	_, e := strconv.Atoi(RollN)
	if e != nil {
		return c.String(http.StatusBadRequest, IncorrectError)
	}
	AddResult, _ = StudentData.InsertOne(ctx, bson.D{
		{"Name", Name},
		{"Roll", RollN},
		{"Branch", Branch},
		{"UserID", UserID},
	})
	//name = append(name, Name)
	//roll = append(roll, Roll)
	//branch = append(branch, Branch)
	//userID = append(userID, UserID)
	return c.String(http.StatusOK, Done)
}
func deleteGet(c echo.Context) error {
	return c.String(http.StatusOK,
		"Please enter the Category and its value whose entry is to be deleted.")
}
func deletePost(c echo.Context) error {
	Category := c.FormValue("Category")
	Value := c.FormValue("Value")
	var err error
	Delete, err = StudentData.DeleteOne(ctx, bson.M{Category: Value})
	if err != nil {
		return c.String(http.StatusOK, NotFound)
	}
	return c.String(http.StatusOK, "Done")
	//f := false
	//for i, v := range name {
	//	if v == Name {
	//		name = append(name[:i], name[i+1:]...)
	//		roll = append(roll[:i], roll[i+1:]...)
	//		branch = append(branch[:i], branch[i+1:]...)
	//		userID = append(userID[:i], userID[i+1:]...)
	//		f = true
	//	}
	//}
	//if f == false {
	//	return c.String(http.StatusBadRequest,
	//		"Student not found.")
	//} else {
	//	return c.String(http.StatusOK, "Student entry deleted.")
	//}
}
func editGet(c echo.Context) error {
	return c.String(http.StatusOK,
		"Please enter the Find,Value,  parameter and  change")
}
func editPost(c echo.Context) error {
	Find := c.FormValue("Find")
	Value := c.FormValue("Value")
	Parameter := c.FormValue("Parameter")
	Change := c.FormValue("Change")
	var r error
	Result, r = StudentData.UpdateOne(
		ctx,
		bson.M{Find: Value},
		bson.D{
			{"$set", bson.D{{Parameter, Change}}},
		},
	)
	if r != nil {
		return c.String(http.StatusOK, NotFound)
	}
	return c.String(http.StatusOK, Done)
	//f := false
	//for i, v := range name {
	//	if v == Name {
	//		f = true
	//		if Parameter == "Name" {
	//			n := append(name[i+1:])
	//			name = append(name[:i], Change)
	//			name = append(name, n...)
	//		} else if Parameter == "Roll No." {
	//			Roll, e := strconv.Atoi(Change)
	//			if e != nil {
	//				return c.String(http.StatusBadRequest,
	//					"Invalid Roll No.")
	//			}
	//			n := append(roll[i+1:])
	//			roll = append(roll[:i], Roll)
	//			roll = append(roll, n...)
	//		} else if Parameter == "Branch" {
	//			n := append(branch[i+1:])
	//			branch = append(branch[:i], Change)
	//			branch = append(branch, n...)
	//		} else if Parameter == "UserID" {
	//			n := append(userID[i+1:])
	//			userID = append(userID[:i], Change)
	//			userID = append(userID, n...)
	//		} else {
	//			return c.String(http.StatusBadRequest,
	//				"Please enter a valid parameter")
	//		}
	//	}
	//}
	//if f == false {
	//	return c.String(http.StatusBadRequest,
	//		"No such student found")
	//} else {
	//	return c.String(http.StatusOK, "Student Entry edited.")
	//}
}
func findGet(c echo.Context) error {
	return c.String(http.StatusOK,
		"Enter the parameter and its Value to be searched")
}
func findPost(c echo.Context) error {
	Parameter := c.FormValue("Parameter")
	Value := c.FormValue("Value")
	DataCursor, err := StudentData.Find(ctx, bson.M{Parameter: Value})
	if err != nil {
		//log.Fatal(err)
		return c.String(http.StatusOK, "")
	}
	var Data []bson.M
	err = DataCursor.All(ctx, &Data)
	if err != nil {
		return c.String(http.StatusOK, NotFound)
	}
	for _, i := range Data {
		for j, k := range i {
			c.String(http.StatusOK, fmt.Sprintf("%v : %v\n", j, k))
		}

	}
	return c.String(http.StatusOK, Done)
	//f := false
	//if Parameter == "Name" {
	//	for i, v := range name {
	//		if Value == v {
	//			f = true
	//			c.String(http.StatusOK, fmt.Sprintf(
	//				"Name : %v\tRoll No. : %v\tBranch : %v\tUserID : %v",
	//				name[i], roll[i], branch[i], userID[i]))
	//		}
	//	}
	//	if f == false {
	//		return c.String(http.StatusBadRequest,
	//			"No such student found")
	//	} else {
	//		return c.String(http.StatusOK, "Done")
	//	}
	//} else if Parameter == "Roll No." {
	//	Roll, e := strconv.Atoi(Value)
	//	if e != nil {
	//		return c.String(http.StatusBadRequest,
	//			"Please enter a valid roll no.")
	//	}
	//	for i, v := range roll {
	//		if Roll == v {
	//			f = true
	//			c.String(http.StatusOK, fmt.Sprintf(
	//				"Name : %v\tRoll No. : %v\tBranch : %v\tUserID : %v",
	//				name[i], roll[i], branch[i], userID[i]))
	//		}
	//	}
	//	if f == false {
	//		return c.String(http.StatusBadRequest,
	//			"No such student found")
	//	} else {
	//		return c.String(http.StatusOK, "Done")
	//	}
	//} else if Parameter == "Branch" {
	//	for i, v := range branch {
	//		if Value == v {
	//			f = true
	//			c.String(http.StatusOK, fmt.Sprintf(
	//				"Name : %v\tRoll No. : %v\tBranch : %v\tUserID : %v",
	//				name[i], roll[i], branch[i], userID[i]))
	//		}
	//	}
	//	if f == false {
	//		return c.String(http.StatusBadRequest,
	//			"No such student found")
	//	} else {
	//		return c.String(http.StatusOK, "Done")
	//	}
	//} else if Parameter == "UserID" {
	//	for i, v := range userID {
	//		if Value == v {
	//			f = true
	//			c.String(http.StatusOK, fmt.Sprintf(
	//				"Name : %v\tRoll No. : %v\tBranch : %v\tUserID : %v\n",
	//				name[i], roll[i], branch[i], userID[i]))
	//		}
	//	}
	//	if f == false {
	//		return c.String(http.StatusBadRequest,
	//			"No such student found")
	//	} else {
	//		return c.String(http.StatusOK, "Done")
	//	}
	//} else {
	//	return c.String(http.StatusBadRequest,
	//		"Please enter a valid parameter")
	//}
}
func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Aditmoe:MongoDBA@cluster0.dmye4md.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Minute)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	MyDataBase := client.Database("PClub")
	StudentData = MyDataBase.Collection("student")
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

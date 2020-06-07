package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Cab struct {
	vid      int    `json:"vid"`
	vname    string `json:"vname"`
	location int    `json:"location"`
}

/*type User struct {
	uid     int        `json:"uid"`
	uname   string     `json:"uname"`
	ubooking []Bookings `json:"ubooking"`
}*/

type Bookings struct {
	uid      int    `json:"uid"`
	bid      int    `json:"bid"`
	startLoc string `json:"startLoc"`
	endLoc   string `json:"endLoc"`
}

func homePage(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hello world",
	})

	fmt.Println("Endpoint Hit: homePage")
}

func nearBy(c *gin.Context) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/travel")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected")

	key := c.Param("location")
	ki, _ := strconv.Atoi(key)
	kb, kt := ki-3, ki+3

	results, err := db.Query("SELECT vname, location, vid FROM cabs WHERE location BETWEEN ? AND ?", kb, kt)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var cab Cab

		err = results.Scan(&cab.vname, &cab.location, &cab.vid)

		if err != nil {
			panic(err.Error())
		}

		c.JSON(200, gin.H{
			"vname":            cab.vname,
			"location":         cab.location,
			"vid":              cab.vid,
			"current_Location": key,
		})

	}

	defer results.Close()
	defer db.Close()

}

func myBookings(c *gin.Context) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/travel")

	if err != nil {
		panic(err.Error())
	}

	uid := c.Param("uid")
	results, err := db.Query("SELECT uid, bid, startLoc,endLoc FROM bookings WHERE uid=?", uid)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var book Bookings

		err = results.Scan(&book.uid, &book.bid, &book.endLoc, &book.endLoc)

		if err != nil {
			panic(err.Error())
		}

		fmt.Print(book.uid)
		fmt.Print(book.bid)
		fmt.Print(book.startLoc)
		fmt.Print(book.endLoc)

		c.JSON(200, gin.H{
			"uid":      book.uid,
			"bid":      book.bid,
			"startLoc": book.startLoc,
			"endLoc":   book.endLoc,
		})
	}

	defer results.Close()
	defer db.Close()

}

func bookCab(c *gin.Context) {

	fmt.Println("Endpoint Hit: book a cab")

	startLoc := c.Param("startLoc")
	uid := c.Param("uid")
	endLoc := c.Param("endLoc")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/travel")

	if err != nil {
		panic(err.Error())
	}

	_, errQ := db.Query("INSERT INTO bookings(uid, startLoc, endLoc) VALUES (?,?,?)", uid, startLoc, endLoc)

	if errQ != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":            uid,
		"Start_Location": startLoc,
		"End_Location":   endLoc,
	})

}

func main() {

	r := gin.Default()

	r.GET("/", homePage)
	r.GET("/nearBy/:location", nearBy)
	r.POST("/myBookings/:uid", myBookings)
	r.POST("/bookCab", bookCab)

	r.Run()

}

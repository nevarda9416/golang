package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"                  // Gorilla mux: For creating routes and HTtP handlers
	"github.com/jinzhu/gorm"                  // // An ORM tool for MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql" // The MySQL Driver
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var db *gorm.DB
var err error

type Booking struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Homepage!")
	fmt.Println("Endpoint Hit: Homepage")
}

// Create a new booking
func createNewBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create new booking")
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking Booking
	json.Unmarshal(reqBody, &booking)
	db.Create(&booking)
	fmt.Println("Endpoint Hit: Creating New Booking ", r.Body)
	json.NewEncoder(w).Encode(booking)
}

// Reading all bookings
func returnAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings := []Booking{}
	db.Find(&bookings)
	fmt.Println("Endpoint Hit: Reading all bookings")
	json.NewEncoder(w).Encode(bookings)
}

// Reading Booking detail by their Id
func returnSingleBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	bookings := []Booking{}
	db.Find(&bookings)
	for _, booking := range bookings {
		// string to int
		s, err := strconv.Atoi(key)
		if err == nil {
			if booking.Id == s {
				fmt.Println(booking)
				fmt.Println("Endpoint Hit Booking No: ", key)
				json.NewEncoder(w).Encode(booking)
			}
		}
	}
}
func handleRequest() {
	fmt.Println("Starting development server at http://127.0.0.1:10000")
	fmt.Println("Quit the server with CONTROL-C")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/new-booking", createNewBooking).Methods("GET") // POST
	myRouter.HandleFunc("/all-bookings", returnAllBookings)
	myRouter.HandleFunc("/booking/{id}", returnSingleBooking)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func main() {
	// Please define your username and password for MySQL
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True")
	// NOTE: See we're using = to assign the global var
	// instead of := which would assign it only in this function
	if err != nil {
		fmt.Println("Connection failed to open")
	} else {
		fmt.Println("Connection established")
	}
	// create the database. This is a one-time step
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE golang")
	db.Exec("USE golang")
	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&Booking{})
	handleRequest()
}

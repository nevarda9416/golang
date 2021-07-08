// https://medium.com/@hugo.bjarred/rest-api-with-golang-mux-mysql-c5915347fa5b
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux" // Gorilla mux: For creating routes and HTtP handlers
	// // An ORM tool for MySQL
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var db *sql.DB
var err error

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Get all posts")
	var posts []Post
	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO posts(title) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]
	_, err = stmt.Exec(title)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, title FROM posts WHERE id=?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post Post
	for result.Next() {
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE posts SET title=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["title"]
	_, err = stmt.Exec(newTitle, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID=%s was updated", params["id"])
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM posts WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}

func main() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("Connection failed to open")
	} else {
		fmt.Println("Connection established")
	}
	defer db.Close()
	// create the database. This is a one-time step
	// Comment out if running multiple times - You may see an error otherwise
	//db.Exec("CREATE DATABASE golang")
	//db.Exec("USE golang")
	// Migration to create tables for Order and Item schema
	//db.AutoMigrate(&Post{})
	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}

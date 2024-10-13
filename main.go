package main

import (
	"database/sql"
	//	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// type Article struct {
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// type Articles []Article

type usernamejson struct {
	Uname string `json:"username" binding:"required"`
	Pwd   string `json:"password"`
}

// var articles = Articles{
// 	Article{Title: "Test 1", Desc: "Test 1 Desc", Content: "Test Content 1"},
// 	Article{Title: "Test 2", Desc: "Test 2 Desc", Content: "Test Content 2"},
// }

// func allArticles(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("Endpoint Hit: All Articles Endpoint")
// 	json.NewEncoder(w).Encode(articles)
// }

// func allArticles(c *gin.Context) {

// 	c.IndentedJSON(http.StatusOK, articles)
// 	// fmt.Println("Endpoint Hit: All Articles Endpoint")
// 	// json.NewEncoder(w).Encode(articles)
// }

// func allArticles(c *gin.Context) {
// 	articles := []Article{}

// 	rows, err := db.Query("SELECT * FROM articles")
// 	if err != nil {
// 		fmt.Println("Err", err.Error())
// 		return
// 	}

// 	for rows.Next() {
// 		var article Article
// 		var id int
// 		err = rows.Scan(&id, &article.Title, &article.Desc, &article.Content)
// 		if err != nil {
// 			fmt.Println("Err", err.Error())
// 		}
// 		articles = append(articles, article)
// 	}

// 	c.IndentedJSON(http.StatusOK, articles)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Homepage Endpoint Hit")
// }

// func getArticle(c *gin.Context) {
// 	id := c.Param("id")
// 	rows, err := db.Query("SELECT * FROM articles where id=?", id)
// 	if err != nil {
// 		fmt.Println("Err", err.Error())
// 		return
// 	}

// 	if rows.Next() {
// 		var article Article
// 		var id int
// 		err = rows.Scan(&id, &article.Title, &article.Desc, &article.Content)
// 		if err != nil {
// 			fmt.Println("Err", err.Error())
// 		}
// 		c.IndentedJSON(http.StatusOK, article)
// 	}
// }

func validateLogin(c *gin.Context) {
	var json usernamejson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pass string
	err := db.QueryRow("SELECT F_USER_PWD FROM app_mas_users where F_USER_NAME=?", json.Uname).Scan(&pass)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if pass != json.Pwd {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}

func handleRequests() {
	router := gin.Default()

	// router.GET("/articles", allArticles)
	// router.GET("/article/:id", getArticle)
	router.POST("/validatelogin/", validateLogin)

	router.Run("localhost:8080")
	// http.HandleFunc("/", homePage)
	// //http.HandleFunc("/articles", allArticles)
	// log.Fatal(http.ListenAndServe(":8081", nil))
}

var db *sql.DB

func main() {
	db = connectDB()
	defer db.Close()

	handleRequests()
}

func connectDB() *sql.DB {
	// Replace "username", "password", "dbname" with your database credentials
	connectionString := "root:Nalam123@tcp(localhost:3306)/goappdb"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

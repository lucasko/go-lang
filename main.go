package main

import (
    "net/http"
	"fmt"
    "github.com/gin-gonic/gin"
)

// Book represents data about a record Book.
type Book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}

// Books slice to seed record Book data.
var Books = []Book{
    {ID: "1", Title: "book - 1", Author: "test 1", Price: 9.99},
    {ID: "2", Title: "book - 2", Author: "test 2", Price: 19.99},
    {ID: "3", Title: "book - 3", Author: "test 3", Price: 29.99},
}

func main() {
    router := gin.Default()

    // SELECT
    router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)

	// INSERT
    router.POST("/books", postBooks)

    // DELETE
	router.DELETE("/books/:id", deleteBookByID)

    // UPDATE
	router.PUT("/books/:id", updateBookByID)

	// start backend server
    router.Run("localhost:8080")
}
func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, Books)
}

func postBooks(c *gin.Context) {
    var newBook Book

    // Call BindJSON to bind the received JSON to
    // newBook.
    if err := c.BindJSON(&newBook); err != nil {
        return
    }

    // Add the new Book to the slice.
    Books = append(Books, newBook)
    newBook.Title = "haha"
    c.IndentedJSON(http.StatusCreated, newBook)
}


func updateBookByID(c *gin.Context) {
 	id := c.Param("id")
    var newBook Book

	// receiving payload from request
    if err := c.BindJSON(&newBook); err != nil {
        return
    }

    foundBook := getBook(id)
    foundBook.Price = newBook.Price
    c.IndentedJSON(http.StatusCreated, foundBook)
}

func getBook(id string) Book {
 for _, a := range Books {
        if a.ID == id {
            return a
        }
    }
 var al Book
 return al
}

func getBookByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range Books {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func deleteBookByID(c *gin.Context) {
    id := c.Param("id")
  	message := fmt.Sprintf("deleted id=%s", id)
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": message })
}
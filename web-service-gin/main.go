package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// !album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// !albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// !the handler function to an endpoint path
func main() {
	router := gin.Default()          //* Initialize a Gin router using Default.
	router.GET("/albums", getAlbums) //* Use the GET function to associate the GET HTTP method and /albums path with a handler function.
	// Note! you’re passing the name of the getAlbums function. This is different from passing the result of the function, which you would do by passing getAlbums() (note the parenthesis).
	router.Run("localhost:8080") //* Use the Run function to attach the router to an http.Server and start the server.
}

// !getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	//* Serialize the struct into indented JSON and add it to the response
	//First argument: HTTP status code you want to send to the client
	c.IndentedJSON(http.StatusOK, albums) //* StatusOK == 200
}

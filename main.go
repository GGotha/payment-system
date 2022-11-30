package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{Id: "1", Title: "teste", Artist: "Mateus", Price: 56.00},
	{Id: "2", Title: "teste2", Artist: "João", Price: 56.00},
	{Id: "3", Title: "teste3", Artist: "Pires", Price: 56.00},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:3333")

	fmt.Println("Hello world")
}

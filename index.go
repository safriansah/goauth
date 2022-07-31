package main

import (
	"goauth/routers"
	"goauth/services"
	"goauth/utils"
	"log"
	"net/http"
)

func main() {
	// Declaring some variables
	// var name string
	// var alphabet_count int

	// // Calling Scanf() function for
	// // scanning and reading the input
	// // texts given in standard input
	// fmt.Scanf("%s", &name)
	// fmt.Scanf("%d", &alphabet_count)

	// // Printing the given texts
	// fmt.Printf("The word %s containing %d number of alphabets.",
	// 	name, alphabet_count)
	// fmt.Println("Hello, World!")

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })
	// r.GET("/books", controllers.GetBooks)
	// r.POST("/books", controllers.CreateBook)
	// r.Run(":4080")

	var db = utils.GetConnection()
	services.SetDB(db)
	var appRouter = routers.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}

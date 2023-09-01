package main

import "github.com/gin-gonic/gin"

var allUsers []User
var allMessages []Message

func main() {
	route := gin.Default()

	allUsers = []User{}
	allMessages = []Message{}

	route.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"users": allUsers})
	})

	route.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, value := range allUsers {
			if value.Id == id {
				payload := gin.H{
					"name": value.Name,
					"links": map[string]string{
						"aa": "asdads",
					},
				}
				c.JSON(200, payload)
			}
		}
	})

	route.POST("/users", func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")

		newUser := User{
			Id:   id,
			Name: name,
		}

		allUsers = append(allUsers, newUser)
		c.JSON(201, gin.H{})
	})

	route.GET("/users/:id/messages/:id", func(c *gin.Context) {

	})

	route.Run(":8088")
}

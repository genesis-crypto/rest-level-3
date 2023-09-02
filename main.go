package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var allUsers []User
var allMessages []Message

func main() {
	route := gin.Default()

	allUsers = []User{}
	allMessages = []Message{}

	route.GET("/users", func(c *gin.Context) {
		userRepresentations := []gin.H{}

		selfLink := Link{
			Href: "/users",
			Rel:  "self",
		}

		for _, value := range allUsers {
			userRepresentation := gin.H{
				"id":   value.Id,
				"name": value.Name,
			}

			userRepresentation["_links"] = map[string]Link{
				"self": Link{
					Href: fmt.Sprintf("/users/%s", value.Id),
					Rel:  "self",
				},
			}

			userRepresentations = append(userRepresentations, userRepresentation)
		}

		response := gin.H{
			"users": userRepresentations,
			"_links": map[string]Link{
				"self": selfLink,
			},
		}

		c.JSON(200, response)
	})

	route.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		var user User
		for _, value := range allUsers {
			if value.Id == id {
				user = value
				break
			}
		}

		if user.Id == "" {
			c.JSON(404, gin.H{"message": "User not found"})
			return
		}

		selfLink := Link{
			Href: fmt.Sprintf("/users/%s", user.Id),
			Rel:  "self",
		}

		userRepresentation := gin.H{
			"id":   user.Id,
			"name": user.Name,
			"_links": map[string]Link{
				"self": selfLink,
			},
		}

		c.JSON(200, userRepresentation)
	})

	route.POST("/users", func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")

		newUser := User{
			Id:   id,
			Name: name,
		}

		allUsers = append(allUsers, newUser)
		c.JSON(201, gin.H{"message": "User created"})
	})

	route.Run(":8088")
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

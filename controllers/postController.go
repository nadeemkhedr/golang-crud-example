package controllers

import (
	"example-crud/initializers"
	"example-crud/models"

	"github.com/gin-gonic/gin"
)


func PostsCreate(c *gin.Context) {
	// get data from request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// return the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}


func PostsShow(c *gin.Context) {
	// get id from the url
	id := c.Param("id")

	// get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// return the posts
	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsUpdate(c *gin.Context) {
	// get id from the url
	id := c.Param("id")

	// get the data from req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// send it
	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsDelete(c *gin.Context) {
	// get id from the url
	id := c.Param("id")

	// get the posts
	initializers.DB.Delete(&models.Post{}, id)

	// return the posts
	c.Status(200)
}

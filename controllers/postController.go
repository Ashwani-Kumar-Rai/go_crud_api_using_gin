package controllers

import (
	"go_crud_api_using_gin/initializers"
	"go_crud_api_using_gin/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get data off request body 
	var body struct{
		Body string
		Title string 
	}
	c.Bind(&body)

	// Create a post 
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it 
	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsIndex(c *gin.Context){
	//Get the posts
	var posts []models.Post 
	initializers.DB.Find(&posts)

	//Respond with them 
	c.JSON(200, gin.H{
		"post": posts,
	})
	
}

func PostsShow(c *gin.Context){

	// Get id off URL
	id := c.Param("id")

	//Get the posts
	var post models.Post 
	initializers.DB.Find(&post, id)

	//Respond with them 
	c.JSON(200, gin.H{
		"post": post,
	})
	
}

func PostsUpdate(c *gin.Context){

	// Get id off URL
	id := c.Param("id")

	//Get the data off request body 
	var body struct{
		Body string
		Title string 
	}
	c.Bind(&body)

	//Find the post were updating 
	var post models.Post 
	initializers.DB.Find(&post, id)

	// update it 
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// Respond with it 
	c.JSON(200, gin.H{
		"post": post,
	})
	
}


func PostsDelete(c *gin.Context){

	// Get id off URL
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{},id)
	
	// Respond with it 
	c.Status(200)
	
}
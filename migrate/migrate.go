package main

import (
	"example-crud/initializers"
	"example-crud/models"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDB()
}


func main() {
  initializers.DB.AutoMigrate(&models.Post{})
}

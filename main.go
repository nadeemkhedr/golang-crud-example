package main

import (
	"example-crud/initializers"

	"github.com/gin-gonic/gin"
)

var ninjaWeapons = map[string]string{
	"ninjaStar": "Beginner Ninja Star - Damage 1",
}

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongo",
		})
	})

	r.GET("/weapon", GetWeapon)

	r.Run()
}

func GetWeapon(c *gin.Context) {
	weaponType := c.Query("type")
	name := c.Query("name")
	c.JSON(200, gin.H{
		"weapon": ninjaWeapons[weaponType],
		"name":   name,
	})
}

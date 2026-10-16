package main

// import "github.com/gin-gonic/gin"

// func Ping(c *gin.Context) { // Context menggantikan w , r atau write , read
// number := c.Query("number")
// 	c.JSON(200, gin.H{ // c = context
// 		"message": "Pong " + number,
// 	})
// }

// func Halo(c *gin.Context) {
// 	nama := c.Param("nama")
// 	umur := c.Param("umur")
// 	c.JSON(200, gin.H{
// 		"message": "Hai " + nama, "umur " : umur,
// 	})
// }

import (
  "fmt"
  "golangframework/database"
  "golangframework/router"
  "log"
)

func main() {
  db, err := database.Connect()
  if err != nil {
    log.Fatal("Gagal terhubung ke database: ", err)
  }
  defer db.Close()
  fmt.Println("✅ Terhubung ke database PostgreSQL")

  r := router.SetupRouter(db)

  fmt.Println("🚀 Server Gin berjalan di http://localhost:8080")
  if err := r.Run(":8080"); err != nil {
    log.Fatal("Gagal menjalankan server: ", err)
  }
}

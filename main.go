package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	// "github.com/Bharadwaj-Srikumar/kistenmeister/tree/main/datenbank"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./kistenmeister/Kistenmeister.db"

/*type DB struct {
	Database *sql.DB
}*/

/*type Routes struct {
	Database datenbank.DB
}*/

func setupRouter() *gin.Engine {

	r := gin.Default()

	// Route zum Erstellen der Datenbank, wenn nicht bereits vorhanden
	r.GET("/createDB", func(c *gin.Context) {
		// Prüfen, ob die Datenbank bereits existiert
		if _, err := os.Stat(dbPath); err == nil { // wenn Datenbank schon vorhanden
			c.String(http.StatusOK, "Datenbank existiert bereits")
		} else if db, err := sql.Open("sqlite3", dbPath); err == nil { // wenn Datenbank nicht vorhanden
			defer db.Close()
			c.String(http.StatusCreated, "Datenbank erfolgreich erstellt")
		} else { // Fehler bei der Eingabe des Pfads
			c.String(http.StatusInternalServerError, fmt.Sprintf("Fehler bei der Datenbankerstellung: %v", err))
			return
		}
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Read box
	//TODO prüfen ob HEAD oder GET besser passt
	r.GET("/box/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		//TODO Get box data from SQlite

		//TODO Return box data
		c.JSON(http.StatusOK, gin.H{"box read": id})
	})

	// Create box
	r.POST("/box", func(c *gin.Context) {
		//TODO Create box data in SQlite

		//TODO Return box data
		c.JSON(http.StatusOK, gin.H{"box create": "newid"})
	})

	// Update box
	//TODO prüfen ob PATCH oder PUT besser passt
	r.PUT("/box/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		//TODO Update box data in SQlite

		//TODO Return box data
		c.JSON(http.StatusOK, gin.H{"box update": id})
	})

	// Delete box
	r.DELETE("/box/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		//TODO Delete box data in SQlite

		//TODO Return box data
		c.JSON(http.StatusOK, gin.H{"box delete": id})
	})

	// Get all boxes
	r.GET("/boxes", func(c *gin.Context) {
		//TODO Get all boxes data from SQlite

		//TODO Return all boxes data
		c.JSON(http.StatusOK, gin.H{"boxes": "all"})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Serve in 0.0.0.0:8080
	r.Run(":8080")
}

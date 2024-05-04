package main

import (
	"database/sql"
	//"fmt"
	"log"
	"net/http"

	//"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./kistenmeister.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type kiste struct {
	id               int    `json:"id"`
	Name             string `json:"Name"`
	Beschreibung     string `json:"Beschreibung"`
	Ersteller        string `json:"Ersteller"`
	Erstellungsdatum string `json:"Erstellungsdatum"`
	Änderer          string `json:"Änderer"`
	Änderungssdatum  string `json:"Änderungssdatum"`
	Verantwortlicher string `json:"Verantwortlicher"`
	//QRCode           string `json:"QRCode"`
	Ort string `json:"Ort"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ConnectDatabase()
	r := gin.Default()
	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/kisten", getKisten)
		v1.GET("/kiste/:id", getKiste)
		v1.POST("/kiste", createKiste)
		v1.PUT("/kiste/:id", aktKiste)
		v1.DELETE("/kiste/:id", löschenKiste)
	}
	// Listen and Serve in 0.0.0.0:8080
	r.Run(":8080")
}

// Read box
// TODO prüfen ob HEAD oder GET besser passt
func getKiste(c *gin.Context) {
	ID := c.Params.ByName("id")
	//TODO Get box data from SQlite

	//TODO Return box data
	c.JSON(http.StatusOK, gin.H{"Kiste lesen": ID})
}

// Create box
func createKiste(c *gin.Context) {
	//TODO Create box data in SQlite

	//TODO Return box data
	c.JSON(http.StatusOK, gin.H{"Kiste erstellen": "newid"})
}

// Update box
// TODO prüfen ob PATCH oder PUT besser passt
func aktKiste(c *gin.Context) {
	ID := c.Params.ByName("id")

	//TODO Update box data in SQlite

	//TODO Return box data
	c.JSON(http.StatusOK, gin.H{"Kiste aktualiseren": ID})
}

// Delete box
func löschenKiste(c *gin.Context) {
	ID := c.Params.ByName("id")

	//TODO Delete box data in SQlite

	//TODO Return box data
	c.JSON(http.StatusOK, gin.H{"Kiste löschen": ID})
}

func getKisten(c *gin.Context) {

	alleKisten, err := GetKisten()
	checkErr(err)

	if alleKisten == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": alleKisten})
	}

}

// Get all boxes
func GetKisten() ([]kiste, error) {

	rows, err := DB.Query("SELECT id, Name, Beschreibung, Ersteller, Erstellungsdatum, Änderer, Änderungssdatum, Verantwortlicher, Ort from Kiste;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleKisten := make([]kiste, 0)

	for rows.Next() {
		einzelneKiste := kiste{}
		err = rows.Scan(&einzelneKiste.id, &einzelneKiste.Name, &einzelneKiste.Beschreibung, &einzelneKiste.Ersteller, &einzelneKiste.Erstellungsdatum, &einzelneKiste.Änderer, &einzelneKiste.Änderungssdatum, &einzelneKiste.Verantwortlicher, &einzelneKiste.Ort)

		if err != nil {
			return nil, err
		}

		alleKisten = append(alleKisten, einzelneKiste)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleKisten, err
}

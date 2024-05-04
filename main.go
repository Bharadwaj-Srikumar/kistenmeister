package main

import (
	"database/sql"
	"log"
	"net/http"

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
	// json Tags müssen zu jedem Attribut hinzugefügt werden
	id               int
	Name             string
	Beschreibung     string
	Ersteller        string
	Erstellungsdatum string
	Änderer          string
	Änderungssdatum  string
	Verantwortlicher string
	//QRCode  string -- funktioniert nicht
	Ort string
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
		v1.POST("/kiste", erstelleKiste)
		v1.PUT("/kiste/:id", aktualisereKiste)
		v1.DELETE("/kiste/:id", löscheKiste)
	}
	// Listen and Serve in 0.0.0.0:8080
	r.Run(":8080")
}

// Lese alle Ksiten in der Datenbank
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

// Lese Kiste mit angegebenem id
// TODO prüfen ob HEAD oder GET besser passt
func getKiste(c *gin.Context) {
	ID := c.Params.ByName("id")
	//TODO Get box data from SQlite

	//TODO Return box data
	c.JSON(http.StatusOK, gin.H{"Kiste lesen": ID})
}

// Erstelle neue Kiste
func erstelleKiste(c *gin.Context) {
	//TODO Create box data in SQlite

	//TODO Return box data
	c.JSON(http.StatusOK, gin.H{"Kiste erstellen": "newid"})
}

// Aktualisere Kiste
// TODO prüfen ob PATCH oder PUT besser passt
func aktualisereKiste(c *gin.Context) {
	ID := c.Params.ByName("id")
	//TODO Update box data in SQlite

	//TODO Return box data
	c.JSON(http.StatusOK, gin.H{"Kiste aktualiseren": ID})
}

// Lösche Kiste
func löscheKiste(c *gin.Context) {
	ID := c.Params.ByName("id")

	//TODO Delete box data in SQlite

	//TODO Return box data
	c.JSON(http.StatusOK, gin.H{"Kiste löschen": ID})
}

// Funktion zum Abfragen aller Kiste
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

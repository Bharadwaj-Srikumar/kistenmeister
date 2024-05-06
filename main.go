package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var dbPath = "./test.db"

func ConnectDatabase() error {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		return err
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	} else {
		DB = db
		return nil
	}
}

type kiste struct {
	// json Tags müssen zu jedem Attribut hinzugefügt werden
	id               int
	Name             string
	Beschreibung     string
	Ersteller        string
	Erstellungsdatum time.Time
	Änderer          string
	Änderungssdatum  time.Time
	Verantwortlicher string
	// QRCode        blob -- konvertieren von BLOB in string klappt aktuell nicht
	Ort string
}

type kommentar struct {
	// json Tags müssen zu jedem Attribut hinzugefügt werden
	id               int
	Kommentar        string
	Ersteller        string
	Erstellungsdatum time.Time
	box_id           int
}

type bild struct {
	// json Tags müssen zu jedem Attribut hinzugefügt werden
	id int
	// bild          blob -- konvertieren von BLOB in string klappt aktuell nicht
	Ersteller        string
	Erstellungsdatum time.Time
	box_id           int
}

func setupRouter() *gin.Engine {

	r := gin.Default()

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
		alleKisten, err := GetBoxes()
		checkErr(err)

		if alleKisten == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
			return
		} else {
			//TODO Return all boxes data
			c.JSON(http.StatusOK, gin.H{"data": alleKisten})
		}
	})

	// Get all comments
	r.GET("/comments", func(c *gin.Context) {
		//TODO Get all boxes data from SQlite
		alleKommentare, err := GetComments()
		checkErr(err)

		if alleKommentare == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
			return
		} else {
			//TODO Return all boxes data
			c.JSON(http.StatusOK, gin.H{"data": alleKommentare})
		}
	})

	// Get all comments
	r.GET("/pictures", func(c *gin.Context) {
		//TODO Get all boxes data from SQlite
		alleBilder, err := GetPictures()
		checkErr(err)

		if alleBilder == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
			return
		} else {
			//TODO Return all boxes data
			c.JSON(http.StatusOK, gin.H{"data": alleBilder})
		}
	})

	return r
}

// Funktion zum Abfragen aller Kiste
func GetBoxes() ([]kiste, error) {

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

// Funktion zum Abfragen aller kommentare
func GetComments() ([]kommentar, error) {

	rows, err := DB.Query("SELECT id, Kommentar, Ersteller, Erstellungsdatum, box_id from Kommentare;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleKommentare := make([]kommentar, 0)

	for rows.Next() {
		einzelneKommentar := kommentar{}
		err = rows.Scan(&einzelneKommentar.id, &einzelneKommentar.Kommentar, &einzelneKommentar.Ersteller, &einzelneKommentar.Erstellungsdatum, &einzelneKommentar.box_id)

		if err != nil {
			return nil, err
		}

		alleKommentare = append(alleKommentare, einzelneKommentar)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleKommentare, err
}

func GetPictures() ([]bild, error) {

	rows, err := DB.Query("SELECT id, Ersteller, Erstellungsdatum, box_id from bilder;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleBilder := make([]bild, 0)

	for rows.Next() {
		einzelnesBild := bild{}
		err = rows.Scan(&einzelnesBild.id, &einzelnesBild.Ersteller, &einzelnesBild.Erstellungsdatum, &einzelnesBild.box_id)

		if err != nil {
			return nil, err
		}

		alleBilder = append(alleBilder, einzelnesBild)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleBilder, err
}

// Funktion zum Prüfen nach Error -> verwendet in den Router-Funktionen
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ConnectDatabase()
	r := setupRouter()
	r.Run(":8080")
}

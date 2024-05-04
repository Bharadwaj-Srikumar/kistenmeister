package db

import (
	"database/sql"
	"log"

	// "time"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Database *sql.DB
}

/*type kisten struct {
	id               int
	Name             string
	Beschreibung     string
	Ersteller        string
	Erstellungsdatum string
	Änderer          string
	Änderungsdatum   string
	Verantwortlicher string
	QRCode           string
	Ort              string
}*/

func (d *DB) InitDB() {
	db, err := sql.Open("sqlite3", "./kistenmeister/datenbank/db")
	if err != nil {
		log.Fatal(err)
	}
	d.Database = db

	initStmt := `
	CREATE TABLE IF NOT EXISTS Kiste (id INTEGER PRIMARY KEY, Name TEXT, Beschreibung TEXT, Ersteller TEXT, Erstellungsdatum DATETIME, Änderer TEXT, Änderungssdatum DATETIME, Verantwortlicher TEXT, QRCode BLOB, Ort TEXT);
	CREATE TABLE IF NOT EXISTS Kommentare(id INTEGER primary key not null, Kommentar text, Ersteller text,  Erstellungsdatum datetime, FOREIGN KEY(box_id) REFERENCES kisten(id));
	CREATE TABLE bilder(id INTEGER primary key not null, bild BLOB,Ersteller text, Erstellungsdatum datetime,FOREIGN KEY(box_id) REFERENCES kisten(id)),`

	_, err = d.Database.Exec(initStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, initStmt)
	}
}

/*
func (d *DB) AddPhoto(name string, timeTaken int) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into PHOTOS(FILENAME, LAST_VIEWED, TIME_TAKEN) values(?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(name, 0, timeTaken)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

func (d *DB) UpdateLastViewed(name string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE PHOTOS SET LAST_VIEWED=? WHERE FILENAME=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(time.Now().Unix(), name)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

func (d *DB) GetNextSlideshowPhoto() (PhotoDetails, error) {
	rows, err := d.Database.Query("select FILENAME, LAST_VIEWED, TIME_TAKEN from PHOTOS ORDER BY LAST_VIEWED LIMIT 1")
	if err != nil {
		return PhotoDetails{}, err
	}
	defer rows.Close()

	for rows.Next() {
		details := PhotoDetails{}
		err := rows.Scan(&details.Name, &details.LastViewed, &details.TimeTaken)
		if err != nil {
			return PhotoDetails{}, nil
		}
		return details, nil
	}

	err = rows.Err()
	return PhotoDetails{}, err
}

func (d *DB) GetPhotoList(count int, offset int) ([]PhotoDetails, error) {

	return []PhotoDetails{}, nil
}
*/

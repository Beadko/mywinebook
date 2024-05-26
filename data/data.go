package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS mywinebook (
        "idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "wineType" TEXT,
        "country" TEXT
      );`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Mywinebook table created")
}

func InsertNote(name string, wineType string) {
	insertNoteSQL := `INSERT INTO mywinebook(name, wineType) VALUES (?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(name, wineType)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Note added successfully")
}

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM mywinebook ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var name string
		var wineType string
		var country string
		row.Scan(&idNote, &name, &wineType, &country)
		log.Println("[", country, "] ", name, "—", wineType)
	}
}

func AmendNote(newName string, oldName string) {
	updateNoteSQL := `UPDATE mywinebook SET name = ? WHERE name = ?`
	statement, err := db.Prepare(updateNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(newName, oldName)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Note updated successfully")
}

func DeleteNote(name string) {
	deleteNoteSQL := `DELETE FROM mywinebook WHERE name =?`
	statement, err := db.Prepare(deleteNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(name)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Note deleted successfully")
}

package data

import (
	"database/sql"
	"log"

	"github.com/Beadko/mywinebook/internal/wine"
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
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "wineType" TEXT
	);`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Mywinebook table created")
}

func AddWine(name string, wineType string) error {
	insertNoteSQL := `INSERT INTO mywinebook(name, wineType) VALUES (?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(name, wineType)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Wine added successfully")
	defer db.Close()
	return nil
}

func GetWines() ([]wine.Wine, error) {
	row, err := db.Query("SELECT id, name, wineType FROM mywinebook ORDER BY id")
	if err != nil {
		return nil, err
	}

	defer row.Close()
	wines := []wine.Wine{}

	for row.Next() {
		w := wine.Wine{}
		row.Scan(&w.ID, &w.Name, &w.Type)
	}
	return wines, nil
}

func UpdateWine(wine.Wine) (wine.Wine, error) {
	updateNoteSQL := `UPDATE mywinebook SET name = ?, wineType = ? WHERE id = ?`
	statement, err := db.Prepare(updateNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	w := wine.Wine{}
	_, err = statement.Exec(&w.ID, &w.Name, &w.Type)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Wine updated successfully")
	return w, nil
}

func DeleteWine(name string) error {
	deleteNoteSQL := `DELETE FROM mywinebook WHERE name =?`
	statement, err := db.Prepare(deleteNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(name)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Wine deleted successfully")
	return nil
}

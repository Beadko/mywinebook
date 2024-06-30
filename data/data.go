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

func InitDB() {
	db.Exec(`CREATE TABLE IF NOT EXISTS wine_types (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE
			);
			INSERT INTO wine_types(name) VALUES ("Red"), ("White"), ("Rose"), ("Sparkling");`)
	db.Exec(`CREATE TABLE IF NOT EXISTS wines (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				name TEXT,
				wineType INTEGER REFERENCES wine_types
	);`)
	log.Println("Database created")
}

func GetWineTypes() ([]wine.WineType, error) {
	rows, err := db.Query(`SELECT * FROM wine_types ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	wineTypes := []wine.WineType{}
	for rows.Next() {
		wt := wine.WineType{}
		if err := rows.Scan(&wt.ID, &wt.Name); err != nil {
			return nil, err
		}
		wineTypes = append(wineTypes, wt)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(wineTypes)
	return wineTypes, nil
}

func AddWine(name string, wineType int) error {
	insertNoteSQL := `INSERT INTO wines(name, wineType) VALUES (?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(name, wineType)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Wine added successfully")
	return nil
}

func GetWines() ([]wine.Wine, error) {
	rows, err := db.Query("SELECT id, name, wineType FROM wines ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	wines := []wine.Wine{}

	for rows.Next() {
		w := wine.Wine{}
		if err := rows.Scan(&w.ID, &w.Name, &w.TypeID); err != nil {
			return nil, err
		}
		wines = append(wines, w)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(wines)
	return wines, nil
}

func GetWine(id string) (wine.Wine, error) {
	log.Printf("Getting wine %s", id)
	row := db.QueryRow("SELECT * FROM wines WHERE id = ?", id)
	w := wine.Wine{}
	if err := row.Scan(&w.ID, &w.Name, &w.Type); err != nil {
		return wine.Wine{}, err
	}
	if err := row.Err(); err != nil {
		return wine.Wine{}, err
	}
	return w, nil
}

func UpdateWine(id string, name string, wineType int) error {
	_, err := db.Exec(`UPDATE wines SET name = ?, wineType = ? WHERE id = ?`, name, wineType, id)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Wine updated successfully")
	return nil
}

func DeleteWine(id string) error {
	deleteNoteSQL := `DELETE FROM wines WHERE id = ?`
	statement, err := db.Prepare(deleteNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(id)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Wine deleted successfully")
	return nil
}

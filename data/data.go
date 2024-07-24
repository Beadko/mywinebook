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
	db, err = sql.Open("sqlite3", "./sqlite-database.db?_foreign_keys=on")
	if err != nil {
		return err
	}
	return db.Ping()
}

func InitDB() {

	db.Exec(`
		CREATE TABLE IF NOT EXISTS wine_types (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE
			);
			INSERT INTO wine_types(name) VALUES ("Red"), ("White"), ("Rose"), ("Sparkling");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS countries (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE
		);
		INSERT INTO countries(name) VALUES ("France"), ("Italy"), ("Australia");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS wines (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			wineType INT unsigned NOT NULL,
			country INT unsigned NOT NULL,
			FOREIGN KEY (wineType) REFERENCES wine_types(id),
			FOREIGN KEY (country) REFERENCES countries(id)
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

func GetCountries() ([]wine.Country, error) {
	rows, err := db.Query(`SELECT * FROM countries ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	countries := []wine.Country{}
	for rows.Next() {
		c := wine.Country{}
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		countries = append(countries, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(countries)
	return countries, nil
}

func AddWine(name string, wineType int, country int) error {
	insertNoteSQL := `INSERT INTO wines(name, wineType, country) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(name, wineType, country)
	if err == nil {
		log.Println("Wine added successfully")
		return nil
	}
	return err
}

func GetWines() ([]wine.Wine, error) {
	rows, err := db.Query("SELECT * FROM wines ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	wines := []wine.Wine{}

	for rows.Next() {
		w := wine.Wine{}
		if err := rows.Scan(&w.ID, &w.Name, &w.TypeID, &w.CountryID); err != nil {
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
	if err := row.Scan(&w.ID, &w.Name, &w.TypeID, &w.CountryID); err != nil {
		return wine.Wine{}, err
	}
	if err := row.Err(); err != nil {
		return wine.Wine{}, err
	}
	return w, nil
}

func UpdateWine(id string, name string, wineType int, country int) error {
	_, err := db.Exec(`UPDATE wines SET name = ?, wineType = ?, country = ? WHERE id = ?`, name, wineType, country, id)
	if err == nil {
		log.Println("Wine updated successfully")
		return nil
	}
	return err
}

func DeleteWine(id string) error {
	deleteNoteSQL := `DELETE FROM wines WHERE id = ?`
	statement, err := db.Prepare(deleteNoteSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	if err == nil {
		log.Println("Wine deleted successfully")
		return nil
	}
	return err
}

func AddCountry(name string) error {
	_, err := db.Exec(`INSERT INTO countries(name) VALUES (?)`, name)
	if err == nil {
		log.Println("Country added successfully")
		return nil
	}
	return err
}

func AddWineType(name string) error {
	_, err := db.Exec(`INSERT INTO wine_types(name) VALUES (?)`, name)
	if err == nil {
		log.Println("Wine type added successfully")
		return nil
	}
	return err
}

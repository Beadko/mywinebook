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
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE
			);
			INSERT INTO wine_types(name) VALUES ("Red"), ("White"), ("Rose"), ("Sparkling"), ("Orange");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS countries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE
		);
		INSERT INTO countries(name) VALUES ("France"), ("Italy"), ("Australia"), ("Spain"), ("New Zealand"), ("Chile"), ("Germany"), ("Malta"), ("USA"), ("Argentina"), ("South Africa"), ("Portugal"), ("Hungary"), ("Georgia");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS aromas (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE
		);
		INSERT INTO aromas(name) VALUES ("Fruity"), ("Vegetal"), ("Floral"), ("Earthy"), ("Woody"), ("Spicy"), ("Mineral");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS intensities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO instensities(name) VALUES ("Weak"), ("Medium"), ("Pronounced");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS flavours (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO flavours(name) VALUES ("Fruity"), ("Vegetal"), ("Floral"), ("Earthy"), ("Woody"), ("Spicy"), ("Mineral"), ("Herbal");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS sweetness (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO sweetness(name) VALUES ("Bone dry"), ("Dry"), ("Off dry"), ("Medium sweet"), ("Sweet"), ("Very sweet");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS acidities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO acidities(name) VALUES ("Tart"), ("Crisp"), ("Fresh"), ("Smooth"), ("Not acidic");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS tannins (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO tannins(name) VALUES ("Very light"), ("Light"), ("Medium"), ("Full Bodied"), ("Heavy");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS bodies (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO bodies(name) VALUES ("Soft"), ("Round"), ("Dry"), ("Hard");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS clarities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO clarities(name) VALUES ("Clear"), ("Slightly hazy"), ("Hazy");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS finishes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO finishes(name) VALUES ("Short"), ("Medium"), ("Long"), ("Very long");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS colours (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO colours(name) VALUES ("Purple"), ("Ruby"), ("Garnet"), ("Tawny"), ("Straw"), ("Yellow"), ("Golden"), ("Blush"), ("Salmon"), ("Pink");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS colour_depths (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO colour_depths(name) VALUES ("Pale"), ("Medium"), ("Deep");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS balances (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO balances(name) VALUES ("Good"), ("Fair"), ("Unbalanced");`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS wines (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100),
			wineType INT unsigned,
			country INT unsigned,
			score INT unsigned,
			year INT unsigned,
			producer VARCHAR(100),
			alcohol DECIMAL(4, 2),
			aroma INT unsigned,
			intensity INT unsigned,
			palate INT unsigned,
			flavour INT unsigned,
			sweetness INT unsigned,
			acidity INT unsigned,
			tannin INT unsigned,
			body INT unsigned,
			clarity INT unsigned,
			finish INT unsigned,
			colour INT unsigned,
			colour_depth INT unsigned,
			balance INT unsigned,
			FOREIGN KEY (wineType) REFERENCES wine_types(id),
			FOREIGN KEY (country) REFERENCES countries(id),
			FOREIGN KEY (aroma) REFERENCES aromas(id),
			FOREIGN KEY (instensity) REFERENCES intensities(id),
			FOREIGN KEY (flavour) REFERENCES flavours(id),
			FOREIGN KEY (sweetness) REFERENCES sweetness(id),
			FOREIGN KEY (acidity) REFERENCES acidities(id),
			FOREIGN KEY (tannin) REFERENCES tannins(id),
			FOREIGN KEY (body) REFERENCES bodies(id),
			FOREIGN KEY (clarity) REFERENCES clarities(id),
			FOREIGN KEY (finish) REFERENCES finishes(id),
			FOREIGN KEY (colour) REFERENCES colours(id)
			FOREIGN KEY (colour_depth) REFERENCES colour_depths(id)
			FOREIGN KEY (balance) REFERENCES balances(id)
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

func AddWine(wine wine.Wine) error {
	insertNoteSQL := `INSERT INTO wines(name, wineType, country, score) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(wine.Name, wine.TypeID, wine.CountryID, wine.Score)
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
		if err := rows.Scan(&w.ID, &w.Name, &w.TypeID, &w.CountryID, &w.Score); err != nil {
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
	if err := row.Scan(&w.ID, &w.Name, &w.TypeID, &w.CountryID, &w.Score); err != nil {
		return wine.Wine{}, err
	}
	if err := row.Err(); err != nil {
		return wine.Wine{}, err
	}
	return w, nil
}

func UpdateWine(wine wine.Wine, id string) error {
	_, err := db.Exec(`UPDATE wines SET name = ?, wineType = ?, country = ?, score = ? WHERE id = ?`, wine.Name, wine.TypeID, wine.CountryID, wine.Score, id)
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

func UpdateCountry(id string, name string) error {
	_, err := db.Exec(`UPDATE countries SET name = ? WHERE id = ?`, name, id)
	if err == nil {
		log.Println("Country updated successfully")
		return nil
	}
	return err
}

func UpdateWineType(id string, name string) error {
	_, err := db.Exec(`UPDATE wine_types SET name = ? WHERE id = ?`, name, id)
	if err == nil {
		log.Println("Wine type updated successfully")
		return nil
	}
	return err
}

func DeleteCountry(id string) error {
	_, err := db.Exec(`DELETE FROM countries WHERE id = ?`, id)
	if err == nil {
		log.Println("Country deleted successfully")
		return nil
	}
	return err
}

func DeleteWineType(id string) error {
	_, err := db.Exec(`DELETE FROM wine_types WHERE id = ?`, id)
	if err == nil {
		log.Println("Wine type deleted successfully")
		return nil
	}
	return err
}

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

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS wine_types (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE
			);
			INSERT INTO wine_types(name) VALUES ("Red"), ("White"), ("Rose"), ("Sparkling"), ("Dessert"), ("Fortified"), ("Orange");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS countries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE
		);
		INSERT INTO countries(name) VALUES ("France"), ("Italy"), ("Australia"), ("Spain"), ("New Zealand"), ("Chile"), ("Germany"), ("Malta"), ("USA"), ("Argentina"), ("South Africa"), ("Portugal"), ("Hungary"), ("Georgia");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS colours (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO colours(name) VALUES ("Purple"), ("Ruby"), ("Garnet"), ("Tawny"), ("Straw"), ("Yellow"), ("Golden"), ("Blush"), ("Salmon"), ("Pink");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS colour_depths (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO colour_depths(name) VALUES ("Pale"), ("Medium"), ("Deep");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS aromas (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE
		);
		INSERT INTO aromas(name) VALUES ("Fruity"), ("Vegetal"), ("Floral"), ("Earthy"), ("Woody"), ("Spicy"), ("Mineral"),("Herbal"), ("Smoky");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS intensities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO intensities(name) VALUES ("Weak"), ("Medium"), ("Pronounced");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS flavours (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO flavours(name) VALUES ("Fruity"), ("Vegetal"), ("Floral"), ("Earthy"), ("Woody"), ("Spicy"), ("Mineral"), ("Herbal"), ("Smoky");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sweetnesses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO sweetnesses(name) VALUES ("Bone dry"), ("Dry"), ("Off dry"), ("Medium sweet"), ("Sweet"), ("Very sweet");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS acidities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO acidities(name) VALUES ("Tart"), ("Crisp"), ("Fresh"), ("Smooth"), ("Not acidic");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tannins (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO tannins(name) VALUES ("Soft"), ("Round"), ("Dry"), ("Hard");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS bodies (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO bodies(name) VALUES ("Very light"), ("Light"), ("Medium"), ("Full Bodied"), ("Heavy");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS clarities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO clarities(name) VALUES ("Clear"), ("Slightly hazy"), ("Hazy");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS finishes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE		
		);
		INSERT INTO finishes(name) VALUES ("Long"), ("Medium"), ("Short");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS balances (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) UNIQUE		
		);
		INSERT INTO balances(name) VALUES ("Good"), ("Fair"), ("Unbalanced");`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
	    PRAGMA foreign_keys = ON;
		CREATE TABLE IF NOT EXISTS wines (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			wine_type INTEGER REFERENCES wine_types(id),
			country INTEGER REFERENCES countries(id),
			score INTEGER,
			producer TEXT, 
			alcohol REAL,
			year INTEGER,
			colour INTEGER REFERENCES colours(id),
			colour_depth INTEGER REFERENCES colour_depths(id),
			clarity INTEGER REFERENCES clarities(id),
			aroma INTEGER REFERENCES aromas(id),
			intensity INTEGER REFERENCES intensities(id),
			flavour INTEGER REFERENCES flavours(id),
			sweetness INTEGER REFERENCES sweetnesses(id),
			acidity INTEGER REFERENCES acidities(id),
			tannin INTEGER REFERENCES tannins(id),
			body INTEGER REFERENCES bodies(id),
			finish INTEGER REFERENCES finishes(id),
			balance INTEGER REFERENCES balances(id),
			notes TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}
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

func AddWine(w wine.Wine) (int, error) {
	insertNoteSQL := `INSERT INTO wines(name, wine_type, country, score, producer, alcohol, year, colour, colour_depth, clarity, aroma, intensity, flavour, sweetness, acidity, tannin, body, finish, balance, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		return 0, err
	}
	result, err := statement.Exec(w.Name, w.TypeID, w.CountryID, w.Score, w.Producer, w.Alcohol, w.Year, w.ColourID, w.DepthID, w.ClarityID, w.AromaID, w.IntensityID, w.FlavourID, w.SweetnessID, w.AcidityID, w.TanninID, w.BodyID, w.FinishID, w.BalanceID, w.Notes)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	id := int(lastInsertId)
	log.Println("Wined added successfully:", id, w)
	return id, nil
}

func GetWines() ([]wine.Wine, error) {
	rows, err := db.Query("SELECT id, name, wine_type, country, score, producer, alcohol, year, colour, colour_depth, clarity, aroma, intensity, flavour, sweetness, acidity, tannin, body, finish, balance, notes FROM wines ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	wines := []wine.Wine{}

	for rows.Next() {
		w := wine.Wine{}
		if err := rows.Scan(&w.ID, &w.Name, &w.TypeID, &w.CountryID, &w.Score, &w.Producer, &w.Alcohol, &w.Year, &w.ColourID, &w.DepthID, &w.ClarityID, &w.AromaID, &w.IntensityID, &w.FlavourID, &w.SweetnessID, &w.AcidityID, &w.TanninID, &w.BodyID, &w.FinishID, &w.BalanceID, &w.Notes); err != nil {
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
	row := db.QueryRow("SELECT id, name, wine_type, country, score, producer, alcohol, year, colour, colour_depth, clarity, aroma, intensity, flavour, sweetness, acidity, tannin, body, finish, balance, notes FROM wines WHERE id = ?", id)
	w := wine.Wine{}
	if err := row.Scan(&w.ID, &w.Name, &w.TypeID, &w.CountryID, &w.Score, &w.Producer, &w.Alcohol, &w.Year, &w.ColourID, &w.DepthID, &w.ClarityID, &w.AromaID, &w.IntensityID, &w.FlavourID, &w.SweetnessID, &w.AcidityID, &w.TanninID, &w.BodyID, &w.FinishID, &w.BalanceID, &w.Notes); err != nil {
		return wine.Wine{}, err
	}
	if err := row.Err(); err != nil {
		return wine.Wine{}, err
	}
	log.Println(w)
	return w, nil
}

func UpdateWine(w wine.Wine, id string) error {
	_, err := db.Exec(`UPDATE wines SET name = ?, wine_type = ?, country = ?, score = ?, producer = ?, alcohol = ?, year = ?, colour = ?, colour_depth = ?, clarity = ?, aroma = ?, intensity = ?, flavour = ?, sweetness = ?, acidity = ?, tannin = ?, body = ?, finish = ?, balance = ?, notes = ? WHERE id = ?`, w.Name, w.TypeID, w.CountryID, w.Score, w.Producer, w.Alcohol, w.Year, w.ColourID, w.DepthID, w.ClarityID, w.AromaID, w.IntensityID, w.FlavourID, w.SweetnessID, w.AcidityID, w.TanninID, w.BodyID, w.FinishID, w.BalanceID, w.Notes, id)
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

func AddCountry(name string) (int, error) {
	result, err := db.Exec(`INSERT INTO countries(name) VALUES (?)`, name)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	id := int(lastInsertId)
	log.Println("Country added successfully:", id, name)
	return id, nil
}

func AddWineType(name string) (int, error) {
	result, err := db.Exec(`INSERT INTO wine_types(name) VALUES (?)`, name)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	id := int(lastInsertId)
	log.Println("Wine type added successfully:", id, name)
	return id, nil
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

func GetDepths() ([]wine.Depth, error) {
	rows, err := db.Query(`SELECT * FROM colour_depths ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	depths := []wine.Depth{}
	for rows.Next() {
		d := wine.Depth{}
		if err := rows.Scan(&d.ID, &d.Name); err != nil {
			return nil, err
		}
		depths = append(depths, d)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(depths)
	return depths, nil
}

func GetClarities() ([]wine.Clarity, error) {
	rows, err := db.Query(`SELECT * FROM clarities ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clarities := []wine.Clarity{}
	for rows.Next() {
		c := wine.Clarity{}
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		clarities = append(clarities, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(clarities)
	return clarities, nil
}

func GetAromas() ([]wine.Aroma, error) {
	rows, err := db.Query(`SELECT * FROM aromas ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	aromas := []wine.Aroma{}
	for rows.Next() {
		a := wine.Aroma{}
		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			return nil, err
		}
		aromas = append(aromas, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(aromas)
	return aromas, nil
}

func AddAroma(name string) error {
	_, err := db.Exec(`INSERT INTO aromas(name) VALUES (?)`, name)
	if err == nil {
		log.Println("Aroma added successfully")
		return nil
	}
	return err
}

func UpdateAroma(id string, name string) error {
	_, err := db.Exec(`UPDATE aromas SET name = ? WHERE id = ?`, name, id)
	if err == nil {
		log.Println("Aroma updated successfully")
		return nil
	}
	return err
}

func DeleteAroma(id string) error {
	_, err := db.Exec(`DELETE FROM aromas WHERE id = ?`, id)
	if err == nil {
		log.Println("Aroma deleted successfully")
		return nil
	}
	return err
}

func GetFlavours() ([]wine.Flavour, error) {
	rows, err := db.Query(`SELECT * FROM flavours ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flavour := []wine.Flavour{}
	for rows.Next() {
		f := wine.Flavour{}
		if err := rows.Scan(&f.ID, &f.Name); err != nil {
			return nil, err
		}
		flavour = append(flavour, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(flavour)
	return flavour, nil
}

func AddFlavour(name string) error {
	_, err := db.Exec(`INSERT INTO flavours(name) VALUES (?)`, name)
	if err == nil {
		log.Println("Flavour added successfully")
		return nil
	}
	return err
}

func UpdateFlavour(id string, name string) error {
	_, err := db.Exec(`UPDATE flavours SET name = ? WHERE id = ?`, name, id)
	if err == nil {
		log.Println("Flavour updated successfully")
		return nil
	}
	return err
}

func DeleteFlavour(id string) error {
	_, err := db.Exec(`DELETE FROM flavours WHERE id = ?`, id)
	if err == nil {
		log.Println("Flavour deleted successfully")
		return nil
	}
	return err
}
func GetBodies() ([]wine.Body, error) {
	rows, err := db.Query(`SELECT * FROM bodies ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bodies := []wine.Body{}
	for rows.Next() {
		b := wine.Body{}
		if err := rows.Scan(&b.ID, &b.Name); err != nil {
			return nil, err
		}
		bodies = append(bodies, b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(bodies)
	return bodies, nil
}
func GetAcidities() ([]wine.Acidity, error) {
	rows, err := db.Query(`SELECT * FROM acidities ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	acidities := []wine.Acidity{}
	for rows.Next() {
		a := wine.Acidity{}
		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			return nil, err
		}
		acidities = append(acidities, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(acidities)
	return acidities, nil
}
func GetSweetness() ([]wine.Sweetness, error) {
	rows, err := db.Query(`SELECT * FROM sweetnesses ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sweetness := []wine.Sweetness{}
	for rows.Next() {
		s := wine.Sweetness{}
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		sweetness = append(sweetness, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(sweetness)
	return sweetness, nil
}

func GetTannins() ([]wine.Tannin, error) {
	rows, err := db.Query(`SELECT * FROM tannins ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tannins := []wine.Tannin{}
	for rows.Next() {
		t := wine.Tannin{}
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		tannins = append(tannins, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(tannins)
	return tannins, nil
}

func GetFinishes() ([]wine.Finish, error) {
	rows, err := db.Query(`SELECT * FROM finishes ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	finishes := []wine.Finish{}
	for rows.Next() {
		f := wine.Finish{}
		if err := rows.Scan(&f.ID, &f.Name); err != nil {
			return nil, err
		}
		finishes = append(finishes, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(finishes)
	return finishes, nil
}

func GetBalances() ([]wine.Balance, error) {
	rows, err := db.Query(`SELECT * FROM balances ORDER by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	balances := []wine.Balance{}
	for rows.Next() {
		b := wine.Balance{}
		if err := rows.Scan(&b.ID, &b.Name); err != nil {
			return nil, err
		}
		balances = append(balances, b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(balances)
	return balances, nil
}

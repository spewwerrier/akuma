package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/luitel777/akuma/internal/interface/akuma"
	_ "github.com/mattn/go-sqlite3"
)

type RETRIVE_TYPE uint8

const (
	HASH RETRIVE_TYPE = 0
	NAME RETRIVE_TYPE = 1
)

func ClearHashmap() {
	db, err := sql.Open("sqlite3", "database/akuma.db")
	if err != nil {
		log.Println(err)
	}
	db.Exec("DROP TABLE hashmaps")
	db.Close()
}

func createTableHashmaps(db *sql.DB) {
	_, err := db.Exec(
		"CREATE TABLE hashmaps (id integer PRIMARY KEY AUTOINCREMENT, name varchar(256) not null unique, hashmaps varchar(256) not null)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`create table chapters (
		  id integer PRIMARY KEY AUTOINCREMENT,
		  volume integer,
		  h_id FOREIGNKEY REFERENCES hashmaps(id)
		  on delete cascade
		  on update cascade
		  on delete restrict
		  )`)
	if err != nil {
		log.Fatal(err)
	}
}

func insertIntoTable(db *sql.DB, name string) {
	name = strings.ReplaceAll(name, "'", "''")
	sequel := fmt.Sprintf("insert into hashmaps (name, hashmaps) values ('%s', '%s')", name, akuma.CreateHashmaps(name))

	_, err := db.Exec(sequel)
	if err != nil {
		log.Fatal(err)
	}
}

func RetriveManga(name string, retrive RETRIVE_TYPE) string {
	db, err := sql.Open("sqlite3", "database/akuma.db")
	if err != nil {
		log.Fatal(err)
	}
	name = strings.ReplaceAll(name, "'", "''")

	var sequel string
	if retrive == HASH {
		sequel = fmt.Sprintf("select hashmaps from hashmaps where name = '%s'", name)
	} else if retrive == NAME {
		sequel = fmt.Sprintf("select name from hashmaps where hashmaps = '%s'", name)
	}

	resp, err := db.Query(sequel)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Close()

	var retrived string
	for resp.Next() {
		resp.Scan(&retrived)
	}
	db.Close()
	return retrived
}

func RetriveID(hash string) int {
	db, err := sql.Open("sqlite3", "database/akuma.db")
	if err != nil {
		log.Fatal(err)
	}
	hash = strings.ReplaceAll(hash, "'", "''")

	var sequel string
	sequel = fmt.Sprintf("select id from hashmaps where hashmaps = '%s'", hash)

	resp, err := db.Query(sequel)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Close()

	var retrived int
	for resp.Next() {
		resp.Scan(&retrived)
	}
	db.Close()
	return retrived
}

func GenerateUniqueIdenfitiers() {
	err := os.Mkdir("database", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	_, err = os.Create("database/akuma.db")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "database/akuma.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableHashmaps(db)

	for _, file := range akuma.GetMangaList() {
		insertIntoTable(db, file.Name())
	}
	db.Close()
}

package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

/*
manga archives might be structured like this

goodnight punpun v01 (2016) (omnibus edition) (digital) (1r0n).cbz
Goodnight Punpun v02 (2016) (Omnibus Edition) (Digital) (1r0n).cbz
Goodnight Punpun v04 (2016) (Omnibus Edition) (Digital) (1r0n).cbz
Goodnight Punpun v06 (2017) (Omnibus Edition) (Digital) (1r0n).cbz


so instead of localhost:3333/manga/a39b12b3571e909b14d9620b66343b84/Goodnight Punpun v06 (2017) (Omnibus Edition) (Digital) (1r0n).cbz
this will become localhost:3333/manga/a39b12b3571e909b14d9620b66343b84/06

to do this we create table for each hash
hash is already implemented for manga archives, not the contents inside it

*/

func returnHashes(db *sql.DB) []string {

	query := `select hashmaps from hashmaps`
	resp, err := db.Query(query)
	if err != nil {
		log.Print("possibly no manga added")
	}

	defer resp.Close()

	var hash string
	var hashes []string

	for resp.Next() {
		resp.Scan(&hash)
		hashes = append(hashes, hash)
	}
	return hashes
}

func InsertChapters(volume int, id int) {
	db, err := sql.Open("sqlite3", "database/akuma.db")
	if err != nil {
		log.Panic(err)
	}
	query := fmt.Sprintf("insert into chapters (volume, h_id) VALUES (%d, %d)", volume, id)
	db.Exec(query)
	db.Close()
}

func CheckIfMangaExists(id int) int {

	db, err := sql.Open("sqlite3", "database/akuma.db")
	if err != nil {
		log.Fatal(err)
	}
	sequel := fmt.Sprintf("select h_id from chapters where h_id=%d", id)

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

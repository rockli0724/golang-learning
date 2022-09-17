package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	db *sql.DB
	id int64
)

func parseCmdLineFlags() {
	flag.Int64Var(&id, "id", 0, "query id")
	flag.Parse()
}

//  go run mysql_access.go -id 2
func main() {
	parseCmdLineFlags()
	remainArgs := flag.Args()
	for _, arg := range remainArgs {
		fmt.Printf("remaining arg: %s \n", arg)
	}

	config := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		AllowNativePasswords: true,
		DBName:               "recordings",
	}
	var err error
	if db, err = sql.Open("mysql", config.FormatDSN()); err != nil {
		panic("init db error =" + err.Error())
	}
	if err = db.Ping(); err != nil {
		panic("ping db error = " + err.Error())
	}
	fmt.Println("init db success")

	albumList, err := queryAlbumList("John Coltrane") //John Coltrane
	if err != nil {
		panic("queryAlbumList err " + err.Error())
	}
	fmt.Printf("album find %v \n", albumList)

	album, err := queryAlbumById(id)
	if err != nil {
		panic("queryAlbumList err " + err.Error())
	}
	fmt.Printf("queryAlbumById %v \n", album)

	insertId, err := addAlbum(Album{
		Title:  "rockli",
		Artist: "this is a artible",
		Price:  234.45,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", insertId)

}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

//queryAlbumList queryAlbumList
func queryAlbumList(name string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var album Album
		if err = rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, album)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

//queryAlbumById queryAlbumById
func queryAlbumById(id int64) (Album, error) {
	var album Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumsById %d: no such album", id)
		}
		return album, fmt.Errorf("queryAlbumById %d: %v", id, err)
	}
	return album, nil
}

//queryAlbumById queryAlbumById
func queryPreAlbumById(id int64) (Album, error) {

	prepare, err := db.Prepare("SELECT * FROM album WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	var album Album
	if err = prepare.QueryRow(id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumsById %d: no such album", id)
		}
		return album, fmt.Errorf("queryAlbumById %d: %v", id, err)
	}
	return album, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

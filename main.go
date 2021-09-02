package main

import (
    "database/sql"
    "encoding/json"
    "time"
    "fmt"
    "regexp"
    "io/ioutil"
    "log"
    // "os"
    _ "github.com/go-sql-driver/mysql"
)

type Album struct {
    PhotoCount  string   `json:"photo_count"`
    ID          int      `json:"id"`
    URL         string   `json:"url"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    ViewCount   string   `json:"view_count"`
    Created     string   `json:"created"`
    LastUpdated string   `json:"last_updated"`
    CoverPhoto  string   `json:"cover_photo"`
    Photos      []string `json:"photos"`
}

type Albums struct {
    Albums      []Album  `json:"albums"`
}

func fetchFiles() []string {
    files, err := ioutil.ReadDir("./data/account")
    if err != nil {
        log.Fatal(err)
    }

    var list []string

    for _, f := range files {
        list = append(list, f.Name())
    }

    return list
}

func findPhotos() []string {
    var files = fetchFiles()
    var photos []string
    var re = regexp.MustCompile(`(?m)photo_[0-9]+.json`)

    for _, f := range files {
        if (re.MatchString(f)) {
            photos = append(photos, f)
        }
    }

    return photos
}

func resetTable(db *sql.DB) {
    truncate, err := db.Query("TRUNCATE TABLE photos")
    if err != nil {
        panic(err.Error())
    }

    truncate.Close()
}

func insertPhotos(db *sql.DB, photos []string) {
    var re = regexp.MustCompile(`(?m)photo_([0-9]+).json`)

    for _, filename := range photos {
        var flickrId = re.FindStringSubmatch(filename)[1]
        insert, err := db.Query("INSERT INTO photos (filename, flickr_id) VALUES (?, ?) RETURNING id", filename, flickrId)
        // if there is an error inserting, handle it
        if err != nil {
            panic(err.Error())
        }
        insert.Close()
    }
}

func findAlbums() Albums {
    file, err := ioutil.ReadFile("data/account/albums.json")

    if err != nil {
        panic(err.Error())
    }

    data := Albums{}

    _ = json.Unmarshal([]byte(file), &data)

    fmt.Printf("Found %d \n", len(data.Albums))

    return data
}

func main() {
    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/flickr")
    if err != nil {
        panic(err)
    }

    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    // var photos = findPhotos()
    var albums = findAlbums()

    fmt.Println(albums.Albums[0].Title)

    //resetTable(db)

    //fmt.Printf("Found %d files\n", len(photos))

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    // insertPhotos(db, photos)

    fmt.Println("Finished!")
}
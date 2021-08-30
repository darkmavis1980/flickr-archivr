package main

import (
    "database/sql"
    "time"
    "fmt"
    "regexp"
    "io/ioutil"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func fetchFiles() []string {
    files, err := ioutil.ReadDir("./account")
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

func main() {
    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/flickr")
    if err != nil {
        panic(err)
    }

    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    var photos = findPhotos()

    // for _, f := range photos {
    // 	fmt.Println(f)
    // }
    fmt.Printf("Found %d files\n", len(photos))

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    var re = regexp.MustCompile(`(?m)photo_([0-9]+).json`)

    for _, filename := range photos {
        var flickrId = re.FindStringSubmatch(filename)[1]
        insert, err := db.Query("INSERT INTO photos (filename, flickr_id) VALUES (?, ?)", filename, flickrId)

        // if there is an error inserting, handle it
        if err != nil {
            panic(err.Error())
        }
        insert.Close()
    }

    fmt.Println("Finished!")
}
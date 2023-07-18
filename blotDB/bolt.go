package main

import (
	"github.com/boltdb/bolt"
)

func main() {
	db, _ := bolt.Open("my.db", 0600, nil)
	defer db.Close()

}

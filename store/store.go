package store

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	models "go-api/models"
	utils "go-api/utils"
)

type Store struct {
	Db models.Db
}

func New() Store {
	db := models.Db{}
	store := Store{Db: db}

	utils.LoadKeys(&db)

	defer func() {
		go tempTimer(store)
	}()

	return store
}

// Get Item
func (k Store) Get(key string) string {
	return k.Db[key]
}

// Add Item
func (k Store) Add(value string) string {
	count := len(k.Db)
	count += 1

	newId := strconv.Itoa(count)

	k.Db[newId] = value
	return newId
}

// Flush Items
func (k Store) Flush() {
	for k2 := range k.Db {
		delete(k.Db, k2)
	}
}

func tempTimer(store Store) {
	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		log.Println("Ticker and tock")

		dbJson, _ := json.Marshal(store.Db)
		os.WriteFile("temp/data.json", dbJson, os.ModeAppend)
	}
}

package title

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"io/ioutil"
	"path/filepath"
	"time"
)

type Title struct {
	Id      string    `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
}

// Here we are using a json file as our "Database" for this demo
// A DB instance can be defined to replace the json storage
const Store = "store/title/title.json"

// Insert a new title to DB
func InsertTitle(title Title) error {

	var titles []Title

	// Get Store aboslute path
	absPath, _ := filepath.Abs(Store)

	// Reading content from DB
	contents, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(contents, &titles); err != nil {
		panic(err)
	}

	// Validate title
	if err := validateTitle(title.Title, titles); err != nil {
		return errors.Wrap(err, 0)
	}

	// Append new title to DB
	titles = append(titles, title)
	data, err := json.Marshal(titles)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(absPath, data, 0644); err != nil {
		return errors.Wrap(err, 0)
	}
	return nil
}

// Find all titles in DB
func FindTitles() ([]Title, error) {

	var titles []Title

	// Get Store aboslute path
	absPath, _ := filepath.Abs(Store)

	// Reading content from DB
	contents, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	// If empty content
	if len(contents) == 0 {
		return []Title{}, nil
	}

	if err := json.Unmarshal(contents, &titles); err != nil {
		panic(err)
	}
	return titles, nil
}

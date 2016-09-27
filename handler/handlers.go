// Handlers that handle http request
// Multiple handler files can be added to this location
// so we structure handler methods by http endpoint
package handler

import (
	"encoding/json"
	"fmt"
	"github.com/eagameslib-api/store/title"
	"net/http"
	"time"
)

// Printing some text on the Index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! The API is currently up and running...")
}

// GET "/tite" handler
func TitleGet(w http.ResponseWriter, r *http.Request) {
	setCors(w)

	titles, err := title.FindTitles()
	if err != nil {
		panic(err)
	}

	// Prepare response in JSON format
	setJsonHeader(w)
	if err := json.NewEncoder(w).Encode(prepareResponse(true, titles, "")); err != nil {
		panic(err)
	}

}

// POST "/tite" handler
func TitlePost(w http.ResponseWriter, r *http.Request) {
	setCors(w)

	// Unmarshall Data from HTTP Json Body
	var postData TitleName
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postData)
	if err != nil {
		panic(err)
	}

	// Generate unique id
	id := generateUniqueId()

	postTitle := title.Title{
		id,
		postData.Name,
		time.Now(),
	}

	// Prepare response in JSON format
	var response JsonResponse
	if err := title.InsertTitle(postTitle); err != nil {
		response = prepareResponse(false, postTitle, err.Error())
	} else {
		response = prepareResponse(true, postTitle, "")
	}

	setJsonHeader(w)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

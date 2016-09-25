package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    //"github.com/gorilla/mux"
    "github.com/eagames/store/title"
    "github.com/rs/xid"
    )

func Index(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    fmt.Fprintln(w, "Welcome! The API is currently up and running...")
}

func TitleGet(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")

    titles, err := title.FindTitles()
    if err != nil {
        panic(err)
    }
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(PrepareResponse(true,titles,"")); err != nil {
        panic(err)
    }
    
}

type TitleName struct {
    Name string `json:"title"`
}

func TitlePost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    var postData TitleName
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&postData)
    if err != nil {
        panic(err)
    }
    
    // generate unique id
    // TODO: seperate this out as a util
    guid := xid.New()
    id := guid.String()

    postTitle := title.Title{
        id,
        postData.Name,
        time.Now(),
    }

    var response JsonResponse

    if err := title.InsertTitle(postTitle); err != nil {
        response = PrepareResponse(false, postTitle, err.Error())
    } else {
        response = PrepareResponse(true, postTitle, "")
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    }
}

type JsonResponse struct {
    Success bool
    Body   interface{}
    Error string
}

func PrepareResponse(success bool, message interface{}, error string) JsonResponse{
    return JsonResponse{success, message, error}
}
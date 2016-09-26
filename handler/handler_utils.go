package handler

import (
	"net/http"
    "github.com/rs/xid"
    "github.com/eagames/config"
)

type JsonResponse struct {
    Success bool
    Body   interface{}
    Error string
}

type TitleName struct {
    Name string `json:"title"`
}

func prepareResponse(success bool, message interface{}, error string) JsonResponse{
    return JsonResponse{success, message, error}
}

func setCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", config.CORS)
}

func setJsonHeader(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func generateUniqueId() string {
    guid := xid.New()
    return guid.String()
}
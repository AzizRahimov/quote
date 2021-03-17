package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

func RespJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}


}


func IsTimePassed(check, date time.Time) bool {


	return check.After(date)
}


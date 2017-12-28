package main


import (
    "net/http"
    "encoding/json"
//    "fmt"
)

// Generate a function to return json through http.ResponseWriter 
func handlerJson(js []byte, err error) func(http.ResponseWriter) {
    return func (w http.ResponseWriter) {
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(js)
    }
}

/*
   Default  Handler

*/
func JonDon(w http.ResponseWriter, r *http.Request) {
    var action_err error
    msg := "JonDon message"
    if action_err != nil {
        handlerJson([]byte(""), action_err)(w)
    } else {
        js, json_err := json.Marshal(msg)
        handlerJson(js, json_err)(w)
    }
}


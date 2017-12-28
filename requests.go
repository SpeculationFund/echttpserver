package main


import (
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "net/http"
)


/*

    Define the RESTful API for crawler and calculator
*/
func ConfigRouter() http.Handler {
    r := mux.NewRouter()

    /*
        Ticker API
    */
    tickerAPI := r.PathPrefix("/ticker").Subrouter()
    tickerAPI.HandleFunc("", JonDon)
    tickerAPI.HandleFunc("/{ec}", JonDon)

    /*
        Market API
    */
    actionAPI := r.PathPrefix("/market").Subrouter()
    actionAPI.HandleFunc("/", JonDon)

    // CORS Configuration
    corsObj := handlers.AllowedOrigins([]string{"*"})

    h := handlers.CORS(corsObj)(r)

    return h
}



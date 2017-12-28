package main

import (
    "net/http"
)



func main() {

    // Configurate the RESTful API
    r := ConfigRouter()

    // Start http server
    http.ListenAndServe(":8000", r)

}


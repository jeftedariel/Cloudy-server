package setups;

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func SetupRoutes() {
    http.HandleFunc("/upload", uploadFile)
    http.HandleFunc("/download". downloadFile)
    http.ListenAndServe(":8080", nil)
}

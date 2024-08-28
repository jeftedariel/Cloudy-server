package main

import (
    "log"
    "net/http"
    "path/filepath"
    "strings"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
    //This is to look to the file that will be downloaded
    filePath := r.URL.Path[len("/download/"):]// And the folder storage which is where we save all the stuff
    filePath = filepath.Join("storage", filePath)

    // Prevent directory traversal attacks
    if !strings.HasPrefix(filepath.Clean(filePath), "storage") {
        http.Error(w, "Invalid file path", http.StatusBadRequest)
        return
    }
    //then if we can find it, we send it back
    http.ServeFile(w, r, filePath)
}


func main() {
    http.HandleFunc("/download/", downloadHandler)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Error starting server:", err)
    }
}


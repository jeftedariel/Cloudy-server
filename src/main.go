package main

import (
    "log"
    "net/http"
    "path/filepath"
    "math/rand"
    "strings"
    "strconv"
    "time"
    "fmt"
    "os"
    "io"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
    //This is to look to the file that will be downloaded
    filePath := r.URL.Path[len("/download/"):]// And the folder storage which is where we save all the stuff
    filePath = filepath.Join("storage", filePath)

    if !strings.HasPrefix(filepath.Clean(filePath), "storage") {
        http.Error(w, "Invalid file path", http.StatusBadRequest)
        return
    }
    //then if we can find it, we send it back
    http.ServeFile(w, r, filePath)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(20000 << 20) // 20Gibityes of limit

    // gettin the file from the form data
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error obtaining the file", http.StatusBadRequest)
        log.Println("Error obtaining the file:", err)
        return
    }
    defer file.Close()
    
    //creates a random number for the folder
    rand.Seed(time.Now().UnixNano())
    min := 10000000
    max := 99999999
    folderId := strconv.Itoa(rand.Intn(max-min+1)+min)
    fmt.Println("A New File has been uploaded, saved at: ", folderId)

    errFolder := os.Mkdir(filepath.Join("storage",folderId), 0755)
    if err != nil {
      fmt.Println(errFolder) 
      return
   }

    filePath := filepath.Join("storage",folderId,handler.Filename)
    dst, err := os.Create(filePath)

    if err != nil {
        http.Error(w, "Error saving the file", http.StatusInternalServerError)
        log.Println("Error saving the file:", err)
        return
    }
    defer dst.Close()

    // Copy the uploaded file's content to the destination file
    _, err = io.Copy(dst, file)
    if err != nil {
        http.Error(w, "Error copying the file", http.StatusInternalServerError)
        log.Println("Error copying the file:", err)
        return
    }

    // Respond with a success message
    w.Write([]byte("The File has been uploaded successfully"))
}



func main() {
    http.HandleFunc("/download/", downloadFile)
    http.HandleFunc("/upload", uploadFile)
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Error starting server:", err)
    }
}


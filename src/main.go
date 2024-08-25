package main
import (
  "fmt"
  "os"
  //"time"
  //"net/http"
  "github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    fmt.Println("Error loading .env file")
  }

  return os.Getenv(key)
}

func main() {
  dotenv := goDotEnvVariable("PORT")

  fmt.Printf(dotenv + "\n")
}

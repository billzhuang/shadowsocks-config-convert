package main

import "net/http"
import "os"

func main() {
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir(os.Getenv("USERPROFILE")+"\\Dropbox\\backup"))))
}

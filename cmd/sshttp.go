package main

import "net/http"
import "os"

func main() {
	panic(http.ListenAndServe(":8119", http.FileServer(http.Dir(os.Getenv("USERPROFILE")+"\\Dropbox\\backup\\win\\"))))
}

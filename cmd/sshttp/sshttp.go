package main

import (
	"fmt"
	sc "github.com/shadowsocks-configfile/shadowsocksconfigfile"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8119", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", 3600))
	if path == "win.pac" {
		fmt.Fprintf(w, "%s", readContent())
	} else {
		fmt.Fprintf(w, "Hi there, I love bill!")
	}
}

func readContent() string {
	dat, err := ioutil.ReadFile(os.Getenv("USERPROFILE") + "\\Dropbox\\backup\\win\\win.pac")
	sc.Check(err)

	return strings.Replace(string(dat), "COMPUTERNAME", os.Getenv("COMPUTERNAME"), -1)
}

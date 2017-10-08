package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

func init() {
	http.HandleFunc("/pageRender", pageRender)
}

func main() {
	log.Println("Server running at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pageRender(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(r.FormValue("url"))
	if err != nil {
		http.Error(w, "Bad URL "+r.FormValue("url"), http.StatusBadRequest)
		return
	}

	tmpFile, err := ioutil.TempFile("", "screenshot-")
	if err != nil {
		http.Error(w, "Unable to create temp file "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	querySelector := r.FormValue("querySelector")
	clickSelector := r.FormValue("clickSelector")
	if r.FormValue("click") != "" {
		clickSelector = r.FormValue("click")
	}
	size := r.FormValue("size")

	// Setup command to be executed
	args := []string{"phantomjs", "/var/lib/render.js",
		url.String(), tmpFile.Name(), querySelector, clickSelector, size}
	cmd := exec.Command("xvfb-run", args...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Error rendering page "+err.Error(), http.StatusInternalServerError)
		w.Write(b)
		return
	}
	log.Println(string(b))

	log.Println("Serving image: ", tmpFile.Name())
	http.ServeFile(w, r, tmpFile.Name())
}

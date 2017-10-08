package main

import (
	"log"
	"net/http"
	"net/url"
	"os/exec"
)

func init() {
	http.HandleFunc("/pageRender", pageRender)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pageRender(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(r.FormValue("url"))
	if err != nil {
		http.Error(w, "Bad URL "+r.FormValue("url"), http.StatusBadRequest)
		return
	}

	querySelector := r.FormValue("querySelector")
	clickSelector := r.FormValue("clickSelector")
	size := r.FormValue("size")

	args := []string{"phantomjs", "/var/lib/render.js", url.String(), querySelector, clickSelector, size}
	cmd := exec.Command("xvfb-run", args...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Error rendering page "+err.Error(), http.StatusInternalServerError)
		w.Write(b)
		return
	}
	log.Println(string(b))

	http.ServeFile(w, r, "/tmp/screenshot.jpg")
}

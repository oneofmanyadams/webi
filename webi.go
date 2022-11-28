package webi

import (
	//"html/template"
	"log"
	"net/http"
)

type Webi struct {
	SourceFilesLocation string
	TemplatePath string
	StylePath string
}

func (s *Webi) Start() {
	http.HandleFunc("/style.css", func(w http.ResponseWriter, req *http.Request) {
 		http.ServeFile(w, req, s.StylePath)
 	})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
 		http.ServeFile(w, req, s.TemplatePath)
 	})
	log.Fatal(http.ListenAndServe(":8080",nil))
}
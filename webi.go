package webi

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type TemplateContent struct {
	Header string
	Content []string
}

func Start(tc TemplateContent) {
	http.HandleFunc("/style.css", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		io.WriteString(w, style_string)
 	})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmp, err := template.New("template.html").Parse(template_string)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Error loading template.", http.StatusInternalServerError)			
			return
		}
		tmp.Execute(w, tc)

 	})
	log.Fatal(http.ListenAndServe(":8080",nil))
}
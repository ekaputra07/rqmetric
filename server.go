package main

import (
  "fmt"
  "log"
  "net/http"
  "html/template"

  "github.com/gobuffalo/packr"
)

const (
  PORT = 5000
)

type TemplateData struct {
  Requests []Request
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
  tplBox := packr.NewBox("./templates")
  indexStr, _ := tplBox.FindString("index.html")

  tpl, _ := template.New("index").Parse(indexStr)

  var reqs []Request
  reqs = append(reqs,
    Request{"http://localhost", 1,10,float64(5.5),2000, 1,2,3,4},
    Request{"http://localhost", 2,50,float64(1.5),200, 2,2,3,4},
    Request{"http://localhost", 6,43,float64(3.5),500, 3,2,3,4},
    Request{"http://localhost", 2,12,float64(4.5),7000, 5,2,3,4},
    Request{"http://localhost", 6,13,float64(7.5),900, 4,2,3,4},
    Request{"http://localhost", 9,15,float64(9.5),100, 6,2,3,4},)

  data := TemplateData{reqs}

  tpl.Execute(w, data)
}

func Serve(id string) {
  mux := http.NewServeMux()

  staticBox := packr.NewBox("./static")
  fs := http.FileServer(staticBox)

  mux.Handle("/static/", fs)
  mux.HandleFunc("/", handleIndex)

  log.Printf("Listening on http://localhost:%v...\n", PORT)
  if err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), mux); err != http.ErrServerClosed {
    log.Fatal(`Server failed to start: `, err)
  }
}
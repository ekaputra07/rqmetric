package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gobuffalo/packr"
)

// TemplateData is a data model that will be passed as a template data.
type TemplateData struct {
	Requests   []Request
	AppName    string
	AppVersion string
}

// loadData loads data from the CSV file and returns it as a TemplateData object.
func loadData(id string) TemplateData {
	rows := ReadCSV(fmt.Sprintf("rqmetric_output_%v.csv", id))
	var reqs []Request
	for index, row := range rows {
		// skip the header
		if index == 0 {
			continue
		}

		URL := row[0]
		MinTime, _ := strconv.Atoi(row[1])
		MaxTime, _ := strconv.Atoi(row[2])
		AvgTime, _ := strconv.ParseFloat(row[3], 64)
		Count, _ := strconv.Atoi(row[4])
		OkResponseCount, _ := strconv.Atoi(row[5])
		RedirectResponseCount, _ := strconv.Atoi(row[6])
		ClientErrorCount, _ := strconv.Atoi(row[7])
		ServerErrorCount, _ := strconv.Atoi(row[8])

		reqs = append(reqs, Request{URL, MinTime, MaxTime, AvgTime, Count, OkResponseCount, RedirectResponseCount, ClientErrorCount, ServerErrorCount})
	}

	return TemplateData{reqs, AppName, AppVersion}
}

// Serve starts a simple web server to serve the data on specific port.
func Serve(id string, port string) {

	data := loadData(id)
	tplBox := packr.NewBox("./templates")
	staticBox := packr.NewBox("./static")

	mux := http.NewServeMux()

	// serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticBox)))

	// serve index page
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		indexStr, _ := tplBox.FindString("index.html")
		tpl, _ := template.New("index").Parse(indexStr)
		tpl.Execute(w, data)
	})

	log.Printf("Listening on http://localhost:%v ...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux); err != http.ErrServerClosed {
		log.Fatal(`Server failed to start: `, err)
	}
}

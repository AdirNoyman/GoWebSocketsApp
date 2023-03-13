package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"log"
	"net/http"
)

var views = jet.NewSet(jet.NewOSFileSystemLoader("./html"), jet.InDevelopmentMode())

func Home(w http.ResponseWriter, r *http.Request) {

	err := renderPage(w, "home.jet", nil)
	if err != nil {

		log.Println("There was an error 😩", err)

	}

}

// data = the data we want to pass to our html template(tmpl)
func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {

	view, err := views.GetTemplate(tmpl)

	if err != nil {

		log.Println("There was an error 😩", err)
		return err
	}

	// the third argument is for data that can be passed in the session context
	err = view.Execute(w, data, nil)
	if err != nil {

		log.Println("There was an error 😩", err)
		return err
	}

	return nil
}

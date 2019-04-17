package home

import (
	"html/template"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

var box = packr.NewBox("../../templates/")

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	appRouter := router.PathPrefix("/").Subrouter()
	appRouter.HandleFunc("", Home)
	return router
}

func Home(w http.ResponseWriter, r *http.Request) {
	html := box.String("index.html")
	t, _ := template.New("index").Parse(html)
	t.Execute(w, "hello world")
}

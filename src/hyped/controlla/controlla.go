package controlla

import (
	"html/template"
	"net/http"
)

// Register Controllers
var (
	indexController index
  topController top
)

// InitializeStatic manages routing setup and static files
func InitializeStatic() {
	// Static public files
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/styles/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
}

// Assign templates to the respective controllers
func configureTemplates(templates *template.Template) {
	indexController.indexTemplate = templates.Lookup("index.html")
}

// Apply Route Handlers
func InitializeRoutes(templates *template.Template) {
	configureTemplates(templates)

	indexController.ApplyEndpoint()
	topController.ApplyEndpoint()
}

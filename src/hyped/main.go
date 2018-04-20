package main

import (
	"html/template"
	"hyped/controlla"
	"net/http"
)

func main() {
	// Setup template prior to handling requests
	templates := processTemplates()

  // Handle request for statics resourses
  controlla.InitializeStatic()

  // Handle request for defined endpoints
  controlla.InitializeRoutes(templates)

	// 	Set server to listen on port 8000
	http.ListenAndServe(":8000", nil)
}

func processTemplates() *template.Template {
	result := template.New("templates")
	const basePath = `templates`
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}

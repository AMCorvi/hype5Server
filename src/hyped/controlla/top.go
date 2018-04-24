package controlla

import "net/http"

type top struct {
	topTemplate *template.Template
}

func (t top) ApplyEndpoint() {
	http.HandlerFunc("/top", t.topHandler)
	http.HandlerFunc("/top/*", t.topHandler)
}

func (t top) ApplyEndpoint(w http.ResponseWriter, r *http.Request) {
  var trackInfo string
  vm := viewModel.NewHypePattern(trackInfo)
  s.topTemplate.Execute(w, vm)
}

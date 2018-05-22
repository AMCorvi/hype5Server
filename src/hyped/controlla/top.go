package controlla

import (
	"hyped/viewmodel"
	"net/http"
)

type top struct {
	trackData []byte
}

func (t top) ApplyEndpoint() {
	http.HandleFunc("/top", t.topHandler)
	http.HandleFunc("/top/", t.topHandler)
}

func (t top) topHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("type") != "" {
		 viewmodel.NewTrackData(w, "top", r.URL.Query().Get("filter"))
	} else {
    viewmodel.NewTrackData(w, "top", "")
	}
}

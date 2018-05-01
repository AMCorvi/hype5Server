package controlla

import (
	"hyped/model"
	"net/http"
)

type top struct {
}

func (t top) ApplyEndpoint() {
	http.HandleFunc("/top", t.topHandler)
}

func (t top) topHandler(w http.ResponseWriter, r *http.Request) {
	trackData := model.QueryTrackData()
	w.Write(trackData)
}

package viewmodel

import (
	"encoding/json"
	"hyped/model"
	"log"
	"net/http"
)

// // TrackInfo structure of individual track datas returned by specific request pattern
// type TrackInfo struct {
// 	// Artist is the name of the track artist
// 	Artist string `json:artist`
// 	// BlogName the name of the publication posting about the track
// 	BlogName string `json:blogname`
// 	// Descrip is a excerpt of the orginating BlogPost of the Track
// 	Descrip string `json:"excerpt"`
// 	// Song is the name of the Track
// 	Song string `json:"trackname"`
// 	// Link to the blog post providing of showcasing the track
// 	Link string `json:"blogurl"`
//
// 	// Thumbnail is string that contains at minimum a 'color' (in rgb format)
// 	// associated the tracks popularity on Hypemachine.
// 	Thumbnail string `json:"imageData`
// }

func NewTrackData(w http.ResponseWriter, trackfilter, tracktype string) {
	// Query for track data and convert into json
	// Then write JSON to respone io.Writer
	var trackData []model.TrackInfo = model.QueryTrackData(trackfilter, tracktype)
	enc := json.NewEncoder(w)
	err := enc.Encode(trackData)
	if err != nil {
		log.Println(err)
	}
}

package viewmodel

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// TrackInfo structure of individual track datas returned by specific request pattern
type TrackInfo struct {
	// Artist is the name of the track artist
	Artist string
	// BlogName the name of the publication posting about the track
	BlogName string
	// Descrip is a excerpt of the orginating BlogPost of the Track
	Descrip string
	// Song is the name of the Track
	Song string
	// Link to the blog post providing of showcasing the track
	Link string

	// Thumbnail is string that contains at minimum a 'color' (in rgb format)
	// associated the tracks popularity on Hypemachine.
	Thumbnail string
}

func NewTrackListing(trackdata []byte) []TrackInfo {

	filename := "./src/hyped/model/trackinfo.json"

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}

	// Convert JSON data in structs of type TrackInfo
	var result [][]TrackInfo
	err = json.Unmarshal(file, &result)
	if err != nil {
		log.Println("json:Marshal error", err)
	}

	// Return []TrackInfo data. unwrapping the top level array
	return result[0]
}

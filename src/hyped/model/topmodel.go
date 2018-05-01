package model

import (
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

func QueryTrackData(type, filter) []byte {

  // Receive type and filter type
  var cmd *Cmd

  if type != nil {
    if filter != nil {
      cmd = exec.Command("casperjs","--verbose", fmt.Sprintf("--hypetype=%s",type)
      , fmt.Sprintf("--hypefilter=%s", filter))
    } else {
      cmd = exec.Command("casperjs","--verbose", fmt.Sprintf("--hypetype=%s", type))
    }
  }

  var data bytes.Buffer
	cmd.Stdout = &data
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}



  //Convert JSON data in structs of type TrackInfo
  var result [][]TrackInfo
  err = json.Unmarshal(data, &result)
  if err != nil {
    log.Println("Marshalling Error:", err)
  }

  // Return []TrackInfo data unwrapping the top level array
  return result[0]

}

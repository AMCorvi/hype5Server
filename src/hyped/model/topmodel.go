package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// TrackInfo structure of individual track datas returned by specific request pattern
type TrackInfo struct {
	// Artist is the name of the track artist
	Artist string `json:artist`
	// BlogName the name of the publication posting about the track
	BlogName string `json:blogname`
	// Descrip is a excerpt of the orginating BlogPost of the Track
	Descrip string `json:"excerpt"`
	// Song is the name of the Track
	Song string `json:"trackname"`
	// Link to the blog post providing of showcasing the track
	Link string `json:"blogurl"`

	// Thumbnail is string that contains at minimum a 'color' (in rgb format)
	// associated the tracks popularity on Hypemachine.
	Thumbnail string `json:"imageData`
}

// QueryTrackData return array of TrackInfo structs based on the provided
// filter and type requested based on the endpoint
func QueryTrackData(hfilter, htype string) []TrackInfo {

  // Variable to receive output of crawler command
  var data bytes.Buffer

  // setup command
  var cmd *exec.Cmd = configureCmd(hfilter, htype, data)

	// Run Crawler command
	err := cmd.Run()
	if err != nil {
    log.Println("Command Run Error: ", err)
	}

	//Convert JSON data in structs of type TrackInfo
	var result [][]TrackInfo
	err = json.Unmarshal(
		[]byte(data.String()),
		&result)

  // handle error
	if err != nil {
		log.Println("Marshalling Error: ", err)
	}

	// Return []TrackInfo data unwrapping the top level array
	return result[0]
}

// Helper function for processing exec call
func configureCmd(hfilter, htype string, data bytes.Buffer) *exec.Cmd {
  // Receive type and filter type
  var cmd *exec.Cmd

  if hfilter != "" {
    if htype != "" {
      // If type and filter are provided run crawler command with both flags
      cmd = exec.Command("casperjs", "--verbose", "crawlerjs", fmt.Sprintf("--hypetype=%s", htype), fmt.Sprintf("--hypefilter=%s", hfilter))
    } else {
      // if filter is not provided but type is provided run crawler script with type flag
      cmd = exec.Command("casperjs", "--verbose", "crawlerjs", fmt.Sprintf("--hypefilter=%s", hfilter))
    }
    // Soley run crawler command if filter and type are not provided
    cmd = exec.Command("casperjs", "--verbose", "crawlerjs")
  }

  // Configure destination for crawler commands standard output
  cmd.Stdout = &data

  return cmd
}

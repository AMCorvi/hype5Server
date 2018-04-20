package viewmodel

import "fmt"

// TrackInfo structure of individual track datas returned by specific request pattern
type TrackInfo struct {
	Artist    string
	BlogName  string
	Excerpt   string
	TrackName string

	// ImageData is map that contains at minimum an 'color' (in rgb format)
	// associated the tracks popularity on Hypemachine.
	// Optionally: (Subject to future implentation): map shall contain jpg
	// 'image' provided for the track by the blog it came from: map shall
	// contain jpg image provided for the track by the blog it came from.
	ImageData map[string]string
}

// Pattern is the property description of the any possible request pattern.
// All field are expected to be explictily provided with exception to
// the 'DemoOutput' slice which should be generated  programatically.
type Pattern struct {
	Type       string
	Category   []string
	DemoCode   string
	DemoOutput []TrackInfo
}

func NewPattern() {
	example := Pattern{
		Type:     "Latest",
		Category: []string{"top", "remixes", "noremixes"},
		DemoCode: "hype5server.com/noremixes?noremixes",
	}

	fmt.Println(example)
	// TODO: make request to endpoint that serves thats serves pattern
	// and loop over JSON in order to construct array of TrackInfo elements

}

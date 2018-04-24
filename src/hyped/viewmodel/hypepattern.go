package viewmodel

import "fmt"


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

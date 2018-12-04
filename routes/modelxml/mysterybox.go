package modelxml

import "encoding/xml"

type Minigames struct {
	XMLName xml.Name `xml:"Minigames"`
	Version string   `xml:"version,attr"`
}

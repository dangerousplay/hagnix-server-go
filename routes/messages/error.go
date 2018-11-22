package messages

import "encoding/xml"

type Error struct {
	XMLName xml.Name `xml:"Error"`
	RawXml  string   `xml:",innerxml"`
}

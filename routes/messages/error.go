package messages

import "encoding/xml"

var DefaultError = &Error{RawXml: "Internal server error"}
var BadRequest = &Error{RawXml: "Bad Request"}

type Error struct {
	XMLName xml.Name `xml:"Error"`
	RawXml  string   `xml:",innerxml"`
}

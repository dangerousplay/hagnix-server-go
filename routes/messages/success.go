package messages

import "encoding/xml"

var DefaultSuccess = &Success{}

type Success struct {
	XMLName xml.Name `xml:"Success"`
	Message string   `xml:",innerxml"`
}

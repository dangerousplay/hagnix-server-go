package messages

import "encoding/xml"

type Sucess struct {
	XMLName xml.Name `xml:"Sucess"`
	Message string   `xml:",innerxml"`
}

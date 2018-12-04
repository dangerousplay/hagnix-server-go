package modelxml

import (
	"encoding/xml"
)

type Package struct {
	XMLName     xml.Name `xml:"Package"`
	Id          int      `xml:"id,attr"`
	Name        string   `xml:"Name"`
	Price       int      `xml:"Price"`
	Quantity    int      `xml:"Quantity"`
	MaxPurchase int      `xml:"MaxPurchase"`
	Weight      int      `xml:"Weight"`
	BgURL       string   `xml:"BgURL"`
	EndDate     string   `xml:"EndDate"`
}

type PackageWrapper struct {
	XMLName  xml.Name `xml:"Packages"`
	Packages []Package
}

type PackageResponse struct {
	XMLName  xml.Name `xml:"PackageResponse"`
	Packages PackageWrapper
}

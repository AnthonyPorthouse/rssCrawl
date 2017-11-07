package feeds

import (
	"encoding/xml"
	"time"
)

type AtomFeed struct {
	XMLName xml.Name
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Link    atomLink `xml:"link"`
	Updated atomDate `xml:"updated"`

	Entries []Entry `xml:"entry"`
}

type Entry struct {
	XMLName xml.Name
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Updated atomDate `xml:"updated"`
	Link    string   `xml:"link"`
	Summary string   `xml:"summary"`
	Content string   `xml:"content"`
}

type atomLink struct {
	Href     string `xml:"href,attr"`
	Rel      string `xml:"rel,attr"`
	Type     string `xml:"type,attr"`
	Hreflang string `xml:"hreflang,attr"`
	Title    string `xml:"title,attr"`
	Length   string `xml:"length,attr"`
}

const (
	LINK_ALTERNATE = "alternate"
	LINK_ENCLOSURE = "enclosure"
	LINK_RELATED   = "related"
	LINK_SELF      = "self"
	LINK_VIA       = "via"
)

type atomDate struct {
	time.Time
}

func (date *atomDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)

	parse, err := time.Parse(time.RFC3339, v)
	if err == nil {
		*date = atomDate{parse}
		return nil
	}
	return err
}

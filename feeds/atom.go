package feeds

import (
	"encoding/xml"
	"time"
)

type AtomFeed struct {
	XMLName xml.Name
	ID      string `xml:"id"`
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Updated string `xml:"updated"`

	Entries []Entry `xml:"entry"`
}

type Entry struct {
	XMLName xml.Name
	ID      string `xml:"id"`
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
	Link    string `xml:"link"`
	Summary string `xml:"summary"`
	Content string `xml:"content"`
}

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

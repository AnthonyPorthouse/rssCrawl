package feeds

import (
	"encoding/xml"
	"time"
)

type AtomFeed struct {
	XMLName xml.Name

	// Required
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Updated atomDate `xml:"updated"`

	// Recommended
	Authors []atomPerson `xml:"author"`
	Link    atomLink     `xml:"link"`

	// Optional
	Categories   []atomPerson `xml:"category"`
	Contributors []string     `xml:"contributor"`
	Generator    string       `xml:"generator"`
	Icon         string       `xml:"icon"`
	Logo         string       `xml:"logo"`
	Rights       string       `xml:"rights"`
	Subtitle     string       `xml:"subtitle"`

	Entries []Entry `xml:"entry"`
}

type Entry struct {
	XMLName xml.Name

	// Required
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Updated atomDate `xml:"updated"`

	// Recommended
	Authors []atomPerson `xml:"author"`
	Content string       `xml:"content"`
	Links   []atomLink   `xml:"link"`
	Summary string       `xml:"summary"`

	// Optional
	Categories   []string     `xml:"category"`
	Contributors []atomPerson `xml:"contributor"`
	Published    atomDate     `xml:"published"`
	Rights       string       `xml:"rights"`
}

type atomLink struct {
	Href     string `xml:"href,attr"`
	Rel      string `xml:"rel,attr"`
	Type     string `xml:"type,attr"`
	Hreflang string `xml:"hreflang,attr"`
	Title    string `xml:"title,attr"`
	Length   string `xml:"length,attr"`
}

type atomPerson struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri"`
	Email string `xml:"email"`
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

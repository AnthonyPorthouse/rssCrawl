package feeds

import (
	"encoding/xml"
	"time"
)

type AtomFeed struct {
	// Required
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Updated atomDate `xml:"updated"`

	// Recommended
	Authors []atomPerson `xml:"author,omitempty"`
	Link    atomLink     `xml:"link,omitempty"`

	// Optional
	Categories   []atomPerson `xml:"category,omitempty"`
	Contributors []string     `xml:"contributor,omitempty"`
	Generator    string       `xml:"generator,omitempty"`
	Icon         string       `xml:"icon,omitempty"`
	Logo         string       `xml:"logo,omitempty"`
	Rights       string       `xml:"rights,omitempty"`
	Subtitle     string       `xml:"subtitle,omitempty"`

	Entries []Entry `xml:"entry"`
}

type Entry struct {
	// Required
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Updated atomDate `xml:"updated"`

	// Recommended
	Authors []atomPerson `xml:"author,omitempty"`
	Content string       `xml:"content,omitempty"`
	Links   []atomLink   `xml:"link,omitempty"`
	Summary string       `xml:"summary,omitempty"`

	// Optional
	Categories   []string     `xml:"category,omitempty"`
	Contributors []atomPerson `xml:"contributor,omitempty"`
	Published    atomDate     `xml:"published,omitempty"`
	Rights       string       `xml:"rights,omitempty"`
}

type atomLink struct {
	Href     string `xml:"href,attr"`
	Rel      string `xml:"rel,attr,omitempty"`
	Type     string `xml:"type,attr,omitempty"`
	Hreflang string `xml:"hreflang,attr,omitempty"`
	Title    string `xml:"title,attr,omitempty"`
	Length   string `xml:"length,attr,omitempty"`
}

type atomPerson struct {
	Name  string `xml:"name,omitempty"`
	URI   string `xml:"uri,omitempty"`
	Email string `xml:"email,omitempty"`
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

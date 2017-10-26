package feeds

import (
	"encoding/xml"
	"time"
)

type RSSFeed struct {
	XMLName xml.Name
	Version string  `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	XMLName         xml.Name
	Title           string  `xml:"title"`
	Link            string  `xml:"link"`
	Language        string  `xml:"language"`
	Description     string  `xml:"description"`
	LastBuildDate   rssDate `xml:"lastBuildDate"`
	UpdatePeriod    string  `xml:"updatePeriod"`
	UpdateFrequency string  `xml:"updateFrequency"`
	Items           []Item  `xml:"item"`
}

type Item struct {
	XMLName     xml.Name
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Categories  []string `xml:"category"`
	Comments    string   `xml:"comments"`
	PubDate     rssDate  `xml:"pubDate"`
	GUID        string   `xml:"guid"`
	Description string   `xml:"description"`
	Encoded     string   `xml:"encoded"`
}

type rssDate struct {
	time.Time
}

func (date *rssDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)

	parse, err := time.Parse(time.RFC1123Z, v)
	if err == nil {
		*date = rssDate{parse}
		return nil
	}

	parse, err = time.Parse(time.RFC1123, v)
	if err == nil {
		*date = rssDate{parse}
		return nil
	}

	return err
}
